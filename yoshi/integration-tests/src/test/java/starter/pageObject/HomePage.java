package starter.pageObject;

import net.serenitybdd.core.annotations.findby.FindBy;
import net.serenitybdd.core.pages.WebElementFacade;

public class HomePage {

	/*
	 * @FindBy(id = "partnerNum") private WebElementFacade partnerNum;
	 * 
	 * @FindBy(id = "partnerName") private WebElementFacade partnerName;
	 * 
	 * @FindBy(id = "lName") private WebElementFacade lName;
	 * 
	 * @FindBy(id = "fName") private WebElementFacade fName;
	 * 
	 * @FindBy(id = "mName") private WebElementFacade mName;
	 * 
	 * @FindBy(id = "phoneNum") private WebElementFacade phoneNum;
	 * 
	 * @FindBy(id = "email") private WebElementFacade email;
	 * 
	 * @FindBy(id = "cancel") private WebElementFacade cancel;
	 * 
	 * @FindBy(id = "next") private WebElementFacade next;
	 */
	
	@FindBy(xpath = "//*[@id=\"container\"]/div/main/div/div/div[1]/div/div[1]/div/input") private WebElementFacade partnerNum;
	@FindBy(xpath = "//*[@id=\"container\"]/div/main/div/div/div[1]/div/div[2]/div/input") private WebElementFacade partnerName;
	@FindBy(id = "//*[@id=\"container\"]/div/main/div/div/div[2]/div/div[1]/div/input") private WebElementFacade lName;
	@FindBy(id = "//*[@id=\"container\"]/div/main/div/div/div[2]/div/div[2]/div/input") private WebElementFacade fName;
	@FindBy(id = "//*[@id=\"container\"]/div/main/div/div/div[2]/div/div[3]/div/input") private WebElementFacade mName;
	@FindBy(id = "//*[@id=\"container\"]/div/main/div/div/div[3]/div/div[1]/div/input") private WebElementFacade phoneNum;
	@FindBy(id = "//*[@id=\"container\"]/div/main/div/div/div[3]/div/div[2]/div/input") private WebElementFacade email;
	@FindBy(id = "//*[@id=\"Cancel\"]/button") private WebElementFacade cancel;
	@FindBy(xpath = "//*[@id=\"Next\"]/button") private WebElementFacade next;
	
	public void typeParterNumInput(String text) {
		partnerNum.type(text);
	}

	
	public void clickNext() {
		next.click();
	}
	
	public void clickCancel() {
		cancel.click();
	}


}
