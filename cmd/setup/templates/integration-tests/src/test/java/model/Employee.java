package model;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public enum Employee {

	API_TEST_USER("test", "test", "test@psi-it.com", "123-567-8901", "2677 Prosperity Ave", Arrays.asList("test1","test2","test3"),
			"test");

	private String firstName;
	private String lastName;
	private String email;
	private String phone;
	private String address;
	private List<String> skills;
	private String id;

	private Employee(String firstName, String lastName, String email, String phone, String address,
			List<String> skills, String id) {
		this.firstName = firstName;
		this.lastName = lastName;
		this.email = email;
		this.phone = phone;
		this.address = address;
		this.skills = skills;
		this.id = id;
	}

	public String getFirstName() {
		return firstName;
	}

	public void setFirstName(String firstName) {
		this.firstName = firstName;
	}

	public String getLastName() {
		return lastName;
	}

	public void setLastName(String lastName) {
		this.lastName = lastName;
	}

	public String getEmail() {
		return email;
	}

	public void setEmail(String email) {
		this.email = email;
	}

	public String getPhone() {
		return phone;
	}

	public void setPhone(String phone) {
		this.phone = phone;
	}

	public String getAddress() {
		return address;
	}

	public void setAddress(String address) {
		this.address = address;
	}

	public List<String> getSkills() {
		return skills;
	}

	public void setSkills(ArrayList<String> skills) {
		this.skills = skills;
	}

	public String getId() {
		return id;
	}

	public void setId(String id) {
		this.id = id;
	}

	@Override
	public String toString() {
		return "EmployeeModel{" + "firstName='" + firstName + '\'' + ", lastName='" + lastName + '\'' + ", email='"
				+ email + '\'' + ", phone='" + phone + '\'' + ", address='" + address + '\'' + ", skills=" + skills
				+ ", id='" + id + '\'' + '}';
	}
}