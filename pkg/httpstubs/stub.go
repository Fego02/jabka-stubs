package httpstubs

import (
	"encoding/json"
	"io"
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
		Response:   StubResponse{},
		Properties: StubProperties{IsLoggingEnabled: true, Delay: 0},
	}
}
func (stub *Stub) String() string {
	if stub.Name != nil {
		return *stub.Name
	}
	return "anonymous stub"
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
