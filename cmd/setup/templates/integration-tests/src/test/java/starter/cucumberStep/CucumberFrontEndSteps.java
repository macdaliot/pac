package starter.cucumberStep;

import cucumber.api.PendingException;
import cucumber.api.java.en.Then;
import cucumber.api.junit.Cucumber;
import net.thucydides.core.annotations.Steps;
import starter.serenityStep.SerenityFrontEndSteps;


public class CucumberFrontEndSteps {

    @Steps
    SerenityFrontEndSteps steps;

    @Then("^the page will have a status of 200$")
    public void checkFrontEndStatus() throws Throwable {
        steps.checkFrontEndStatus();
    }

}