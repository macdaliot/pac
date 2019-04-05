Feature: ACME API Features

  @InitialAPI
  Scenario Outline: Initial GET Request to endpoint
    Then the <endpoint> will have a status of 200

    Examples: 
      | endpoint            |
      | /api/expensereports |

  @API
  Scenario: Invalid Page
    Given I am on an invalid page
    Then the status is 200 and there is a page error message

  @API
  Scenario Outline: GET
    Given I get a 200 status from GET <endpoint>

    Examples: 
      | endpoint |
      | null     |

  @API
  Scenario Outline: POST
    Given I get a 200 status from POST <endpoint>

    Examples: 
      | endpoint |
      | null     |
