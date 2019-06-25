package starter.serenityStep;

import static org.junit.Assert.assertTrue;

import net.thucydides.core.annotations.Step;
import net.thucydides.core.util.EnvironmentVariables;
import net.thucydides.core.util.SystemEnvironmentVariables;
import starter.pageObject.LoginPage;

public class SerenityLoginSteps {

	LoginPage page;
	EnvironmentVariables variables = SystemEnvironmentVariables.createEnvironmentVariables();
	String baseFrontEndUri = variables.getProperty("baseFrontEndUri");

	/**
	 * User opens the application
	 * 
	 * @throws Throwable
	 */
	@Step
	public void openApplication() {
		page.openAt(baseFrontEndUri);
	}

	/**
	 * User logs in with password and username
	 * 
	 * @param username
	 * @param password
	 * @throws Throwable
	 */
	@Step
	public void userLogin(String username, String password) {
		page.clickInitialLoginButton();
		page.typeUserNameInput(username);
		page.typePasswordInput(password);
		page.clickLogin();
	}

	/**
	 * User is logged in with name showing in top right corner
	 * 
	 * @param username
	 */
	@Step
	public void loggedIn(String username, String displayName) {
		assertTrue(page.getUserNameHeader().equals(displayName));
	}

	@Step
	public void notLoggedIn() {
		assertTrue(page.getAuth0HeaderText());
	}
	
	@Step 
	public void loginErrorMessage(String error) {
		if (error.equals("wrongEmailPassword")) {
			System.out.println("the error is: " + page.getErrorMessage());
			assertTrue(page.getErrorMessage().equalsIgnoreCase("WRONG EMAIL OR PASSWORD."));
		}
	}

}
