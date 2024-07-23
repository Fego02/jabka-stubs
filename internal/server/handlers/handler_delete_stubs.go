package handlers

import (
	"github.com/Fego02/jabka-stubs/pkg/httpstubs"
	"net/http"
)

type StubsDeleteHandler struct {
	Stubs *httpstubs.Stubs
}

func (h *StubsDeleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	h.Stubs.Mutex.Lock()
	defer h.Stubs.Mutex.Unlock()
	h.Stubs.Map = make(map[int]*httpstubs.Stub, 10)
	LogWithRequestDetails(LevelStubsOperations, r,
		"stubs delete:success", "status", http.StatusOK)
}
