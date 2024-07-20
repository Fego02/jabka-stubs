package httpstubs

import (
	"errors"
	"net/http"
	"net/url"
	"regexp"
)

func isValidRegex(pattern string) bool {
	_, err := regexp.Compile(pattern)
	return err == nil
}

func validateName(name string) error {
	if name == "" {
		return errors.New("invalid Name")
	}
	return nil
}

func validateURL(URL string) error {
	_, err := url.ParseRequestURI(URL)
	if err != nil {
		return errors.New("invalid URL")
	}

	return nil
}

func validateMethod(method string) error {
	switch method {
	case http.MethodGet, http.MethodHead, http.MethodPost, http.MethodPut, http.MethodPatch,
		http.MethodDelete, http.MethodConnect, http.MethodOptions, http.MethodTrace:
	default:
		return errors.New("invalid Method")
	}
	return nil
}

func validateHeaders(headers map[string]string) error {
	for headerKey := range headers {
		if headerKey == "" {
			return errors.New("invalid Headers Data")
		}
	}
	return nil
}

func validateRequestHeaders(requestHeaders map[string]string) error {
	if validateHeaders(requestHeaders) != nil {
		return errors.New("invalid Request Headers Data")
	}
	return nil
}

func validateResponseHeaders(responseHeaders map[string]string) error {
	if validateHeaders(responseHeaders) != nil {
		return errors.New("invalid Response Headers Data")
	}
	return nil
}

func validateStatusCode(statusCode int) error {
	if statusCode < 100 || statusCode > 999 {
		return errors.New("invalid Status Code")
	}
	return nil
}

func validateRequest(request *StubRequest) error {
	if request.URLMatches != "" && request.URL != "" {
		return errors.New("invalid URL Data. Only one of the URL Matches or URL Match must be provided")
	} else if request.URLMatches != "" {
		//if !isValidRegex(request.URLMatches) {
		//	return errors.New("invalid URL regexp")
		//}
	} else {
		if err := validateURL(request.URL); err != nil {
			return err
		}
	}

	if err := validateMethod(request.Method); err != nil {
		return err
	}
	//if len(request.HeadersMatches) != 0 && len(request.Headers) != 0 {
	//	return errors.New("invalid Headers Data. Only of the Headers Matches or Headers must be provided")
	//} else if len(request.HeadersMatches) != 0 {
	//	for headerKey := range request.HeadersMatches {
	//		if headerKey == "" {
	//			return errors.New("invalid Headers Data")
	//		}
	//	}
	//	return nil
	//}
	if err := validateRequestHeaders(request.Headers); err != nil {
		return err
	}

	return nil
}

func validateResponse(response *StubResponse) error {
	if err := validateResponseHeaders(response.Headers); err != nil {
		return err
	}
	if err := validateStatusCode(response.Status); err != nil {
		return err
	}

	return nil
}

func validateDelay(delay int) error {
	if delay < 0 {
		return errors.New("invalid Delay")
	}
	return nil
}

func validateProperties(properties *StubProperties) error {
	if err := validateDelay(properties.Delay); err != nil {
		return err
	}

	return nil
}
func (stub *Stub) validate() error {
	if err := validateName(stub.Name); err != nil {
		return err
	}
	if err := validateRequest(&stub.Request); err != nil {
		return err
	}
	if err := validateResponse(&stub.Response); err != nil {
		return err
	}
	if err := validateProperties(&stub.Properties); err != nil {
		return err
	}

	return nil
}
