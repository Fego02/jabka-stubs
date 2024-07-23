package handlers

import (
	"net/http"
)

type StubsExtraHandler struct {
}

func (h *StubsExtraHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ErrMethodNotAllowed.HttpError(&w)
	LogWithRequestDetails(LevelStubsOperations, r,
		"stub extra: invalid method", "status", ErrMethodNotAllowed.Status)
}
