package httpstubs

import "github.com/Fego02/jabka-stubs/pkg/utils"

type StubRequestMethod struct {
	Method      *string  `json:"method"`
	MethodsList []string `json:"methods_list"`
}

func (stubRequestMethod *StubRequestMethod) Matches(method string) bool {
	if stubRequestMethod.Method != nil {
		return *stubRequestMethod.Method == method
	}
	if stubRequestMethod.MethodsList != nil {
		for _, m := range stubRequestMethod.MethodsList {
			if m == method {
				return true
			}
		}
		return false
	}

	return true
}

func (stubRequestMethod *StubRequestMethod) Validate() error {
	if utils.ContainAtLeastTwoNotNils(stubRequestMethod.Method, stubRequestMethod.MethodsList) {
		return ErrRequestMethodOverloaded
	}
	if stubRequestMethod.Method != nil {
		if !utils.IsMethodValid(*stubRequestMethod.Method) {
			return ErrInvalidRequestMethod
		}
		return nil
	}

	if stubRequestMethod.MethodsList != nil {
		for _, m := range stubRequestMethod.MethodsList {
			if !utils.IsMethodValid(m) {
				return ErrInvalidRequestMethod
			}
		}
		return nil
	}

	return nil
}
