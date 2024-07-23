package httpstubs

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"
)

type Stub struct {
	Name       *string        `json:"name"`
	Request    StubRequest    `json:"request"`
	Response   StubResponse   `json:"response"`
	Properties StubProperties `json:"properties"`
}

func NewStub() *Stub {
	return &Stub{
		Request:    StubRequest{},
		Response:   StubResponse{Status: 200},
		Properties: StubProperties{IsLoggingEnabled: true, Delay: 0},
	}
}

func (stub *Stub) Matches(r *MyRequest) bool {
	return stub.Request.Matches(r)
}

func (stub *Stub) WriteToResponse(w *http.ResponseWriter) error {
	return stub.Response.WriteToResponse(w)
}

func (stub *Stub) Serve(r *http.Request, w *http.ResponseWriter) error {
	if stub.Properties.Delay != 0 {
		time.Sleep(time.Duration(stub.Properties.Delay) * time.Millisecond)
	}
	return stub.Response.WriteToResponse(w)
}

func (stub *Stub) GetName() string {
	if stub.Name != nil {
		return *stub.Name
	}
	return AnonymousStubName
}

func (stub *Stub) ReadFromRequest(r *http.Request) error {
	contentType := r.Header.Get("Content-Type")
	if strings.HasPrefix(contentType, "multipart/form-data") {
		err := r.ParseMultipartForm(MaxPartFileSize)
		if err != nil {
			return ErrInvalidMultipart
		}
		return stub.ReadFromMultipart(r.MultipartForm)
	}
	if contentType == "application/json" {
		return stub.ReadFromSimpleRequest(r)
	}

	return ErrInvalidContentType
}

func (stub *Stub) ReadFromSimpleRequest(r *http.Request) error {
	if err := stub.ReadFromJson(r.Body); err != nil {
		return err
	}
	return nil
}

func (stub *Stub) ReadFromJson(r io.Reader) error {
	err := json.NewDecoder(r).Decode(stub)
	if err != nil {
		return ErrInvalidJson
	}

	if err := stub.Validate(); err != nil {
		return err
	}

	return nil
}

func (stub *Stub) Validate() error {
	if err := stub.Request.Validate(); err != nil {
		return err
	}
	if err := stub.Response.Validate(); err != nil {
		return err
	}
	if err := stub.Properties.Validate(); err != nil {
		return err
	}
	return nil
}
