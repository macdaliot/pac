Feature: PAC API Features

  @API
  Scenario Outline: GET Request to endpoint
    Then the <endpoint> will have a status of <status> with token: <token>

    Examples:
      | endpoint  | status | token |

  @API
  Scenario: Invalid Page
    Given I am on an invalid page
    Then the status is 200 and there is a page error message

  @API
  Scenario Outline: POST
    Given I get a 200 status from POST <endpoint>

    Examples: 
      | endpoint  |
