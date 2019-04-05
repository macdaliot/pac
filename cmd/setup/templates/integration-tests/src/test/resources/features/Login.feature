Feature: Login Test

  @Login
  Scenario Outline: Check if application login is working - positive
    Given the user is on the application
    And the user <username> logs in with the password test
    Then the user <password> will be logged in

    Examples: 
      | username        | password |
      | test@acme.com | test     |

  @Login
  Scenario Outline: Check if application login is working - negative
    Given the user is on the application
    And the user <username> logs in with the password <password>
    Then the user will not be logged in
    And there will be a <errorMessage> error message 

    Examples: 
      | username        | password | errorMessage       |
      | test@acme.com | test2    | wrongEmailPassword |