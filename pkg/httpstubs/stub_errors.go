package httpstubs

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidJson             = errors.New("invalid JSON")
	ErrInvalidContentType      = errors.New("invalid Content Type")
	ErrInvalidMultipart        = errors.New("invalid multipart")
	ErrCannotReadRequestBody   = errors.New("cannot read request body file")
	ErrCannotReadResponseBody  = errors.New("cannot read response body file")
	ErrRequestMethodOverloaded = errors.New("request method is overloaded. " +
		"Only one of 'method' and 'methods_list' can be provided")
	ErrInvalidRequestMethod = errors.New("invalid method")
	ErrRequestUrlOverloaded = errors.New("request url is overloaded. " +
		"Only one of 'url', 'url_matches' and 'url_not_matches' can be provided")
	ErrInvalidRequestUrl           = errors.New("invalid url")
	ErrInvalidRequestUrlMatches    = errors.New("invalid url_matches regex")
	ErrInvalidRequestUrlNotMatches = errors.New("invalid url_not_matches regex")
	ErrRequestBodyOverloaded       = errors.New("request body is overloaded. " +
		"Only one of 'body', 'body_matches', 'body_bin', 'body_not_matches', 'body_bin_matches' " +
		"and 'body_bin_not_matches' can be provided")
	ErrInvalidRequestBodyMatches       = errors.New("invalid body_matches")
	ErrInvalidRequestBodyNotMatches    = errors.New("invalid body_not_matches")
	ErrInvalidRequestBodyBinMatches    = errors.New("invalid body_bin_matches")
	ErrInvalidRequestBodyBinNotMatches = errors.New("invalid body_bin_not_matches")
	ErrRequestHeadersOverloaded        = errors.New("request headers are overloaded. " +
		"Header name can be provided only once in 'headers', 'headers_matches' and 'headers_not_matches'")
	ErrInvalidRequestsHeaderName       = errors.New("invalid request header name")
	ErrInvalidRequestsHeaderValueRegex = errors.New("invalid request header value regex")
	ErrCannotReadHTTPRequestBody       = errors.New("cannot read request body file")
	ErrInvalidResponseHeaderName       = errors.New("invalid response header name")
	ErrInvalidStatusCode               = errors.New("invalid status code")
	ErrResponseBodyOverloaded          = errors.New("response body is overloaded. " +
		"Only one of 'body' and  'body_bin' can be provided")
	ErrResponseBodyWrite = errors.New("cannot write response body to response writer")
)

type PartNotFoundError struct {
	PartName string
}

func (e PartNotFoundError) Error() string {
	return fmt.Sprintf("part %s not found", e.PartName)
}

func NewPartNotFoundError(partName string) PartNotFoundError {
	return PartNotFoundError{PartName: partName}
}

type PartFoundMoreThanOnceError struct {
	PartName string
}

func (e PartFoundMoreThanOnceError) Error() string {
	return fmt.Sprintf("part %s found more than once", e.PartName)
}

func NewPartFoundMoreThanOnceError(partName string) PartFoundMoreThanOnceError {
	return PartFoundMoreThanOnceError{PartName: partName}
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
