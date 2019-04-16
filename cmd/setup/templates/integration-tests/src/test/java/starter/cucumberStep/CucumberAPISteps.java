package starter.cucumberStep;

import com.fasterxml.jackson.core.JsonProcessingException;

import cucumber.api.java.en.Given;
import cucumber.api.java.en.Then;
import net.thucydides.core.annotations.Steps;
import starter.serenityStep.SerenityAPISteps;


public class CucumberAPISteps {
	
	@Steps
	SerenityAPISteps steps;
	
	@Then("^the (.+) will have a status of (.+) with token: (.+)$")
	public void GETApiTest(String endpoint, int expectedStatusCode, boolean withToken) throws Throwable {
		steps.GETApiTest(endpoint, expectedStatusCode, withToken);
	}
	
	@Given("^I get a 200 status from POST (.+)$")
	public void post200Status(String endpoint) throws Throwable {
		//steps.post200Status(endpoint);
	}
	
	@Given("^I deserialize and store object from GET (.+)$") 
	public void deserializeGETObject(String endpoint) throws Throwable {
		System.out.println("The endpoint is: " + endpoint);
		//steps.deserializeGETObject(endpoint);
	}
	
    @Given("^I am on an invalid page$")
    public void navigateToInvalidPage() throws JsonProcessingException {
    	//steps.navigateToInvalidPage();
    }

    @Then("^the status is 200 and there is a page error message$")
    public void invalidPageError() {
    	//steps.invalidPageError();
    }

}
