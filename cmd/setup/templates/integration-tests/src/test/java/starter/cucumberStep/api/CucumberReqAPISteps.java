package starter.cucumberStep.api;

import cucumber.api.PendingException;
import cucumber.api.java.en.Then;
import cucumber.api.junit.Cucumber;
import net.thucydides.core.annotations.Steps;
import org.junit.runner.RunWith;
import starter.serenityStep.api.SerenityReqAPISteps;

@RunWith(Cucumber.class)
public class CucumberReqAPISteps {

    @Steps
    SerenityReqAPISteps steps;

    @Then("^the information is sent and the (.+) will have a status of 200 after request$")
    public void checkStatus(String endpoint) throws Exception {
        steps.checkStatus(endpoint);
    }

}