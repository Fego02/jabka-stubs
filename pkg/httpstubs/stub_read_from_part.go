package httpstubs

import (
	"github.com/Fego02/jabka-stubs/pkg/utils"
	"io"
	"mime/multipart"
	"strings"
)

func (stub *Stub) ReadFromMultipart(form *multipart.Form) error {

	if err := stub.ReadStubData(form); err != nil {
		return err
	}
	if err := stub.ReadOptionalPartByKey(RequestBodyPartKey, &stub.Request.BodyBin, form); err != nil {
		return err
	}
	if err := stub.ReadOptionalPartByKey(ResponseBodyPartKey, &stub.Response.BodyBin, form); err != nil {
		return err
	}
	if err := stub.Validate(); err != nil {
		return err
	}

	return nil
}

func (stub *Stub) ReadStubData(form *multipart.Form) error {
	stubDataReader, err := ReaderFromMultiPartFormByKey(form, StubDataPartKey)
	if err != nil {
		return err
	}
	if closer, ok := stubDataReader.(io.Closer); ok {
		defer utils.HandleClose(closer)
	}
	return stub.ReadFromJson(stubDataReader)
}

func (stub *Stub) ReadOptionalPartByKey(key string, dst *[]byte, form *multipart.Form) error {
	partReader, err := ReaderFromMultiPartFormByKey(form, key)
	switch err.(type) {
	case PartNotFoundError:
	case nil:
		if closer, ok := partReader.(io.Closer); ok {
			defer utils.HandleClose(closer)
		}
		*dst, err = io.ReadAll(partReader)
		if err != nil {
			return err
		}
	default:
		return err
	}
	return nil
}

func ReaderFromMultiPartFormByKey(form *multipart.Form, key string) (io.Reader, error) {
	fileHeader, isFileFound := form.File[key]
	text, isTextFound := form.Value[key]
	if (!isFileFound || len(fileHeader) == 0) && !isTextFound {
		return nil, NewPartNotFoundError(key)
	}
	if isTextFound && isFileFound {
		return nil, NewPartFoundMoreThanOnceError(key)
	}

	if isTextFound {
		return strings.NewReader(text[0]), nil
	}

	reader, err := fileHeader[0].Open()
	if err != nil {
		return nil, NewCannotOpenStubPartError(key)
	}

	return reader, nil
}
