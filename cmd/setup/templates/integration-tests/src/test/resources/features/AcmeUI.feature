Feature: Template feature file for UI tests

  @UI @scenario1
  Scenario: An unbounded number of skills can be provided and stored with the profile
    Given i open the home page

  @UI @scenario2
  Scenario Outline: Scenario Outline Template
    Given i open the home page
    When i login with username <username> and password <password>

    Examples:
      | username      | password |
      | test@test.com | password |
