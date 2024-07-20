package httpstubs

import (
	"encoding/json"
	"io"
)

type Stub struct {
	Name       string         `json:"name"`
	Request    StubRequest    `json:"request"`
	Response   StubResponse   `json:"response"`
	Properties StubProperties `json:"properties"`
}

type StubRequest struct {
	Method         string            `json:"method"`
	URL            string            `json:"url"`
	BodyText       string            `json:"body"`
	Headers        map[string]string `json:"headers"`
	BodyBinary     []byte
	URLMatches     string            `json:"url_matches"`
	BodyMatches    string            `json:"body_matches"`
	BodyContains   string            `json:"body_contains"`
	HeadersMatches map[string]string `json:"headers_matches"`
}

type StubResponse struct {
	Status     int               `json:"status"`
	BodyText   string            `json:"body"`
	Headers    map[string]string `json:"headers"`
	BodyBinary []byte
}

type StubProperties struct {
	IsLoggingEnabled bool `json:"is_logging_enabled"`
	Delay            int  `json:"delay"`
}

const maxFileSize = 10 << 20

func NewStub() *Stub {
	return &Stub{
		Request:    StubRequest{BodyText: "null"},
		Response:   StubResponse{BodyText: "null"},
		Properties: StubProperties{IsLoggingEnabled: true, Delay: 0},
	}
}

func (stub *Stub) ShallowCopy(stubSource *Stub) {
	stub.Name = stubSource.Name
	stub.Request = stubSource.Request
	stub.Response = stubSource.Response
	stub.Properties = stubSource.Properties
}

func (stub *Stub) ReadFromJson(r io.Reader) error {
	tmpStub := NewStub()
	err := json.NewDecoder(r).Decode(tmpStub)
	if err != nil {
		return ErrInvalidJSON
	}

	if err := tmpStub.validate(); err != nil {
		return err
	}
	stub.ShallowCopy(tmpStub)

	if stub.Request.BodyText != "null" {
		stub.Request.BodyBinary = []byte(stub.Request.BodyText)
	}
	if stub.Response.BodyText != "null" {
		stub.Response.BodyBinary = []byte(stub.Response.BodyText)
	}

	return nil
}
