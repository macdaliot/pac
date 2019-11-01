package starter.serenityStep;

import com.fasterxml.jackson.databind.ObjectMapper;
import io.restassured.specification.RequestSpecification;
import net.serenitybdd.rest.SerenityRest;
import net.thucydides.core.annotations.Managed;
import net.thucydides.core.annotations.Step;
import net.thucydides.core.util.EnvironmentVariables;
import net.thucydides.core.util.SystemEnvironmentVariables;
import org.openqa.selenium.WebDriver;
import org.openqa.selenium.support.ui.ExpectedConditions;
import starter.pageObject.HomePage;

import java.io.IOException;
import java.util.Properties;

import static org.junit.Assert.*;

public class SerenityFrontEndSteps extends SerenitySteps {

    private ObjectMapper objectMapper = new ObjectMapper();
    private EnvironmentVariables variables = SystemEnvironmentVariables.createEnvironmentVariables();
    private String baseFrontEndUri = variables.getProperty("baseFrontEndUri");
    private Properties serenityProperties;
    private static RequestSpecification requestSpecification;
    private HomePage homePage;

    @Step
    public void checkFrontEndStatus() throws IOException {

        SerenityRest
                .given()
                .baseUri(baseFrontEndUri).relaxedHTTPSValidation()
                .when()
                .get()
                .then()
                .assertThat().statusCode(200);

    }

    @Step
    public void NavigateToHomePage(WebDriver driver) {
        driver.get(baseFrontEndUri);
    }

    @Step
    public void checkGovWebsiteBannerRender() {
        assertNotNull((homePage.getElementWebFacade("govWebsiteBanner")));

    }

    @Step
    public void checkLoginButtonRender() throws IOException {
        assertNotNull((homePage.getElementWebFacade("loginButton")));
    }

    @Step
    public void checkLoginButtonClick() throws IOException {
        ExpectedConditions.elementToBeClickable(homePage.getElementWebFacade("loginButton"));
    }

}
