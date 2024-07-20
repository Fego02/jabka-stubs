package httpstubs

import (
	"errors"
	"github.com/Fego02/jabka-stubs/src/stubs/utils"
	"io"
	"net/http"
)

func readRequestBody(r *http.Request) ([]byte, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, errors.New("failed to read body")
	}
	defer utils.HandleClose(r.Body)

	return body, nil
}

// Опять исправить случай при ошибке, чтобы не портить структуру

func (request *StubRequest) ReadFromRequest(r *http.Request) error {
	request.Method = r.Method
	request.URL = r.URL.String()
	request.Headers = make(map[string]string)

	for key, values := range r.Header {
		request.Headers[key] = values[0]
	}
	request.Headers["Host"] = r.Host

	body, err := readRequestBody(r)
	if err != nil {
		return err
	}

	request.BodyBinary = body

	return nil
}
