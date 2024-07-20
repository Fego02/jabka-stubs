package handlers

import (
	httpstubs "github.com/Fego02/jabka-stubs/src/stubs/http-stubs"
	"net/http"
)

type StubsExtraHandler struct {
	StubsPtr *httpstubs.Stubs
}

func (h *StubsExtraHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
