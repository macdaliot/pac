package starter.Steps;

import Utils.DataGenerator;
import cucumber.api.java.en.And;
import cucumber.api.java.en.Given;
import cucumber.api.java.en.Then;
import cucumber.api.java.en.When;
import net.serenitybdd.core.Serenity;
import net.serenitybdd.core.pages.WebElementFacade;
import net.thucydides.core.util.EnvironmentVariables;
import net.thucydides.core.util.SystemEnvironmentVariables;
import org.apache.commons.lang3.RandomStringUtils;
import starter.Pages.BasePageObject;

import java.util.List;

import static org.assertj.core.api.Assertions.assertThat;

public class AcmeSteps {

    BasePageObject page;
    DataGenerator dataGenerator = new DataGenerator();
    EnvironmentVariables variables = SystemEnvironmentVariables.createEnvironmentVariables();
    String baseUri = variables.getProperty("baseUri");
    String baseUrl = variables.getProperty("baseUrl");
    String token = variables.getProperty("token");

    @Given("^i open the home page$")
    public void iOpenTheHomePage() throws Throwable {
        page.openAt(baseUrl);
    }

    @When("^i login with username (.+) and password (.+)$")
    public void iLoginWithUsernameAndPassword(String username, String password) throws Throwable {
        System.out.println();
    }
}

