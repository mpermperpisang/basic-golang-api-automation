@example-api
Feature: API Automation

  @put
  Scenario: PUT - Hit /unicorns endpoint
    Given endpoint "/unicorns"
    When post endpoint
    And get id
    And put endpoint
    Then validate put response
