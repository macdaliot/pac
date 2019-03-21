package starter.cucumberStep;

import org.openqa.selenium.WebDriver;
import org.openqa.selenium.chrome.ChromeDriver;

import cucumber.api.java.After;
import cucumber.api.java.en.And;
import cucumber.api.java.en.Given;
import cucumber.api.java.en.Then;
import net.thucydides.core.annotations.Managed;
import net.thucydides.core.annotations.Steps;
import starter.serenityStep.SerenityLoginSteps;

public class CucumberLoginSteps {

	@Steps
	SerenityLoginSteps steps;
	@Managed
	WebDriver driver;

	/**
	 * User opens the application
	 * 
	 * @throws Throwable
	 */
	@Given("^the user is on the application$")
	public void openApplication() {
		steps.openApplication();
	}

	/**
	 * User logs in with password and username
	 * 
	 * @param username
	 * @param password
	 * @throws Throwable
	 */
	@And("^the user (.+) logs in with the password (.+)$")
	public void userLogin(String username, String password) {
		steps.userLogin(username, password);
	}

	/**
	 * User is logged in with name showing in top right corner
	 * 
	 * @param username
	 */
	@Then("^the user (.+) will be logged in$")
	public void loggedIn(String username) {
		steps.loggedIn(username);
	}

	@Then("^the user will not be logged in$")
	public void notLoggedIn() {
		steps.notLoggedIn();
	}

	@After
	public void closeBrowser() {
		driver.quit();
	}

}
