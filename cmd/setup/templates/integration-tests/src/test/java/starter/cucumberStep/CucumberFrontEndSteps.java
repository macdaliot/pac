package starter.cucumberStep;

import cucumber.api.java.en.And;
import cucumber.api.java.en.Given;
import cucumber.api.java.en.Then;
import cucumber.api.java.en.When;
import net.thucydides.core.annotations.Steps;
import starter.serenityStep.SerenityFrontEndSteps;
import starter.serenityStep.SerenityLoginSteps;
import com.fasterxml.jackson.core.JsonProcessingException;

public class CucumberFrontEndSteps {
  @Steps
  SerenityFrontEndSteps steps;

  @Steps
  SerenityLoginSteps loginSteps;

  @Given("^the user is a customer service representative$")
  public void customerServiceRepLogin() {
    String custServRepUsername = "test@psi-it.com";
    String custServPassword = "test";
    loginSteps.openApplication();
    loginSteps.userLogin(custServRepUsername, custServPassword);
  }

  @Given("^the user is a customer$")
  public void customerLogin() {
    String custUsername = "test@psi-it.com";
    String custPassword = "test";
    loginSteps.openApplication();
    loginSteps.userLogin(custUsername, custPassword);
  }

  @Then("^the page will return a status of 200$")
  public void checkFrontEndStatus() throws Throwable {
    steps.checkFrontEndStatus();
  }

  @Given("^the user is a customer service representative $")
  public void searchSetup() throws Throwable {
    steps.searchSetup();
  }

  @When("^the user searches for a customer with account ID: (.+)$")
  public void customerSearch(String id) throws Throwable {
    steps.customerSearch(id);
  }

  @Then("^the page should display customer (.+) for account ID: (.+)$")
  public void checkCustomerLoanInfo(String loan) throws Throwable {
    steps.checkCustomerLoanInfo(loan);
  }

  @And("^the page should display customer (.+)$")
  public void checkCustomerCheckingInfo(String checking) throws Throwable {
    steps.checkCustomerCheckingInfo(checking);
  }

  @And("^the page should display account management (.+)$")
  public void checkCustomerAccountManagementInfo(String activity) throws Throwable {
    steps.checkCustomerAccountManagementInfo(activity);
  }

  @Then("^the loan, checking, and account mgmt activity will show$")
  public void heckDashboardElements(){
    steps.checkDashboardElements();
  }

  @Then("^the user will only be able to see the appropriate dashboard activity$")
  public void appropriateUserInfo(){
    steps.checkDashboardElements();
  }

  @And("^the user will only be able to see the last 4 digits of customer's SSN$")
  public void checkSSNInfo() throws Throwable {
    steps.checkSSNInfo();
  }

  @Given("^I am on an invalid page$")
  public void navigateToInvalidPage() throws JsonProcessingException {
    //steps.navigateToInvalidPage();
  }

  @Then("^the status is 200 and there is a page error message$")
  public void invalidPageError() {
    //steps.invalidPageError();
  }
}
