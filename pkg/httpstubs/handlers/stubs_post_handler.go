package handlers

import (
	"fmt"
	"github.com/Fego02/jabka-stubs/pkg/httpstubs"
	"github.com/Fego02/jabka-stubs/pkg/utils"
	"net/http"
	"strconv"
)

type StubsPostHandler struct {
	StubsPtr *httpstubs.Stubs
}

func (h *StubsPostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	stub := httpstubs.NewStub()
	if err := stub.ReadFromRequest(r); err != nil {
		http.Error(w, utils.CapitalizeFirstLetter(err.Error()), http.StatusBadRequest)
		return
	}

	h.StubsPtr.Add(stub)
	w.Header().Set("Location", r.URL.Path+"/"+strconv.Itoa(len(h.StubsPtr.IdMap)))
	w.WriteHeader(http.StatusCreated)
	_, err := fmt.Fprintf(w, "Stub created successfully for %s on %s\n", stub.String(),
		stub.Request.StubRequestUrl.String())
	if err != nil {
		http.Error(w, "Writing Error", http.StatusInternalServerError)
	}
}
