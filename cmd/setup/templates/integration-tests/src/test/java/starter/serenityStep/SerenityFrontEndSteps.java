package starter.serenityStep;

import java.io.IOException;
import java.util.Properties;
import org.openqa.selenium.WebDriver;
import com.fasterxml.jackson.databind.ObjectMapper;
import io.restassured.specification.RequestSpecification;
import net.serenitybdd.rest.SerenityRest;
import net.thucydides.core.annotations.Managed;
import net.thucydides.core.annotations.Step;
import net.thucydides.core.util.EnvironmentVariables;
import net.thucydides.core.util.SystemEnvironmentVariables;
import starter.pageObject.HomePage;
import static org.junit.Assert.assertTrue;

public class SerenityFrontEndSteps {
    private ObjectMapper objectMapper = new ObjectMapper();
    private EnvironmentVariables variables = SystemEnvironmentVariables.createEnvironmentVariables();
    private String baseFrontEndUri = variables.getProperty("baseFrontEndUri");
    private Properties serenityProperties;
    private static RequestSpecification requestSpecification;
    private HomePage page;

    @Managed
    WebDriver driver;

    @Step
    public void checkFrontEndStatus() throws IOException {
        
        SerenityRest
            .given()
            .baseUri(baseFrontEndUri)
            .when()
            .get()
            .then()
            .assertThat().statusCode(200);

    }

    @Step
    public void searchSetup() throws IOException {
        driver.get(baseFrontEndUri);
    }

    @Step
    public void customerSearch(String id) throws IOException {
        page.search(id);
    }

    @Step
    public void checkDashboardElements() {
        assertTrue(page.elementIsPresent("loan")
                && page.elementIsPresent("checking")
                && page.elementIsPresent("detail")
                && page.elementIsPresent("time"));
    }

    @Step
    public void appropriateDashboardElements() {
        String loan = page.getElementText("loan");
        String checking = page.getElementText("checking");
        String accountmgmt = page.getElementText("time");
        String detail = page.getElementText("detail");
        //hit endpoint
        //get rest response (list)
        try {
            assertTrue(checkCustomerCheckingInfo(checking) &&
                    checkCustomerLoanInfo(loan) && checkCustomerAccountManagementInfo(accountmgmt));
        } catch (Exception e) {
            assertTrue(false);
        }
    }

    @Step
    public void checkSSNInfo() throws IOException {
        String ssn = page.getElementText("ssn");
        assertTrue(ssn.length() == 4);
    }

    @Step
    public boolean checkCustomerLoanInfo(String loan) {
        SerenityRest
        .then();
        return false;
    }

    @Step
    public boolean checkCustomerCheckingInfo(String checking) {
        SerenityRest
        .given()
        .baseUri(baseFrontEndUri);
        return false;
    }

    @Step
    public boolean checkCustomerAccountManagementInfo(String activity) throws IOException {
        SerenityRest
        .given()
        .baseUri(baseFrontEndUri);
        return false;
    }

}