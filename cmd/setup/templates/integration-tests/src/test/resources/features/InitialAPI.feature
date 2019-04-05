Feature: Initial Acme API Features

  @InitialAPIStephen
  Scenario Outline: Users endpoint GET request

    Then the information is sent over and the <endpoint> will have a status of 200 after request

    Examples:
      | endpoint |
      | users   |