Feature: Test Recruiter Login endpoint

  Scenario: Successful Create Newsfeed action
    Given a running api
    When I send a "POST" request to "/newsfeed" with json:
    """
    {
      "title": "title",
      "post": "post",
      "author": "author"
    }
    """
    Then the response code should be 200
    And the response should match json:
    """
    {
      "id": 1
    }
    """