package album

import (
	"net/http"
	"testing"
	"time"
	"veterinaria-server/internal/auth"
	"veterinaria-server/internal/entity"
	"veterinaria-server/internal/test"
	"veterinaria-server/pkg/log"
)

func TestAPI(t *testing.T) {
	logger, _ := log.NewForTest()
	router := test.MockRouter(logger)
	repo := &mockRepository{items: []entity.Album{
		{ID: "123", Name: "album123", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}}
	RegisterHandlers(router.Group(""), NewService(repo, logger), auth.MockAuthHandler, logger)
	header := auth.MockAuthHeader()

	tests := []test.APITestCase{
		{Name: "get all", Method: "GET", URL: "/albums", Body: "", Header: nil, WantStatus: http.StatusOK, WantResponse: `*"total_count":1*`},
		{Name: "get 123", Method: "GET", URL: "/albums/123", Body: "", Header: nil, WantStatus: http.StatusOK, WantResponse: `*album123*`},
		/*{Name: "get unknown", "GET", "/albums/1234", "", nil, http.StatusNotFound, ""},
		{Name: "create ok", "POST", "/albums", `{"name":"test"}`, header, http.StatusCreated, "*test*"},
		{Name: "create ok count", "GET", "/albums", "", nil, http.StatusOK, `*"total_count":2*`},
		{Name: "create auth error", "POST", "/albums", `{"name":"test"}`, nil, http.StatusUnauthorized, ""},
		{Name: "create input error", "POST", "/albums", `"name":"test"}`, header, http.StatusBadRequest, ""},
		{Name: "update ok", "PUT", "/albums/123", `{"name":"albumxyz"}`, header, http.StatusOK, "*albumxyz*"},
		{Name: "update verify", "GET", "/albums/123", "", nil, http.StatusOK, `*albumxyz*`},
		{Name: "update auth error", "PUT", "/albums/123", `{"name":"albumxyz"}`, nil, http.StatusUnauthorized, ""},*/
		{Name: "update input error", Method: "PUT", URL: "/albums/123", Body: `"name":"albumxyz"}`, Header: header, WantStatus: http.StatusBadRequest, WantResponse: ""},
		/*{Name: "delete ok", "DELETE", "/albums/123", ``, header, http.StatusOK, "*albumxyz*"},
		{Name: "delete verify", "DELETE", "/albums/123", ``, header, http.StatusNotFound, ""},
		{Name: "delete auth error", "DELETE", "/albums/123", ``, nil, http.StatusUnauthorized, ""},*/
	}
	for _, tc := range tests {
		test.Endpoint(t, router, tc)
	}
}
