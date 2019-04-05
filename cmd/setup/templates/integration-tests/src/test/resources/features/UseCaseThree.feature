Feature: Use Case Three Front End Tests

  @UseCaseThree
  Scenario Outline: User is able to create a partner organization
    Given the user is on the application
    And the user <username> logs in with the password test
    Then the user <password> will be logged in
    When the attributes are entered for the partner <partner> and manager <manager>
    Then the partner organization will be created

    Examples: 
      | username      | password | partner      | manager      |
      | test@acme.com | test     | Test1        | Test1        |
      | test@acme.com | test     | InvalidTest1 | InvalidTest1 |
