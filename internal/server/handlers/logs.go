package handlers

import (
	"log/slog"
	"net/http"
)

var (
	LevelNonMatchedRequests = slog.Level(1)
	LevelMatchedRequests    = slog.Level(2)
	LevelStubsOperations    = slog.Level(3)
)

func LogWithRequestDetails(level slog.Level, r *http.Request, message string, args ...any) {
	requestSlice := []any{"url", r.URL.String(), "method", r.Method, "headers", r.Header, "body", r.Body}
	args = append(args, requestSlice...)
	slog.Log(nil, level, message, args...)
}
