package handlers

import (
	"github.com/Fego02/jabka-stubs/src/stubs/http-stubs"
	"net/http"
	"time"
)

type RequestsHandler struct {
	StubsPtr *httpstubs.Stubs
}

func (h *RequestsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	request := httpstubs.StubRequest{}
	err := request.ReadFromRequest(r)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusNotFound)
		return
	}

	stubs := h.StubsPtr.GetMatchingStubsByRequest(&request)
	if len(stubs) == 0 {
		http.Error(w, "Stub not found", http.StatusNotFound)
		return
	}
	if len(stubs) != 1 {
		http.Error(w, "Multiple stubs found", http.StatusNotFound)
		return
	}

	stub := stubs[0]
	if err = stub.WriteToResponse(&w); err != nil {
		http.Error(w, "Writing Error", http.StatusInternalServerError)
	}
	if stub.Properties.Delay != 0 {
		time.Sleep(time.Duration(stub.Properties.Delay) * time.Millisecond)
	}
}
