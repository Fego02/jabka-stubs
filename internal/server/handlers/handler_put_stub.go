package handlers

import (
	"fmt"
	"github.com/Fego02/jabka-stubs/pkg/httpstubs"
	"github.com/Fego02/jabka-stubs/pkg/utils"
	"net/http"
	"strconv"
)

type StubPutHandler struct {
	Stubs *httpstubs.Stubs
}

func (h *StubPutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		ErrInvalidId.HttpError(&w)
		LogWithRequestDetails(LevelStubsOperations, r,
			"stub put: invalid id", "status", ErrInvalidId.Status)
		return
	}
	stub, err := getStubFromRequest(r)
	if err != nil {
		http.Error(w, utils.CapitalizeFirstLetter(err.Error()), http.StatusBadRequest)
		LogWithRequestDetails(LevelStubsOperations, r,
			"stub put: fail", "error", err.Error(), "status", http.StatusBadRequest)
		return
	}
	if ok := h.Stubs.Put(stub, id); !ok {
		ErrStubNotFoundById.HttpError(&w)
		LogWithRequestDetails(LevelStubsOperations, r,
			"stub put: stub not found", "status", ErrStubNotFoundById.Status)
		return
	}
	w.WriteHeader(http.StatusOK)

	response := fmt.Sprintf("Stub replaced successfully for %s on %s\n", stub.GetName(),
		stub.Request.StubRequestUrl.String())

	if err = writeResponseBody(response, &w); err != nil {
		ErrWritingResponse.HttpError(&w)
		LogWithRequestDetails(LevelStubsOperations, r,
			"stub put: cannot write response", "status", ErrWritingResponse.Status)
		return
	}

	LogWithRequestDetails(LevelStubsOperations, r,
		"stub put: success", "response", response, "status", http.StatusCreated)
}
