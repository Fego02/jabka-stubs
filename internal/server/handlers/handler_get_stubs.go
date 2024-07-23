package handlers

import (
	"encoding/json"
	"github.com/Fego02/jabka-stubs/pkg/httpstubs"
	"net/http"
)

// Не совсем готово

type StubsGetHandler struct {
	Stubs *httpstubs.Stubs
}

func (h *StubsGetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	h.Stubs.Mutex.RLock()
	defer h.Stubs.Mutex.RUnlock()

	for _, stub := range h.Stubs.Map {
		if err := json.NewEncoder(w).Encode(stub); err != nil {
			ErrFailedToEncodeStub.HttpError(&w)
			LogWithRequestDetails(LevelStubsOperations, r,
				"stubs get:failed to encode", "status", ErrFailedToEncodeStub.Status)
			return
		}
	}
}
