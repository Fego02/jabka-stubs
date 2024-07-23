package httpstubs

import (
	"github.com/Fego02/jabka-stubs/pkg/utils"
	"io"
	"net/http"
)

type MyRequest struct {
	Url     string
	Method  string
	Headers map[string]string
	Body    []byte
}

func (req *MyRequest) Read(r *http.Request) error {
	req.Url = r.URL.String()
	req.Method = r.Method
	headers := make(map[string]string)
	for headerName, headerValues := range r.Header {
		headers[headerName] = headerValues[0]
	}
	headers["Host"] = r.Host

	var err error
	req.Body, err = io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer utils.HandleClose(r.Body)

	return nil
}
