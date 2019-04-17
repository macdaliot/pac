#Feature: Login Test
#
#  @Login
#  Scenario Outline: Check if application login is working - positive
#    Given the user is on the application
#    And the user <username> logs in with the password <password>
#    Then the user <username> will be logged in as <displayName>
#
#    Examples:
#      | username        | password | displayName |
#      | test@psi-it.com | test     | test        |
#
#  @Login
#  Scenario Outline: Check if application login is working - negative
#    Given the user is on the application
#    And the user <username> logs in with the password <password>
#    Then the user will not be logged in
#    And there will be a <errorMessage> error message
#
#    Examples:
#      | username        | password   |   errorMessage     |
#      | test@psi-it.com | test2      | wrongEmailPassword |