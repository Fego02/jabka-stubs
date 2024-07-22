package httpstubs

import (
	"bytes"
	"encoding/hex"
	"github.com/Fego02/jabka-stubs/pkg/utils"
	"regexp"
)

type StubRequestBody struct {
	BodyText           *string `json:"body"`
	BodyTextMatches    *string `json:"body_matches"`
	BodyTextNotMatches *string `json:"body_not_matches"`
	BodyBin            []byte  `json:"body_bin"`
	BodyHexMatches     *string `json:"body_hex_matches"`
	BodyHexNotMatches  *string `json:"body_hex_not_matches"`
}

func (stubRequestBody *StubRequestBody) Matches(body []byte) bool {
	if utils.IsValidUtf8String(body) {
		bodyText := string(body)

		if stubRequestBody.BodyText != nil {
			return bodyText == *stubRequestBody.BodyText
		}

		if stubRequestBody.BodyTextMatches != nil {
			re, err := regexp.Compile(*stubRequestBody.BodyTextMatches)
			if err != nil {
				return false
			}
			return re.MatchString(bodyText)
		}

		if stubRequestBody.BodyTextNotMatches != nil {
			re, err := regexp.Compile(*stubRequestBody.BodyTextNotMatches)
			if err != nil {
				return false
			}
			return !re.MatchString(bodyText)
		}
	}
	if stubRequestBody.BodyBin != nil {
		return bytes.Equal(body, stubRequestBody.BodyBin)
	}

	bodyHex := hex.EncodeToString(body)
	if stubRequestBody.BodyHexMatches != nil {
		re, err := regexp.Compile(*stubRequestBody.BodyHexMatches)
		if err != nil {
			return false
		}
		return re.MatchString(bodyHex)
	}

	if stubRequestBody.BodyHexNotMatches != nil {
		re, err := regexp.Compile(*stubRequestBody.BodyHexNotMatches)
		if err != nil {
			return false
		}
		return !re.MatchString(bodyHex)
	}

	return true
}

func (stubRequestBody *StubRequestBody) Validate() error {
	if utils.ContainAtLeastTwoNotNils(stubRequestBody.BodyText, stubRequestBody.BodyTextMatches, stubRequestBody.BodyTextNotMatches,
		stubRequestBody.BodyBin, stubRequestBody.BodyHexNotMatches, stubRequestBody.BodyHexMatches) {
		return ErrRequestBodyOverloaded
	}

	if stubRequestBody.BodyTextMatches != nil {
		if !utils.IsValidRegex(*stubRequestBody.BodyTextMatches) {
			return ErrInvalidRequestBodyMatches
		}
		return nil
	}

	if stubRequestBody.BodyTextNotMatches != nil {
		if !utils.IsValidRegex(*stubRequestBody.BodyTextNotMatches) {
			return ErrInvalidRequestBodyNotMatches
		}
		return nil
	}

	if stubRequestBody.BodyHexMatches != nil {
		if !utils.IsValidRegex(*stubRequestBody.BodyHexMatches) {
			return ErrInvalidRequestBodyBinMatches
		}
		return nil
	}

	if stubRequestBody.BodyHexNotMatches != nil {
		if !utils.IsValidRegex(*stubRequestBody.BodyHexNotMatches) {
			return ErrInvalidRequestBodyBinNotMatches
		}
		return nil
	}

	return nil
}
