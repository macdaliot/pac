package starter.cucumberStep;

import cucumber.api.java.en.Given;
import cucumber.api.java.en.Then;
import net.thucydides.core.annotations.Managed;
import net.thucydides.core.annotations.Steps;
import net.thucydides.core.util.EnvironmentVariables;
import net.thucydides.core.util.SystemEnvironmentVariables;
import org.openqa.selenium.WebDriver;
import starter.pageObject.LoginPage;
import starter.serenityStep.SerenityAPISteps;

public class CucumberAPISteps {
  LoginPage page;
  EnvironmentVariables variables = SystemEnvironmentVariables.createEnvironmentVariables();
  String baseFrontEndUri = variables.getProperty("baseFrontEndUri");

  @Managed
  WebDriver driver;
  
  @Steps
  SerenityAPISteps steps;

  @Then("^the (.+) will have a status of (.+)$")
  public void httpGet(String endpoint, int expectedStatusCode) throws Throwable {
    steps.httpGet(endpoint, expectedStatusCode);
  }
  
  @Given("^I get a 200 status from POST (.+)$")
  public void post200Status(String endpoint) throws Throwable {
    steps.post200Status(endpoint);
  }

  @Given("^I get a 200 status from GET (.+)$")
  public void get200Status(String endpoint) throws Throwable {
    steps.get200Status(endpoint);
  }

  @Given("^I deserialize and store object from GET (.+)$") 
  public void deserializeGETObject(String endpoint) throws Throwable {
    System.out.println("The endpoint is: " + endpoint);
    //steps.deserializeGETObject(endpoint);
  }  
}
