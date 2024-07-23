package handlers

import (
	"encoding/json"
	"github.com/Fego02/jabka-stubs/pkg/httpstubs"
	"net/http"
	"strconv"
)

type StubGetHandler struct {
	Stubs *httpstubs.Stubs
}

func (h *StubGetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		ErrInvalidId.HttpError(&w)
		LogWithRequestDetails(LevelStubsOperations, r,
			"stub get:invalid id", "status", ErrInvalidId.Status)
		return
	}

	stub := h.Stubs.GetById(id)
	if stub == nil {
		ErrStubNotFoundById.HttpError(&w)
		LogWithRequestDetails(LevelStubsOperations, r,
			"stub get:stub not found", "status", ErrStubNotFoundById.Status)
		return
	}
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(stub); err != nil {
		ErrFailedToEncodeStub.HttpError(&w)
		LogWithRequestDetails(LevelStubsOperations, r,
			"stub get:failed to encode", "status", ErrFailedToEncodeStub.Status)
		return
	}
	LogWithRequestDetails(LevelStubsOperations, r,
		"stub get:success", "status", http.StatusOK)
}
