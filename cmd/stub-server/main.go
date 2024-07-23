package main

import (
	"flag"
	"fmt"
	"github.com/Fego02/jabka-stubs/internal/server"
	"github.com/Fego02/jabka-stubs/internal/server/handlers"
	"github.com/Fego02/jabka-stubs/pkg/utils"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
)

var LogLevel = new(slog.LevelVar)

func main() {
	port := flag.String("port", "8080", "port to listen on")
	address := flag.String("address", "127.0.0.1", "address to bind to")
	logFileName := flag.String("log", "", "path to log file or none to disable logging")
	logForMatchedOnly := flag.Bool("log_for_matched_only", true, "log only for matched requests")
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
	if *logForMatchedOnly {
		LogLevel.Set(handlers.LevelMatchedRequests)
	} else {
		LogLevel.Set(handlers.LevelNonMatchedRequests)
	}
	slog.SetLogLoggerLevel(LogLevel.Level())

	router := server.NewRouter()
	fmt.Printf("Server listening on %s:%s\n", *address, *port)
	log.Fatal(http.ListenAndServe(*address+":"+*port, router))
}
