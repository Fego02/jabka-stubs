package server

import (
	"github.com/Fego02/jabka-stubs/internal/server/handlers"
	"github.com/Fego02/jabka-stubs/pkg/httpstubs"
	"net/http"
)

func NewRouter() *http.ServeMux {
	router := http.NewServeMux()
	stubs := httpstubs.NewStubs()
	router.Handle("POST /stubs/http-stubs", &handlers.StubsPostHandler{Stubs: stubs})
	router.Handle("GET /stubs/http-stubs/{id}", &handlers.StubGetHandler{Stubs: stubs})
	router.Handle("GET /stubs/http-stubs", &handlers.StubsGetHandler{Stubs: stubs})
	router.Handle("DELETE /stubs/http-stubs/{id}", &handlers.StubDeleteHandler{Stubs: stubs})
	router.Handle("DELETE /stubs/http-stubs", &handlers.StubsDeleteHandler{Stubs: stubs})
	router.Handle("PUT /stubs/http-stubs/{id}", &handlers.StubPutHandler{Stubs: stubs})
	router.Handle("/stubs/", &handlers.StubsExtraHandler{})
	router.Handle("/", &handlers.RequestsHandler{StubsPtr: stubs})

	return router
}
