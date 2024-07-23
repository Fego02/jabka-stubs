package httpstubs

type StubRequest struct {
	StubRequestUrl
	StubRequestMethod
	StubRequestHeaders
	StubRequestBody
}

func (stubRequest *StubRequest) Validate() error {
	if err := stubRequest.StubRequestMethod.Validate(); err != nil {
		return err
	}
	if err := stubRequest.StubRequestUrl.Validate(); err != nil {
		return err
	}
	if err := stubRequest.StubRequestBody.Validate(); err != nil {
		return err
	}

	return stubRequest.StubRequestHeaders.Validate()
}

func (stubRequest *StubRequest) Matches(r *MyRequest) bool {
	return stubRequest.StubRequestUrl.Matches(r.Url) && stubRequest.StubRequestMethod.Matches(r.Method) &&
		stubRequest.StubRequestHeaders.Matches(r.Headers) && stubRequest.StubRequestBody.Matches(r.Body)
}
