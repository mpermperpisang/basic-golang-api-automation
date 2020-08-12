@example-api
Feature: API Automation

  @post
  Scenario: POST - Hit /unicorns endpoint
    Given endpoint "/unicorns"
    When post endpoint
    Then validate post response
