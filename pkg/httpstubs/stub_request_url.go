package httpstubs

import (
	"fmt"
	"github.com/Fego02/jabka-stubs/pkg/utils"
	"regexp"
)

type StubRequestUrl struct {
	Url           *string `json:"url"`
	UrlMatches    *string `json:"url_matches"`
	UrlNotMatches *string `json:"url_not_matches"`
}

func (stubRequestUrl *StubRequestUrl) Matches(url string) bool {
	if stubRequestUrl.Url != nil {
		return *stubRequestUrl.Url == url
	}
	if stubRequestUrl.UrlMatches != nil {
		re, err := regexp.Compile(*stubRequestUrl.UrlMatches)
		if err != nil {
			return false
		}
		return re.MatchString(url)
	}

	if stubRequestUrl.UrlNotMatches != nil {
		re, err := regexp.Compile(*stubRequestUrl.UrlNotMatches)
		if err != nil {
			return false
		}
		return !re.MatchString(url)
	}

	return true
}

func (stubRequestUrl *StubRequestUrl) Validate() error {
	if utils.ContainAtLeastTwoNotNils(stubRequestUrl.Url, stubRequestUrl.UrlMatches, stubRequestUrl.UrlNotMatches) {
		return ErrRequestUrlOverloaded
	}

	if stubRequestUrl.Url != nil {
		if !utils.IsUrlValid(*stubRequestUrl.Url) {
			return ErrInvalidRequestUrl
		}
		return nil
	}

	if stubRequestUrl.UrlMatches != nil {
		if !utils.IsValidRegex(*stubRequestUrl.UrlMatches) {
			return ErrInvalidRequestUrlMatches
		}
		return nil
	}

	if stubRequestUrl.UrlNotMatches != nil {
		if !utils.IsValidRegex(*stubRequestUrl.UrlNotMatches) {
			return ErrInvalidRequestUrlNotMatches
		}
		return nil
	}

	return nil
}

func (stubRequestUrl *StubRequestUrl) String() string {
	if stubRequestUrl.Url != nil {
		return *stubRequestUrl.Url
	}
	if stubRequestUrl.UrlMatches != nil {
		return fmt.Sprintf("url matching regex: %s", *stubRequestUrl.UrlMatches)
	}
	if stubRequestUrl.UrlNotMatches != nil {
		return fmt.Sprintf("url not matching regex: %s", *stubRequestUrl.UrlNotMatches)
	}
	return AnyUrlString
}
