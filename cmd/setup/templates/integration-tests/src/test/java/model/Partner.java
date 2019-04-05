package model;

public enum Partner {

	PartnerTest1("4A103DDB-807D-EC33-3BA5-8EBB965F3970", "Reed", "Maynard", "Mara",
			"Phasellus.fermentum@Suspendissenonleo.net", "(024) 1303 9282");

	private String id;
	private String first_name;
	private String last_name;
	private String middle_name;
	private String email;
	private String ph_number;

	private Partner(String id, String first_name, String last_name, String middle_name, String email,
			String ph_number) {
		this.id = id;
		this.first_name = first_name;
		this.last_name = last_name;
		this.middle_name = middle_name;
		this.email = email;
		this.ph_number = ph_number;
	}

}
