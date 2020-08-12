@example-api
Feature: API Automation

  @delete
  Scenario: DELETE - Hit /unicorns endpoint
    Given endpoint "/unicorns"
    When post endpoint
    And get id
    And delete endpoint
    Then validate delete response
