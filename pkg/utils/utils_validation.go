package utils

import (
	"net/http"
	urlPkg "net/url"
	"reflect"
	"regexp"
	"unicode/utf8"
)

func IsUrlValid(url string) bool {
	_, err := urlPkg.ParseRequestURI(url)
	return err == nil
}

func IsMethodValid(method string) bool {
	switch method {
	case http.MethodGet, http.MethodHead, http.MethodPost, http.MethodPut, http.MethodPatch,
		http.MethodDelete, http.MethodConnect, http.MethodOptions, http.MethodTrace:
		return true
	default:
		return false
	}
}

func IsValidRegex(pattern string) bool {
	_, err := regexp.Compile(pattern)
	return err == nil
}

func IsValidUtf8String(data []byte) bool {
	return utf8.Valid(data)
}

func ContainAtLeastTwoNotNils(args ...interface{}) bool {
	notNilIsFound := false
	for _, arg := range args {
		if !isNil(arg) {
			if notNilIsFound {
				return true
			}
			notNilIsFound = true
		}
	}

	return false
}

func isNil(value interface{}) bool {
	v := reflect.ValueOf(value)

	if !v.IsValid() || (v.Kind() == reflect.Ptr && v.IsNil()) ||
		(v.Kind() == reflect.Slice && v.IsNil()) ||
		(v.Kind() == reflect.Map && v.IsNil()) ||
		(v.Kind() == reflect.Chan && v.IsNil()) ||
		(v.Kind() == reflect.Interface && v.IsNil()) {
		return true
	}

	return false
}
