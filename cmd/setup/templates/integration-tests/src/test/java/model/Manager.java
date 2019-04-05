package model;

public enum Manager {
	
	ManagerTest1("Test", "Test", "7031231234", "test@acme.com"),
	ManagerTest2("Test2", "Test2", "7031231234", "test2@acme.com"),
	ManagerInvalidTest1("InvalidTest", "InvalidTest", "number", "test@acme.com");
	
	private String fname;
	private String lname;
	private String number;
	private String email;

	private Manager(String fname, String lname, String number, String email) {
		this.fname = fname;
		this.lname = lname;
		this.number = number;
		this.email = email;
	}
}
