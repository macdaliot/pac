package starter.pageObject;

import net.serenitybdd.core.annotations.findby.FindBy;
import net.serenitybdd.core.pages.PageObject;
import net.serenitybdd.core.pages.WebElementFacade;

import java.util.HashMap;
import java.util.Map;

public class HomePage extends PageObject implements PageObjectInterface {

	//Example Elements
	@FindBy(xpath = "//*[@id=\"container\"]/div/div")
	private WebElementFacade govWebsiteBanner;

	@FindBy(xpath = "//*[@id=\"container\"]/div/header/div[2]/a/div/button")
	private WebElementFacade loginButton;

	private Map<String, WebElementFacade> map;

	public WebElementFacade getElementWebFacade(String element) {
		switch (element) {
			case "govWebsiteBanner":
				return govWebsiteBanner;
			case "loginButton":
				return loginButton;
			default: return null;
		}
	}


	public void clickElement(WebElementFacade element) {
		element.click();
	}

	public boolean elementIsPresent(WebElementFacade element) {
		return element.isCurrentlyVisible();
	}

	public String getElementText(WebElementFacade element) {
		return element.getText();
	}
}
