package httpstubs

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidJSON            = errors.New("invalid JSON")
	ErrInvalidContentType     = errors.New("invalid Content Type")
	ErrInvalidMultipart       = errors.New("invalid multipart")
	ErrCannotReadRequestBody  = errors.New("cannot read request body file")
	ErrCannotReadResponseBody = errors.New("cannot read response body file")
)

type StubPartNotFoundError struct {
	PartName string
}

func (e StubPartNotFoundError) Error() string {
	return fmt.Sprintf("part %s not found", e.PartName)
}

func NewStubPartNotFoundError(partName string) StubPartNotFoundError {
	return StubPartNotFoundError{PartName: partName}
}

type StubPartFoundMoreThanOnceError struct {
	PartName string
}

func (e StubPartFoundMoreThanOnceError) Error() string {
	return fmt.Sprintf("part %s found more than once", e.PartName)
}

func NewStubPartFoundMoreThanOnceError(partName string) StubPartFoundMoreThanOnceError {
	return StubPartFoundMoreThanOnceError{PartName: partName}
}

type CannotOpenStubPartError struct {
	PartName string
}

func (e CannotOpenStubPartError) Error() string {
	return fmt.Sprintf("cannot open part %s", e.PartName)
}

func NewCannotOpenStubPartError(partName string) CannotOpenStubPartError {
	return CannotOpenStubPartError{PartName: partName}
}
