package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"veterinaria-server/internal/accesos"
	"veterinaria-server/internal/album"
	"veterinaria-server/internal/auth"
	"veterinaria-server/internal/clientes"
	"veterinaria-server/internal/compra"
	"veterinaria-server/internal/config"
	"veterinaria-server/internal/consultas"
	"veterinaria-server/internal/detalle_compra"
	"veterinaria-server/internal/detalle_examen_cualitativo"
	"veterinaria-server/internal/detalle_examen_cuantitativo"
	"veterinaria-server/internal/detalle_examen_informativo"
	"veterinaria-server/internal/errors"
	"veterinaria-server/internal/especies"
	"veterinaria-server/internal/examen_mascota"
	"veterinaria-server/internal/factura"
	"veterinaria-server/internal/generos"
	"veterinaria-server/internal/healthcheck"
	"veterinaria-server/internal/lote"
	"veterinaria-server/internal/mascotas"
	"veterinaria-server/internal/medida"
	"veterinaria-server/internal/productos"
	"veterinaria-server/internal/proveedor"
	"veterinaria-server/internal/proveedor_producto"
	"veterinaria-server/internal/stock_individual"
	"veterinaria-server/internal/tipo_examen"
	"veterinaria-server/internal/unidad"
	"veterinaria-server/pkg/accesslog"
	"veterinaria-server/pkg/dbcontext"
	"veterinaria-server/pkg/log"

	dbx "github.com/go-ozzo/ozzo-dbx"
	routing "github.com/go-ozzo/ozzo-routing/v2"
	"github.com/go-ozzo/ozzo-routing/v2/content"
	"github.com/go-ozzo/ozzo-routing/v2/cors"
	"github.com/go-ozzo/ozzo-routing/v2/file"
	_ "github.com/go-sql-driver/mysql"
)

// Version indicates the current version of the application.
var Version = "1.0.0"

var flagConfig = flag.String("config", "./config/local.yml", "path to the config file")

func main() {
	flag.Parse()
	// create root logger tagged with server version
	logger := log.New().With(context.TODO(), "version", Version)

	// load application configurations
	cfg, err := config.Load(*flagConfig, logger)
	if err != nil {
		logger.Errorf("failed to load application configuration: %s", err)
		os.Exit(-1)
	}

	// connect to the database
	db, err := dbx.MustOpen("mysql", cfg.DSN)
	if err != nil {
		logger.Error(err)
		os.Exit(-1)
	}
	db.QueryLogFunc = logDBQuery(logger)
	db.ExecLogFunc = logDBExec(logger)
	defer func() {
		if err := db.Close(); err != nil {
			logger.Error(err)
		}
	}()

	// build HTTP server
	address := fmt.Sprintf(":%v", cfg.ServerPort)
	hs := &http.Server{
		Addr:    address,
		Handler: buildHandler(logger, dbcontext.New(db), cfg),
	}

	// start the HTTP server with graceful shutdown
	go routing.GracefulShutdown(hs, 10*time.Second, logger.Infof)
	logger.Infof("server %v is running at %v", Version, address)
	if err := hs.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Error(err)
		os.Exit(-1)
	}
}

