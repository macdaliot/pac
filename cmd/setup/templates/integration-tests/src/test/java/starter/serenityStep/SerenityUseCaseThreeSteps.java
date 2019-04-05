package starter.serenityStep;

import model.Manager;
import model.Partner;

public class SerenityUseCaseThreeSteps {
	
	private Manager manager1;
	private Partner partner1;

	public void partnerManagerAttributesEntered(String partner, String manager) {

		if (manager.equals("Test1")) {
			manager1 = Manager.ManagerTest1;
		} else if (manager.equals("InvalidTest1")) {
			manager1 = Manager.ManagerInvalidTest1;
		}
		if (partner.equals("Test1")) {
			partner1 = Partner.PartnerTest1;
		} else if (partner.equals("InvalidTest1")) {
			//partner1 = Partner.PartnerInvalidTest1;
		}

	}
	
	//check the attributes
	public void partnerIsCreated() {
		
	}
}
