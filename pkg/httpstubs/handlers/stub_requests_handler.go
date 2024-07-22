package handlers

import (
	"github.com/Fego02/jabka-stubs/pkg/httpstubs"
	"net/http"
	"time"
)

type RequestsHandler struct {
	StubsPtr *httpstubs.Stubs
}

func (h *RequestsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	stubs := h.StubsPtr.GetMatchingStubsByRequest(r)
	if len(stubs) == 0 {
		http.Error(w, "Stub not found", http.StatusNotFound)
		return
	}
	if len(stubs) != 1 {
		http.Error(w, "Multiple stubs found", http.StatusNotFound)
		return
	}

	stub := stubs[0]
	if err := stub.Response.WriteToResponse(&w); err != nil {
		http.Error(w, "Writing Error", http.StatusInternalServerError)
	}
	if stub.Properties.Delay != 0 {
		time.Sleep(time.Duration(stub.Properties.Delay) * time.Millisecond)
	}
}
