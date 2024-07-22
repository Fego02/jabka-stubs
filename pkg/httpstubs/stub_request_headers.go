package httpstubs

import (
	"github.com/Fego02/jabka-stubs/pkg/utils"
	"regexp"
)

type StubRequestHeaders struct {
	Headers           map[string]string `json:"headers"`
	HeadersMatches    map[string]string `json:"headers_matches"`
	HeadersNotMatches map[string]string `json:"headers_not_matches"`
}

func (stubRequestHeaders *StubRequestHeaders) Matches(headers map[string]string) bool {
	for stubHeaderName, stubHeaderValue := range stubRequestHeaders.Headers {
		headerValue, ok := headers[stubHeaderName]
		if !ok {
			return false
		}

		if stubHeaderValue != headerValue {
			return false
		}
	}
	for stubHeaderName, stubHeaderValue := range stubRequestHeaders.HeadersMatches {
		headerValue, ok := headers[stubHeaderName]
		if !ok {
			return false
		}

		re, err := regexp.Compile(stubHeaderValue)
		if err != nil {
			return false
		}
		if !re.MatchString(headerValue) {
			return false
		}
	}
	for stubHeaderName, stubHeaderValue := range stubRequestHeaders.HeadersNotMatches {
		headerValue, ok := headers[stubHeaderName]
		if !ok {
			return false
		}

		re, err := regexp.Compile(stubHeaderValue)
		if err != nil {
			return false
		}
		if re.MatchString(headerValue) {
			return false
		}
	}

	return true
}

func (stubRequestHeaders *StubRequestHeaders) Validate() error {
	headerNames := make(map[string]struct{})

	for headerName := range stubRequestHeaders.Headers {
		if _, ok := headerNames[headerName]; ok {
			return ErrRequestHeadersOverloaded
		}
		headerNames[headerName] = struct{}{}
		if headerName == "" {
			return ErrInvalidRequestsHeaderName
		}
	}

	for headerName, headerValue := range stubRequestHeaders.HeadersMatches {
		if _, ok := headerNames[headerName]; ok {
			return ErrRequestHeadersOverloaded
		}
		headerNames[headerName] = struct{}{}
		if headerName == "" {
			return ErrInvalidRequestsHeaderName
		}
		if !utils.IsValidRegex(headerValue) {
			return ErrInvalidRequestsHeaderValueRegex
		}
	}

	for headerName, headerValue := range stubRequestHeaders.HeadersNotMatches {
		if _, ok := headerNames[headerName]; ok {
			return ErrRequestHeadersOverloaded
		}
		headerNames[headerName] = struct{}{}
		if headerName == "" {
			return ErrInvalidRequestsHeaderName
		}
		if !utils.IsValidRegex(headerValue) {
			return ErrInvalidRequestsHeaderValueRegex
		}
	}

	return nil
}
