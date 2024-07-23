package handlers

import (
	"fmt"
	"github.com/Fego02/jabka-stubs/pkg/httpstubs"
	"github.com/Fego02/jabka-stubs/pkg/utils"
	"net/http"
	"strconv"
)

type StubsPostHandler struct {
	Stubs *httpstubs.Stubs
}

func (h *StubsPostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	stub, err := getStubFromRequest(r)
	if err != nil {
		http.Error(w, utils.CapitalizeFirstLetter(err.Error()), http.StatusBadRequest)
		LogWithRequestDetails(LevelStubsOperations, r,
			"stubs post: fail", "error", err.Error(), "status", http.StatusBadRequest)
		return
	}
	id := h.Stubs.Add(stub)
	location := r.URL.Path + "/" + strconv.Itoa(id)
	w.Header().Set("Location", location)
	w.WriteHeader(http.StatusCreated)

	response := fmt.Sprintf("Stub created successfully for %s on %s\nLocation: %s\n", stub.GetName(),
		stub.Request.StubRequestUrl.String(), location)

	if err = writeResponseBody(response, &w); err != nil {
		ErrWritingResponse.HttpError(&w)
		LogWithRequestDetails(LevelStubsOperations, r,
			"stubs post: cannot write response body", "status", ErrWritingResponse.Status)
		return
	}
	LogWithRequestDetails(LevelStubsOperations, r,
		"stubs post: success", "response", response, "status", http.StatusCreated)
}

func getStubFromRequest(r *http.Request) (*httpstubs.Stub, error) {
	stub := httpstubs.NewStub()
	if err := stub.ReadFromRequest(r); err != nil {
		return nil, err
	}
	return stub, nil
}

func writeResponseBody(s string, w *http.ResponseWriter) error {
	_, err := fmt.Fprintf(*w, s)
	return err
}