// buildHandler sets up the HTTP routing and builds an HTTP handler.
func buildHandler(logger log.Logger, db *dbcontext.DB, cfg *config.Config) http.Handler {
	router := routing.New()

	router.Use(
		accesslog.Handler(logger),
		errors.Handler(logger),
		content.TypeNegotiator(content.JSON),
		cors.Handler(cors.AllowAll),
		db.TransactionHandler(),
	)

	healthcheck.RegisterHandlers(router, Version)

	rg := router.Group("/v1")

	authHandler := auth.Handler(cfg.JWTSigningKey)

	album.RegisterHandlers(rg.Group(""),
		album.NewService(album.NewRepository(db, logger), logger),
		authHandler, logger,
	)

	accesos.RegisterHandlers(rg.Group(""),
		accesos.NewService(accesos.NewRepository(db, logger), logger),
		authHandler, logger,
	)

	generos.RegisterHandlers(rg.Group(""),
		generos.NewService(generos.NewRepository(db, logger), logger),
		authHandler, logger,
	)

	especies.RegisterHandlers(rg.Group(""),
		especies.NewService(especies.NewRepository(db, logger), logger),
		authHandler, logger,
	)

	clientes.RegisterHandlers(rg.Group(""),
		clientes.NewService(clientes.NewRepository(db, logger), logger),
		authHandler, logger,
	)

	mascotas.RegisterHandlers(rg.Group(""),
		mascotas.NewService(mascotas.NewRepository(db, logger), logger),
		authHandler, logger,
	)

	tipo_examen.RegisterHandlers(rg.Group(""),
		tipo_examen.NewService(tipo_examen.NewRepository(db, logger), logger),
		authHandler, logger, db,
	)

	detalle_examen_cualitativo.RegisterHandlers(rg.Group(""),
		detalle_examen_cualitativo.NewService(detalle_examen_cualitativo.NewRepository(db, logger), logger),
		authHandler, logger,
	)

	detalle_examen_cuantitativo.RegisterHandlers(rg.Group(""),
		detalle_examen_cuantitativo.NewService(detalle_examen_cuantitativo.NewRepository(db, logger), logger),
		authHandler, logger,
	)

	detalle_examen_informativo.RegisterHandlers(rg.Group(""),
		detalle_examen_informativo.NewService(detalle_examen_informativo.NewRepository(db, logger), logger),
		authHandler, logger,
	)

	examen_mascota.RegisterHandlers(rg.Group(""),
		examen_mascota.NewService(examen_mascota.NewRepository(db, logger), logger),
		authHandler, logger,
	)

	factura.RegisterHandlers(rg.Group(""),
		factura.NewService(factura.NewRepository(db, logger), logger),
		authHandler, logger, db,
	)

	compra.RegisterHandlers(rg.Group(""),
		compra.NewService(compra.NewRepository(db, logger), logger),
		authHandler, logger, db,
	)

	detalle_compra.RegisterHandlers(rg.Group(""),
		detalle_compra.NewService(detalle_compra.NewRepository(db, logger), logger),
		authHandler, logger,
	)

	consultas.RegisterHandlers(rg.Group(""),
		consultas.NewService(consultas.NewRepository(db, logger), logger),
		authHandler, logger,
	)

	proveedor.RegisterHandlers(rg.Group(""),
		proveedor.NewService(proveedor.NewRepository(db, logger), logger),
		authHandler, logger,
	)

	productos.RegisterHandlers(rg.Group(""),
		productos.NewService(productos.NewRepository(db, logger), logger),
		authHandler, logger,
	)

	proveedor_producto.RegisterHandlers(rg.Group(""),
		proveedor_producto.NewService(proveedor_producto.NewRepository(db, logger), logger),
		authHandler, logger,
	)

	lote.RegisterHandlers(rg.Group(""),
		lote.NewService(lote.NewRepository(db, logger), logger),
		authHandler, logger,
	)

	stock_individual.RegisterHandlers(rg.Group(""),
		stock_individual.NewService(stock_individual.NewRepository(db, logger), logger),
		authHandler, logger,
	)

	unidad.RegisterHandlers(rg.Group(""),
		unidad.NewService(unidad.NewRepository(db, logger), logger),
		authHandler, logger,
	)

	medida.RegisterHandlers(rg.Group(""),
		medida.NewService(medida.NewRepository(db, logger), logger),
		authHandler, logger,
	)

	auth.RegisterHandlers(rg.Group(""),
		auth.NewService(db, cfg.JWTSigningKey, cfg.JWTExpiration, logger),
		logger,
	)

	// Serving Static Files
	rg.Get("/files/*", file.Server(file.PathMap{
		"/v1/files": "/resources/",
	}))

	return router
}

// logDBQuery returns a logging function that can be used to log SQL queries.
func logDBQuery(logger log.Logger) dbx.QueryLogFunc {
	return func(ctx context.Context, t time.Duration, sql string, rows *sql.Rows, err error) {
		if err == nil {
			logger.With(ctx, "duration", t.Milliseconds(), "sql", sql).Info("DB query successful")
		} else {
			logger.With(ctx, "sql", sql).Errorf("DB query error: %v", err)
		}
	}
}

// logDBExec returns a logging function that can be used to log SQL executions.
func logDBExec(logger log.Logger) dbx.ExecLogFunc {
	return func(ctx context.Context, t time.Duration, sql string, result sql.Result, err error) {
		if err == nil {
			logger.With(ctx, "duration", t.Milliseconds(), "sql", sql).Info("DB execution successful")
		} else {
			logger.With(ctx, "sql", sql).Errorf("DB execution error: %v", err)
		}
	}
}
