package httpstubs

import "github.com/Fego02/jabka-stubs/pkg/utils"

type StubResponseBody struct {
	BodyText *string `json:"body"`
	BodyBin  []byte  `json:"body_bin"`
}

func (stubResponseBody *StubResponseBody) Validate() error {
	if utils.ContainAtLeastTwoNotNils(stubResponseBody.BodyBin, stubResponseBody.BodyText) {
		return ErrResponseBodyOverloaded
	}

	return nil
}

func (stubResponseBody *StubResponseBody) Bytes() []byte {
	if stubResponseBody.BodyText != nil {
		return []byte(*stubResponseBody.BodyText)
	}
	return stubResponseBody.BodyBin
}
