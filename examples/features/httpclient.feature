Feature: Hit API
  example hit http api restfull

  Scenario: health check api success
    Given "bulba-api" send request to "GET /"
    Then "bulba-api" response code should be 200