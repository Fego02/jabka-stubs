package handlers

import (
	"github.com/Fego02/jabka-stubs/pkg/httpstubs"
	"github.com/Fego02/jabka-stubs/pkg/utils"
	"net/http"
	"time"
)

type RequestsHandler struct {
	StubsPtr *httpstubs.Stubs
}

func (h *RequestsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	startMatch := time.Now()
	stubs := h.StubsPtr.OptimizedGetMatchingStubsByRequest(r)
	matchTime := time.Since(startMatch)
	if len(stubs) == 0 {
		ErrStubNotFound.HttpError(&w)
		LogWithRequestDetails(LevelNonMatchedRequests, r,
			"request not matched any stub", "match time", matchTime, "status", ErrStubNotFound.Status)
		return
	}
	if len(stubs) != 1 {
		ErrMultipleStubsFound.HttpError(&w)
		LogWithRequestDetails(LevelNonMatchedRequests, r,
			"request matched multiple stubs", "match time", matchTime, "status", ErrMultipleStubsFound.Status)
		return
	}
	stub := stubs[0]
	if err := stub.Serve(r, &w); err != nil {
		http.Error(w, utils.CapitalizeFirstLetter(err.Error()), http.StatusBadRequest)
		if stub.Properties.IsLoggingEnabled {
			LogWithRequestDetails(LevelMatchedRequests, r,
				"request matched stub but ended in error", "error", err.Error(),
				"match time", matchTime, "status", http.StatusBadRequest)
		}

		return
	}
	if stub.Properties.IsLoggingEnabled {
		LogWithRequestDetails(LevelMatchedRequests, r,
			"request matched stub", "match time", matchTime, "stub name", stub.GetName(),
			"status", stub.Response.Status)
	}
}
