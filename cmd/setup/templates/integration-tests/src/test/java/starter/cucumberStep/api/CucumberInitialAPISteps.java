package starter.cucumberStep.api;

import cucumber.api.PendingException;
import cucumber.api.java.en.When;
import cucumber.api.java.en.Then;
import cucumber.api.junit.Cucumber;
import net.thucydides.core.annotations.Steps;
import org.junit.runner.RunWith;
import starter.serenityStep.api.SerenityInitialAPISteps;

    public class CucumberInitialAPISteps {

        @Steps
        SerenityInitialAPISteps steps;

        @Then("^the information is sent over and the (.+) will have a status of 200 after request$")
        public void userInfoSent(String endpoint) throws Exception{
            steps.userInfoSent(endpoint);
        }
    }
