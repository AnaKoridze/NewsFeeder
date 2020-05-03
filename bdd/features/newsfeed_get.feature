Feature: Test Resume Contact endpoint

  Scenario: Successful Get Newsfeeds action
    Given a running api
    When I send a "GET" request to "/newsfeeds"
    Then the response code should be 200
    And the response should match json:
    """
    {
    "total": 1,
    "data": [
        {
            "ID": 1,
            "CreatedAt": "2020-05-03T13:52:34.030197Z",
            "UpdatedAt": "2020-05-03T13:52:34.030197Z",
            "DeletedAt": null,
            "title": "title",
            "post": "post",
            "author": "author"
        }
    ]
  }
    """