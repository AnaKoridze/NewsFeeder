Feature: Test unknown endpoint

  Scenario: HTTP GET on an Unknown endpoint
    Given a running api
    When I send a "GET" request to "/dummy"
    Then the response code should be 404
    And the response should match json:
      """
      {
        "status": 404,
        "error": "Not Found"
      }
      """

  Scenario: HTTP POST on an Unknown endpoint
    Given a running api
    When I send a "POST" request to "/dummy"
    Then the response code should be 404
    And the response should match json:
      """
      {
        "status": 404,
        "error": "Not Found"
      }
      """
