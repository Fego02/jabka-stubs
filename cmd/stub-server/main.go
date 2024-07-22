package main

import (
	"flag"
	"fmt"
	"github.com/Fego02/jabka-stubs/pkg/httpstubs"
	"github.com/Fego02/jabka-stubs/pkg/httpstubs/handlers"
	"github.com/Fego02/jabka-stubs/pkg/utils"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
)

// TODO Возможность сохранять заглушки между сессиями и быстро выставлять их
// Проверку на то, что адрес правильно
// Доделать API, хотя бы CRUD
// Продвинутый мэтчинг тела, URL и заголовков
// Журналирование
// Задержка в конце или в начале?

var programLevel = new(slog.LevelVar) // Info by default

func main() {
	port := flag.String("port", "8080", "port to listen on")
	address := flag.String("address", "127.0.0.1", "address to bind to")
	logFileName := flag.String("log", "", "path to log file or none to disable logging")
	logForFoundOnly := flag.Bool("log_for_matched_only", true, "log only for matched requests")

	slog.SetLogLoggerLevel(handlers.LevelMatchedRequests)

	flag.Parse()
	switch *logFileName {
	case "":
	case "none":
		log.SetOutput(io.Discard)
	default:
		file, err := os.OpenFile(*logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("Failed to open log file: %v", err)
		}
		defer utils.HandleClose(file)
		log.SetOutput(file)
	}

	if *logForFoundOnly {
		programLevel.Set(handlers.LevelNonMatchedRequests)
	}
	stubs := httpstubs.NewStubs()
	http.Handle("POST /stubs/http-stubs/", &handlers.StubsPostHandler{StubsPtr: stubs})
	http.Handle("POST /stubs/http-stubs", &handlers.StubsPostHandler{StubsPtr: stubs})

	http.Handle("/", &handlers.RequestsHandler{StubsPtr: stubs})
	http.Handle("GET /stubs/http-stubs/{id}", &handlers.StubGetHandler{StubsPtr: stubs})
	http.Handle("/stubs/", &handlers.StubsExtraHandler{StubsPtr: stubs})

	fmt.Printf("Server listening on %s:%s\n", *address, *port)
	log.Fatal(http.ListenAndServe(*address+":"+*port, nil))
}
