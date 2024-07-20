package main

import (
	"flag"
	"fmt"
	"github.com/Fego02/jabka-stubs/src/stubs/http-stubs"
	"github.com/Fego02/jabka-stubs/src/stubs/http-stubs/handlers"
	"log"
	"net/http"
)

// TODO Возможность сохранять заглушки между сессиями и быстро выставлять
// Проверку на то, что адрес правильно вставили

func main() {
	port := flag.String("port", "8080", "port to listen on")
	address := flag.String("address", "127.0.0.1", "address to bind to")
	flag.Parse()
	stubs := httpstubs.NewStubs()
	http.Handle("POST /stubs/http-stubs", &handlers.StubsPostHandler{StubsPtr: stubs})
	http.Handle("/", &handlers.RequestsHandler{StubsPtr: stubs})
	http.Handle("GET /stubs/http-stubs/{id}", &handlers.StubGetHandler{StubsPtr: stubs})
	http.Handle("/stubs/", &handlers.StubsExtraHandler{StubsPtr: stubs})

	fmt.Printf("Server listening on %s:%s", *address, *port)
	log.Fatal(http.ListenAndServe(*address+":"+*port, nil))
}
