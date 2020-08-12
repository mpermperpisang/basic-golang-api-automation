@example-api @get
Feature: API Automation

  @get-1
  Scenario: GET - Hit empty endpoint
    Given endpoint "/"
    When get endpoint
    Then validate get response

  @get-2
  Scenario: GET - Hit /unicorns endpoint
    Given endpoint "/unicorns"
    When get endpoint
    Then validate get response
