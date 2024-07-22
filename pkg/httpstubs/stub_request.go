package httpstubs

import (
	"github.com/Fego02/jabka-stubs/pkg/utils"
	"io"
	"net/http"
)

type StubRequest struct {
	StubRequestUrl
	StubRequestMethod
	StubRequestHeaders
	StubRequestBody
}

func (stubRequest *StubRequest) Validate() error {
	if err := stubRequest.StubRequestMethod.Validate(); err != nil {
		return err
	}
	if err := stubRequest.StubRequestUrl.Validate(); err != nil {
		return err
	}
	if err := stubRequest.StubRequestBody.Validate(); err != nil {
		return err
	}

	return stubRequest.StubRequestHeaders.Validate()
}

func (stubRequest *StubRequest) Matches(r *http.Request) bool {
	url := r.URL.String()
	method := r.Method
	headers := make(map[string]string)
	for headerName, headerValues := range r.Header {
		headers[headerName] = headerValues[0]
	}
	headers["Host"] = r.Host

	body, err := readRequestBody(r)
	if err != nil {
		return false
	}

	a := stubRequest.StubRequestUrl.Matches(url)
	b := stubRequest.StubRequestMethod.Matches(method)
	c := stubRequest.StubRequestHeaders.Matches(headers)
	d := stubRequest.StubRequestBody.Matches(body)
	return a && b && c && d
}

func readRequestBody(r *http.Request) ([]byte, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, ErrCannotReadHTTPRequestBody
	}
	defer utils.HandleClose(r.Body)

	return body, nil
}
