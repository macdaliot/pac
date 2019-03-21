Feature: Login Test

  Scenario: Check if application login is working
    Given the user is on the application
    And the user test@psi-it.com logs in with the password test
    Then the user test will be logged in

  Scenario: Check if application login is working
    Given the user is on the application
    And the user test@psi-it.com logs in with the password test2
    Then the user will not be logged in
