package starter.pageObject;

import net.serenitybdd.core.annotations.findby.FindBy;
import net.serenitybdd.core.pages.WebElementFacade;
import org.openqa.selenium.NoSuchElementException;

public class HomePage {

	@FindBy(id = "searchBar") private WebElementFacade searchBar;
	@FindBy(id = "searchButton") private WebElementFacade searchButton;

	//shared details
	@FindBy(id = "sub") private WebElementFacade logId;
	@FindBy(id = "type") private WebElementFacade transType;
	@FindBy(id = "time") private WebElementFacade time;
	@FindBy(id = "detail") private WebElementFacade transTypeDet;
	private final String[] transDetTypes = {"Car", "Mortgage", "Misc", "Transfer",
			"Withdrawal", "Deposit", "Address Edit", "SSN Edit", "Phone Edit", "Name Edit"};
	@FindBy(id = "transId") private WebElementFacade transId;
	@FindBy(id = "acctId") private WebElementFacade acctId;
	@FindBy(id = "ssn") private WebElementFacade ssn;
	@FindBy(id = "transTS") private WebElementFacade transTS;
	@FindBy(id = "transStat") private WebElementFacade transStat;
	@FindBy(id = "acctMangement") private WebElementFacade acctMangement;

	//fields shown after account id is entered
	@FindBy(id = "name") private WebElementFacade name;
	@FindBy(id = "phone") private WebElementFacade phone;
	@FindBy(id = "address") private WebElementFacade address;
	@FindBy(id = "creditScore") private WebElementFacade creditScore;

	//Checking
	@FindBy(id = "checkingId") private WebElementFacade checkingId;
	@FindBy(id = "checkingType") private WebElementFacade checkingType;
	//Checking types
	private final String withdrawal_from = "withdrawal_from";
	private final String deposit_to = "deposit_to";
	private final String transfer_both = "transfer_both";
	//
	@FindBy(id = "amount") private WebElementFacade amount;

	private void typeInput(WebElementFacade element, String input) {
		element.sendKeys(input);
	}

	private void clickElement(WebElementFacade element) {
		element.click();
	}

	public void search(String input) {
		typeInput(searchBar, input);
		clickElement(searchButton);
	}

	public String getElementText(String element) {
		String elementText = "";
		try {
			switch (element) {
				case "loan":
					elementText = logId.getText();
					break;
				case "checking":
					elementText = transType.getText();
					break;
				case "detail":
					elementText = transTypeDet.getText();
					break;
				case "time":
					elementText = time.getText();
					break;
				case "account mgmt":
					elementText = acctMangement.getText();
					break;
				case "address":
					elementText = address.getText();
					break;
				case "creditScore":
					elementText = creditScore.getText();
					break;
			}

		} catch(NoSuchElementException e) {
			elementText = null;
		}
		return elementText;
	}

	public boolean elementIsPresent(String element) {
		try {
			switch (element) {
				case "loan":
					logId.getText();
					break;
				case "checking":
					transType.getText();
					break;
				case "account mgmt":
					acctMangement.getText();
					break;
				case "address":
					address.getText();
					break;
				case "creditScore":
					creditScore.getText();
					break;
			}

		} catch(NoSuchElementException e) {
			return false;
		}
		return true;
	}

}
