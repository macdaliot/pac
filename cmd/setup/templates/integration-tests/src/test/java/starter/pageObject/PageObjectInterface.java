package starter.pageObject;

import net.serenitybdd.core.pages.WebElementFacade;

public interface PageObjectInterface {

    WebElementFacade getElementWebFacade(String element);

    /**
     * Type Input for web element
     * @param element
     * @param input
     */
    static void typeInput(WebElementFacade element, String input) {
        element.sendKeys(input);
    }

    /**
     * Click element - web element
     * @param element
     */
    void clickElement(WebElementFacade element);

    String getElementText(WebElementFacade element);

    boolean elementIsPresent(WebElementFacade element);


}
