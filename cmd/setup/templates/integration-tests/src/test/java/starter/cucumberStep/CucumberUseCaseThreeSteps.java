package starter.cucumberStep;

import cucumber.api.java.en.Then;
import cucumber.api.java.en.When;
import net.thucydides.core.annotations.Steps;
import starter.serenityStep.SerenityUseCaseThreeSteps;

public class CucumberUseCaseThreeSteps {
	
	@Steps
	SerenityUseCaseThreeSteps steps;

	@When("^the attributes are entered for the partner (.+) and manager (.+)$") 
	public void partnerManagerAttributesEntered(String partner, String manager) {
		steps.partnerManagerAttributesEntered(partner, manager);
	}
	
	@Then("^the partner organization will be created$")
	public void partnerIsCreated() {
		steps.partnerIsCreated();
	}
}
