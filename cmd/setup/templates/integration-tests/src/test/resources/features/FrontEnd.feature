Feature: Front End Features

  @UIREADY
  Scenario: Page Exists
    Then the page will return a status of 200


  @NOTREADY
  Scenario: Page Does Not Exist
    Given the address bar has a route not recognized by the app
    Then the status is 200 and there is an error message
