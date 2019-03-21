package starter.pageObject;

import java.util.List;

import org.openqa.selenium.By;

import net.serenitybdd.core.annotations.findby.FindBy;
import net.serenitybdd.core.pages.WebElementFacade;
import net.thucydides.core.pages.PageObject;

//fix this later:
//@DefaultUrl("http://api.peachfour.pac.pyramidchallenges.com/api/auth")
public class LoginPage extends PageObject {

	@FindBy(xpath = "//*[@id=\"container\"]/div/header/div[2]/a/div/button")
	private WebElementFacade initialLoginButton;
	@FindBy(css = "input[name='email']")
	private WebElementFacade userNameInput;
	@FindBy(css = "input[name='password']")
	private WebElementFacade passwordInput;
	@FindBy(css = "button[name='submit']")
	private WebElementFacade loginButton;
	@FindBy(css = "span[class='application-title'")
	private WebElementFacade applicationTitle;
	@FindBy(xpath = "//*[@id=\"auth0-lock-error-msg-email\"]/div")
	private WebElementFacade usernameErrorMessage;
	@FindBy(xpath = "//*[@id=\"auth0-lock-error-msg-password\"]/div")
	private WebElementFacade passwordErrorMessage;
	@FindBy(xpath = "//*[@id=\"container\"]/div/header/div[2]/span")
	private WebElementFacade applicationHeader;
	@FindBy(className = "username")
	private WebElementFacade usernameHeader;
	@FindBy(className = "auth0-lock-name")
	private WebElementFacade auth0Header;
	// success messages
	private static final String loginSuccessMessage = "success";
	private static final String loginPasswordFailMessage = "Can't be blank";
	private static final String loginUsernameFailMessage = "Can't be blank";
	private static final String loginButtonText = "Login";
	private static final String headerText = "peachfour";
	private static final String auth0headerText = "ACME Login";
	//username css
	private static final String userNameInputCSS = "input[name='email']";
	//application header 

	
	public String getUserNameInputCSS() {
		return userNameInputCSS;
	}
	
	public void clickInitialLoginButton() {
		initialLoginButton.click();
	}
	
	public boolean getAuth0HeaderText() {
		return auth0Header.getText().contentEquals(auth0headerText);
	}
	
	

	public String getloginButtonText() {
		return loginButtonText;
	}

	public String getheaderText() {
		return headerText;
	}
	
	public String getUserNameHeader() {
		return usernameHeader.getText();
	}

	public String getpasswordErrorMessage() {
		return passwordErrorMessage.getText();
	}

	public String getusernameErrorMessage() {
		return usernameErrorMessage.getText();
	}

	public String getStatusMessage(String status) {
		if (status.equals("fail")) {
			return loginButton.getText();
		} else if (status.equals("success")) {
			// should return the application header correctly
			return applicationHeader.getText();
		} else if (status.equals("failPassword")) {
			return passwordErrorMessage.getText();
		} else if (status.equals("failUsername")) {
			return usernameErrorMessage.getText();
		} else {
			return null;
		}
	}

	public static String getLoginsuccessmessage() {
		return loginSuccessMessage;
	}

	public static String getLoginpasswordfailmessage() {
		return loginPasswordFailMessage;
	}

	public static String getLoginusernamefailmessage() {
		return loginUsernameFailMessage;
	}

	public void typeUserNameInput(String text) {
		userNameInput.type(text);
	}

	public void typePasswordInput(String text) {
		passwordInput.type(text);
	}

	public String getApplicationTitle() {
		return applicationTitle.getText();
	}

	public void clickLogin() {
		loginButton.click();
	}

}
