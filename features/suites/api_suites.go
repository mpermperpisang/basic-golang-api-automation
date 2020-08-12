package suites

import (
	"github.com/basic-golang-api-automation/features/stepdefinitions"
	"github.com/cucumber/godog"
)

// AutomationAPI : suites for API steps
func AutomationAPI(s *godog.Suite) {
	s.Step(`^endpoint "([^"]*)"$`, stepdefinitions.GivenEndpoint)
	s.Step(`get id`, stepdefinitions.GetIDCRUD)
	s.Step(`get endpoint`, stepdefinitions.GetEndpoint)
	s.Step(`post endpoint`, stepdefinitions.PostEndpoint)
	s.Step(`put endpoint`, stepdefinitions.PutEndpoint)
	s.Step(`delete endpoint`, stepdefinitions.DeleteEndpoint)
	s.Step(`validate get response`, stepdefinitions.ValidateGetResponse)
	s.Step(`validate post response`, stepdefinitions.ValidatePostResponse)
	s.Step(`validate put response`, stepdefinitions.ValidatePutResponse)
	s.Step(`validate delete response`, stepdefinitions.ValidateDeleteResponse)
}
