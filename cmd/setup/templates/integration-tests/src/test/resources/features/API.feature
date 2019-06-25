Feature: API Features

  @TEST
  Scenario Outline: HTTP GET Request to Endpoint
    Then the <endpoint> will have a status of <status>

    Examples:
      | endpoint      | status        |


  @NOTREADY
  Scenario Outline: HTTP POST Request to Endpoint
    Then the <endpoint> will have a status of <status>

    Examples:
      | endpoint      | status        |
