Feature: ACME API Features

  @API
  Scenario: Invalid Page
    Given I am on an invalid page
    Then the status is 200 and there is a page error message

#  @API
#  Scenario Outline: GET
#    Given I get a 200 status from GET <endpoint>
#
#    Examples: 
#      | endpoint |
#      | null     |
#
#  @API
#  Scenario Outline: POST
#    Given I get a 200 status from POST <endpoint>
#
#    Examples: 
#      | endpoint |
#      | null     |
#
