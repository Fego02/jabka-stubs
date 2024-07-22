package httpstubs

import (
	"net/http"
)

type StubResponse struct {
	Status  int               `json:"status"`
	Headers map[string]string `json:"headers"`
	StubResponseBody
}

func (stubResponse *StubResponse) Validate() error {
	if err := stubResponse.StubResponseBody.Validate(); err != nil {
		return err
	}
	if err := validateStatusCode(stubResponse.Status); err != nil {
		return err
	}
	if err := validateResponseHeaders(stubResponse.Headers); err != nil {
		return err
	}

	return nil
}

func (stubResponse *StubResponse) WriteToResponse(w *http.ResponseWriter) error {
	(*w).WriteHeader(stubResponse.Status)
	for key, value := range stubResponse.Headers {
		(*w).Header().Set(key, value)
	}

	body := stubResponse.StubResponseBody.Bytes()
	if body != nil {
		_, err := (*w).Write(body)
		if err != nil {
			return ErrResponseBodyWrite
		}
	}
	return nil
}

func validateResponseHeaders(headers map[string]string) error {
	for headerName := range headers {
		if headerName == "" {
			return ErrInvalidResponseHeaderName
		}
	}
	return nil
}

func validateStatusCode(statusCode int) error {
	if statusCode < 100 || statusCode > 999 {
		return ErrInvalidStatusCode
	}
	return nil
}
