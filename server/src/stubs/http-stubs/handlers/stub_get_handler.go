package handlers

import (
	"encoding/json"
	"github.com/Fego02/jabka-stubs/src/stubs/http-stubs"
	"net/http"
	"strconv"
)

type StubGetHandler struct {
	StubsPtr *httpstubs.Stubs
}

func (h *StubGetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	stub := h.StubsPtr.GetById(id)
	if stub == nil {
		http.Error(w, "Stub Not Found By ID", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(stub); err != nil {
		http.Error(w, "Failed to encode stub JSON", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}
