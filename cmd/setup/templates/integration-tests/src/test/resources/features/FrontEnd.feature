Feature: Front End Features

  @FrontEnd
  Scenario: Checking the Front End Status
    Then page will have a status of 200


  @FrontEnd
  Scenario: Check Government Banner render
    Given the user is on the home page
    Then the government banner should render

  @FrontEnd
  Scenario: Check Login Button render
    Given the user is on the home page
    Then the login button should render
    And the user should be able to click on the button




