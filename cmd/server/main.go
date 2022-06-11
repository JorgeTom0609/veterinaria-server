package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"time"

	"veterinaria-server/internal/accesos"
	"veterinaria-server/internal/album"
	"veterinaria-server/internal/auth"
	"veterinaria-server/internal/cita_medica"
	"veterinaria-server/internal/clientes"
	"veterinaria-server/internal/compra"
	"veterinaria-server/internal/config"
	"veterinaria-server/internal/consultas"
	"veterinaria-server/internal/detalle_compra"
	"veterinaria-server/internal/detalle_examen_cualitativo"
	"veterinaria-server/internal/detalle_examen_cuantitativo"
	"veterinaria-server/internal/detalle_examen_informativo"
	"veterinaria-server/internal/detalle_factura"
	"veterinaria-server/internal/detalle_hospitalizacion"
	"veterinaria-server/internal/detalle_servicio_consulta"
	"veterinaria-server/internal/detalle_servicio_hospitalizacion"
	"veterinaria-server/internal/detalle_uso_servicio"
	"veterinaria-server/internal/detalle_uso_servicio_consulta"
	"veterinaria-server/internal/documento_mascota"
	"veterinaria-server/internal/errors"
	"veterinaria-server/internal/especies"
	"veterinaria-server/internal/examen_mascota"
	"veterinaria-server/internal/factura"
	"veterinaria-server/internal/generos"
	"veterinaria-server/internal/healthcheck"
	"veterinaria-server/internal/hospitalizacion"
	"veterinaria-server/internal/lote"
	"veterinaria-server/internal/mascotas"
	"veterinaria-server/internal/medida"
	"veterinaria-server/internal/productos"
	"veterinaria-server/internal/proveedor"
	"veterinaria-server/internal/proveedor_producto"
	"veterinaria-server/internal/receta"
	"veterinaria-server/internal/rol"
	"veterinaria-server/internal/servicio_producto"
	"veterinaria-server/internal/servicios"
	"veterinaria-server/internal/stock_individual"
	"veterinaria-server/internal/tipo_examen"
	"veterinaria-server/internal/unidad"
	"veterinaria-server/internal/usuario_rol"
	"veterinaria-server/internal/usuarios"
	"veterinaria-server/pkg/accesslog"
	"veterinaria-server/pkg/dbcontext"
	"veterinaria-server/pkg/log"

	dbx "github.com/go-ozzo/ozzo-dbx"
	routing "github.com/go-ozzo/ozzo-routing/v2"
	"github.com/go-ozzo/ozzo-routing/v2/content"
	"github.com/go-ozzo/ozzo-routing/v2/cors"
	"github.com/go-ozzo/ozzo-routing/v2/file"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mileusna/crontab"
	"google.golang.org/protobuf/proto"

	_ "github.com/mattn/go-sqlite3"

	"github.com/mdp/qrterminal"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types"
	waLog "go.mau.fi/whatsmeow/util/log"
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
		authHandler, logger, db,
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

	detalle_factura.RegisterHandlers(rg.Group(""),
		detalle_factura.NewService(detalle_factura.NewRepository(db, logger), logger),
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

	servicios.RegisterHandlers(rg.Group(""),
		servicios.NewService(servicios.NewRepository(db, logger), logger),
		authHandler, logger, db,
	)

	servicio_producto.RegisterHandlers(rg.Group(""),
		servicio_producto.NewService(servicio_producto.NewRepository(db, logger), logger),
		authHandler, logger,
	)

	documento_mascota.RegisterHandlers(rg.Group(""),
		documento_mascota.NewService(documento_mascota.NewRepository(db, logger), logger),
		authHandler, logger,
	)

	hospitalizacion.RegisterHandlers(rg.Group(""),
		hospitalizacion.NewService(hospitalizacion.NewRepository(db, logger), logger),
		authHandler, logger, db,
	)

	detalle_hospitalizacion.RegisterHandlers(rg.Group(""),
		detalle_hospitalizacion.NewService(detalle_hospitalizacion.NewRepository(db, logger), logger),
		authHandler, logger,
	)

	detalle_servicio_hospitalizacion.RegisterHandlers(rg.Group(""),
		detalle_servicio_hospitalizacion.NewService(detalle_servicio_hospitalizacion.NewRepository(db, logger), logger),
		authHandler, logger, db,
	)

	detalle_uso_servicio.RegisterHandlers(rg.Group(""),
		detalle_uso_servicio.NewService(detalle_uso_servicio.NewRepository(db, logger), logger),
		authHandler, logger,
	)

	usuarios.RegisterHandlers(rg.Group(""),
		usuarios.NewService(usuarios.NewRepository(db, logger), logger),
		authHandler, logger,
	)

	rol.RegisterHandlers(rg.Group(""),
		rol.NewService(rol.NewRepository(db, logger), logger),
		authHandler, logger,
	)

	usuario_rol.RegisterHandlers(rg.Group(""),
		usuario_rol.NewService(usuario_rol.NewRepository(db, logger), logger),
		authHandler, logger,
	)

	receta.RegisterHandlers(rg.Group(""),
		receta.NewService(receta.NewRepository(db, logger), logger),
		authHandler, logger,
	)

	detalle_servicio_consulta.RegisterHandlers(rg.Group(""),
		detalle_servicio_consulta.NewService(detalle_servicio_consulta.NewRepository(db, logger), logger),
		authHandler, logger, db,
	)

	detalle_uso_servicio_consulta.RegisterHandlers(rg.Group(""),
		detalle_uso_servicio_consulta.NewService(detalle_uso_servicio_consulta.NewRepository(db, logger), logger),
		authHandler, logger,
	)

	cita_medica.RegisterHandlers(rg.Group(""),
		cita_medica.NewService(cita_medica.NewRepository(db, logger), logger),
		authHandler, logger,
	)

	auth.RegisterHandlers(rg.Group(""),
		auth.NewService(db, cfg.JWTSigningKey, cfg.JWTExpiration, logger),
		logger,
	)

	// Serving Static Files
	if runtime.GOOS == "windows" {
		rg.Get("/files/*", file.Server(file.PathMap{
			"/v1/files":                   "/resources/",
			"/v1/files/documentosMascota": "/documentos-mascota",
		}))
	} else {
		rg.Get("/files/*", file.Server(file.PathMap{
			"/v1/files":                   "/root/go/src/github.com/JorgeTom0609/veterinaria-server/resources/",
			"/v1/files/documentosMascota": "/root/go/src/github.com/JorgeTom0609/veterinaria-server/documentos-mascota",
		}))
	}

	wac, err := WAConnect()
	if err != nil {
		fmt.Println(err)
	}

	cron := crontab.New()

	//err = cron.AddJob("*/5 * * * *", func() {
	err = cron.AddJob("00 10 * * *", func() {
		wac, err = WAConnect()
		if err != nil {
			fmt.Println(err)
		}
		ctx := context.Background()
		//Buscar Citas Sin notificar
		scm := cita_medica.NewService(cita_medica.NewRepository(db, logger), logger)
		citas, err1 := scm.GetCitasMedicaSinNotificar(ctx)

		if err1 != nil {
			return
		}

		for i := 0; i < len(citas); i++ {
			_, err = wac.SendMessage(types.JID{
				User:   citas[i].Telefono,
				Server: types.DefaultUserServer,
			}, "", &waProto.Message{
				Conversation: proto.String("Saludos " + citas[i].Duenio +
					", veterinaria DELFICAR le informa que el día *" + Format(citas[i].Fecha) +
					"* tiene agendada una cita médica para su mascota *" + citas[i].Mascota +
					"* por el siguiente motivo: *" + citas[i].Motivo + "*."),
			})
			if err != nil {
				fmt.Println(err)
			} else {
				_, _ = scm.ActualizarCitaMedica(ctx, cita_medica.UpdateCitaMedicaRequest{
					IdCitaMedica:       citas[i].IdCitaMedica,
					IdMascota:          citas[i].IdMascota,
					Motivo:             citas[i].Motivo,
					Fecha:              citas[i].Fecha,
					EstadoNotificacion: "SI",
				})
			}
		}
	})

	if err != nil {
		fmt.Println(err)
		return nil
	}

	err = cron.AddJob("00 08 * * *", func() {
		wac, err = WAConnect()
		if err != nil {
			fmt.Println(err)
		}

		_, err = wac.SendMessage(types.JID{
			User:   "593960270781",
			Server: types.DefaultUserServer,
		}, "", &waProto.Message{
			Conversation: proto.String("Whatsapp funcionando"),
		})
		if err != nil {
			fmt.Println(err)
		}
	})

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return router
}

