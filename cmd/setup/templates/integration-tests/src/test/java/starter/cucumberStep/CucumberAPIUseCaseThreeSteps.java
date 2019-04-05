package starter.cucumberStep;

import java.io.IOException;

import org.apache.http.client.ClientProtocolException;

import cucumber.api.java.en.Then;
import net.thucydides.core.annotations.Steps;
import starter.serenityStep.SerenityAPIUseCaseThreeSteps;

public class CucumberAPIUseCaseThreeSteps {

	@Steps
	SerenityAPIUseCaseThreeSteps steps;
	
	@Then("^the (.+) information will be returned$")
	public void partnerOrganizationInfo(String param) throws ClientProtocolException, IOException {
		if (param.contains("partner")) {
			if (param.contains("Test1")) {
				steps.enterPartnerInfo();
			}
		}
		else if (param.contains("manager")) {
			if (param.contains("Test1")) {
				steps.enterAcmeManager();
			}
		}
	}
	
	@Then("^the (.+) will be updated to (.+)$") 
	public void updateApprover(String approver1, String approver2) {
		steps.updatePartnerInfo();
	}
	
}
