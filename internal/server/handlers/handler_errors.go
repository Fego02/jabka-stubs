package handlers

import (
	"fmt"
	"net/http"
)

var (
	ErrInvalidId          = NewHandlerError("Invalid Id", http.StatusBadRequest)
	ErrStubNotFoundById   = NewHandlerError("Stub not found by Id", http.StatusNotFound)
	ErrFailedToEncodeStub = NewHandlerError("Failed to Encode Stub JSON", http.StatusInternalServerError)
	ErrMethodNotAllowed   = NewHandlerError("Method not allowed", http.StatusMethodNotAllowed)
	ErrWritingResponse    = NewHandlerError("Failed to Write Response", http.StatusInternalServerError)
	ErrStubNotFound       = NewHandlerError("Stub not found", http.StatusNotFound)
	ErrMultipleStubsFound = NewHandlerError("Multiple stubs found", http.StatusForbidden)
	ErrCannotReadRequest  = NewHandlerError("Cannot Read Request", http.StatusInternalServerError)
)

type HandlerError struct {
	ErrorMessage string
	Status       int
}

func (e HandlerError) Error() string {
	return fmt.Sprintf("%s", e.ErrorMessage)
}

func (e HandlerError) HttpError(w *http.ResponseWriter) {
	http.Error(*w, e.ErrorMessage, e.Status)
}

func NewHandlerError(errorMessage string, status int) HandlerError {
	return HandlerError{ErrorMessage: errorMessage, Status: status}
}