func Format(t time.Time) string {
	//days[t.Weekday()][:3], t.Day(), months[t.Month()-1][:3],
	return fmt.Sprintf("%s %02d de %s a las %02d:%02d",
		days[t.Weekday()], t.Day(), months[t.Month()-1], t.Hour(), t.Minute(),
	)
}

var days = [...]string{
	"Domingo", "Lunes", "Martes", "Miércoles", "Jueves", "Viernes", "Sábado"}

var months = [...]string{
	"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio",
	"Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre",
}

func WAConnect() (*whatsmeow.Client, error) {
	container, err := sqlstore.New("sqlite3", "file:wapp.db?_foreign_keys=on", waLog.Noop)
	if err != nil {
		return nil, err
	}
	deviceStore, err := container.GetFirstDevice()
	if err != nil {
		panic(err)
	}
	client := whatsmeow.NewClient(deviceStore, waLog.Noop)
	if client.Store.ID == nil {
		// No ID stored, new login
		qrChan, _ := client.GetQRChannel(context.Background())
		err = client.Connect()
		if err != nil {
			return nil, err
		}
		for evt := range qrChan {
			if evt.Event == "code" {
				qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
			} else {
				fmt.Println("Login event:", evt.Event)
			}
		}
	} else {
		err := client.Connect()
		if err != nil {
			return nil, err
		}
	}
	return client, nil
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
