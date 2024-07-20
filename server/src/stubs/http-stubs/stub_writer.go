package httpstubs

import (
	"errors"
	"net/http"
)

func (stub *Stub) WriteToResponse(w *http.ResponseWriter) error {
	(*w).WriteHeader(stub.Response.Status)
	for key, value := range stub.Response.Headers {
		(*w).Header().Set(key, value)
	}

	if stub.Response.BodyBinary != nil {
		_, err := (*w).Write(stub.Response.BodyBinary)
		if err != nil {
			return errors.New("writing Error")
		}
	}

	return nil
}
