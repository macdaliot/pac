package starter.cucumberStep;

import cucumber.api.java.en.Given;
import cucumber.api.java.en.When;
import cucumber.api.java.en.Then;
import cucumber.api.java.en.And;
import net.thucydides.core.annotations.Managed;
import net.thucydides.core.annotations.Steps;
import org.openqa.selenium.WebDriver;
import starter.serenityStep.SerenityFrontEndSteps;
import starter.serenityStep.SerenityLoginSteps;


public class CucumberFrontEndSteps {

  @Steps
  SerenityFrontEndSteps steps;

  @Managed
  WebDriver driver;

  @Then("^the page will have a status of 200$")
  public void checkFrontEndStatus() throws Throwable {
    steps.checkFrontEndStatus();
  }

  @Given("^the user is on the home page$")
  public void NavigateToHomePage() throws Throwable {
    steps.NavigateToHomePage(driver);
  }

  @Then("the government banner should render$")
  public void checkGovWebsiteBannerRender() throws Throwable {
    steps.checkGovWebsiteBannerRender();
  }

  @Then("^the login button should render$")
  public void checkLoginButtonRender() throws Throwable {
    steps.checkLoginButtonRender();
  }

  @And("^the user should be able to click on the button$")
  public void checkLoginButtonClick() throws Throwable {
    steps.checkLoginButtonClick();
  }

}
