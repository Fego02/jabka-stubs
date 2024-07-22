package httpstubs

import (
	"github.com/Fego02/jabka-stubs/pkg/utils"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
)

// Исправить то, что при ошибке структура портится

func (stub *Stub) ReadFromRequest(r *http.Request) error {
	contentType := r.Header.Get("Content-Type")
	if strings.HasPrefix(contentType, "multipart/form-data") {
		return stub.readFromComplexRequest(r)
	}
	if contentType == "application/json" {
		return stub.readFromSimpleRequest(r)
	}

	return ErrInvalidContentType
}

func (stub *Stub) readFromSimpleRequest(r *http.Request) error {
	if err := stub.ReadFromJson(r.Body); err != nil {
		return err
	}
	return nil
}

// Надо переписать, но уже ладно

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

func (stub *Stub) readFromComplexRequest(r *http.Request) error {
	err := r.ParseMultipartForm(MaxPartFileSize)

	if err != nil {
		return ErrInvalidMultipart
	}

	stubDataReader, err := ReaderFromMultiPartFormByKey(r.MultipartForm, "stub-data")
	if err != nil {
		return err
	}
	if closer, ok := stubDataReader.(io.Closer); ok {
		defer utils.HandleClose(closer)
	}

	if err = stub.ReadFromJson(stubDataReader); err != nil {
		return err
	}

	requestBodyReader, err := ReaderFromMultiPartFormByKey(r.MultipartForm, "request-body")
	switch err.(type) {
	case PartNotFoundError:
	case nil:
		if closer, ok := requestBodyReader.(io.Closer); ok {
			defer utils.HandleClose(closer)
		}
		stub.Request.BodyBin, err = io.ReadAll(requestBodyReader)
		if err != nil {
			return ErrCannotReadRequestBody
		}
	default:
		return err
	}

	responseBodyReader, err := ReaderFromMultiPartFormByKey(r.MultipartForm, "response-body")
	switch err.(type) {
	case PartNotFoundError:
	case nil:
		if err != nil {
			return err
		}
		if closer, ok := responseBodyReader.(io.Closer); ok {
			defer utils.HandleClose(closer)
		}

		stub.Response.BodyBin, err = io.ReadAll(responseBodyReader)
		if err != nil {
			return ErrCannotReadResponseBody
		}
	default:
		return err
	}

	return nil
}
