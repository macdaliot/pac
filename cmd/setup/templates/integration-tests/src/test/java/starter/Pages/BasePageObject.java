package starter.Pages;

import net.serenitybdd.core.annotations.findby.FindBy;
import net.serenitybdd.core.pages.WebElementFacade;
import net.thucydides.core.pages.PageObject;

import java.util.List;

public class BasePageObject extends PageObject{
	
	@FindBy(css="span[class='user-name']") WebElementFacade loginNameTab;
	@FindBy(css="button[title^='Detail']") List<WebElementFacade> editDetailButtons;



}
