Feature: Test health endpoint

  Scenario: Health-Check
    Given a running api
    When I send a "GET" request to "/health"
    Then the response code should be 200
    And the response should match json:
      """
      { "status": "UP" }
      """
