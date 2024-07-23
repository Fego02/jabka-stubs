package handlers

import (
	"github.com/Fego02/jabka-stubs/pkg/httpstubs"
	"net/http"
	"strconv"
)

type StubDeleteHandler struct {
	Stubs *httpstubs.Stubs
}

func (h *StubDeleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		ErrInvalidId.HttpError(&w)
		LogWithRequestDetails(LevelStubsOperations, r,
			"stub delete:failed to encode", "status", ErrInvalidId.Status)
		return
	}

	if ok := h.Stubs.Delete(id); !ok {
		ErrStubNotFoundById.HttpError(&w)
		LogWithRequestDetails(LevelStubsOperations, r,
			"stub delete:stub not found", "status", ErrStubNotFoundById.Status)
		return
	}

	w.WriteHeader(http.StatusOK)
	LogWithRequestDetails(LevelStubsOperations, r,
		"stub delete:success", "status", http.StatusOK)
}
