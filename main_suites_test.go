package main

import (
	"github.com/basic-golang-api-automation/features/suites"
	"github.com/cucumber/godog"
)

func GodogMainSuites(s *godog.Suite) {
	suites.AutomationAPI(s)
}
