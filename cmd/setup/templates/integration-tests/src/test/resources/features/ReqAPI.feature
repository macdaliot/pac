Feature: Requirements ACME API Features

  @ReqAPIStephen
  Scenario Outline: Requirements returns a status of 200

    Then the information is sent and the <endpoint> will have a status of 200 after request
    Examples:
      | endpoint     |
      | requirements |
