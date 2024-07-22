package httpstubs

import "errors"

type StubProperties struct {
	IsLoggingEnabled bool `json:"is_logging_enabled"`
	Delay            int  `json:"delay"`
	Priority         int  `json:"priority"`
}

func (stubProperties *StubProperties) Validate() error {
	if err := validateDelay(stubProperties.Delay); err != nil {
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
