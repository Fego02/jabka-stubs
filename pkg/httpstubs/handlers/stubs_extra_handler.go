package handlers

import (
	"context"
	"github.com/Fego02/jabka-stubs/pkg/httpstubs"
	"log/slog"
	"net/http"
)

type StubsExtraHandler struct {
	StubsPtr *httpstubs.Stubs
}

func (h *StubsExtraHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	slog.Log(nil, LevelMatchedRequests, "foo")
	slog.Info("fee")
	slog.Log(context.Background(), LevelNonMatchedRequests, "boo")
}
