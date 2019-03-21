package starter.serenityStep;

import java.util.List;

import org.junit.Assert;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.ObjectMapper;

import io.restassured.response.Response;
import model.Employee;
import net.serenitybdd.core.Serenity;
import net.serenitybdd.rest.SerenityRest;
import net.thucydides.core.annotations.Step;
import net.thucydides.core.util.EnvironmentVariables;
import net.thucydides.core.util.SystemEnvironmentVariables;

public class SerenityAPISteps {

	private ObjectMapper objectMapper = new ObjectMapper();
	private EnvironmentVariables variables = SystemEnvironmentVariables.createEnvironmentVariables();
	private String baseUri = variables.getProperty("baseUri");
	private String baseUrl = variables.getProperty("baseUrl");
	private String token = variables.getProperty("token");

  @Step
	public void get200Status(String endpoint) throws Throwable {
		Response response = SerenityRest
				.given()
				.header("Authorization",token)
				.baseUri(baseUri)
				.get(endpoint);
		Assert.assertTrue(response.getStatusCode() == 200);
	}

  @Step
	public void post200Status(String endpoint) throws Throwable {
		// Create an object for serialization
		Employee employee = Employee.API_TEST_USER;
		// Serialize the object
		String json = objectMapper.writeValueAsString(employee);
		// POST
		Response response = SerenityRest
				.given()
				.header("Authorization",token)
				.baseUri(baseUri)
				.body(json)
				.post(endpoint);
		

		Assert.assertTrue(response.getStatusCode() == 200);
	}

	
  @Step
	public void deserializeGETObject(String endpoint) throws Throwable {
		// GET response
		Response response = SerenityRest
				.given()
				.header("Authorization",token)
				.baseUri(baseUri)
				.get(endpoint);

		// Deserialize the response body
		// If single object use...
		// EmployeeModel employee = objectMapper.readValue(response.asString(),EmployeeModel.class);
		// If list of objects use...
		List<Employee> employees = objectMapper.readValue(response.asString(), new TypeReference<List<Employee>>(){});

		// Store the deserialized object(s) as a Serenity session variable.
		Serenity.setSessionVariable("employees").to(employees);

		//This can be called elsewhere in the scenario with
		// List<EmployeeModel> employees = Serenity.sessionVariableCalled("employees");
	}
	
  @Step
  public void navigateToInvalidPage() throws JsonProcessingException {
		// Create an object for serialization
    	String endpoint = "null";
		Employee employee = Employee.API_TEST_USER;
		// Serialize the object
		String json = objectMapper.writeValueAsString(employee);
		// POST
		// GET response
		
		 Response response = SerenityRest .given() .header("Authorization",token)
		 .baseUri(baseUri + "/null") .get(endpoint);
		
		/*
		 * SerenityRest. given(). auth(). preemptive(). basic("test@psi-it.com","test").
		 * when().
		 * get("http://integration.peachfive.pac.pyramidchallenges.com/invalid").
		 * then(). assertThat(). body("authentication_result",equalTo("Success"));
		 */
  }
    
  @Step
  public void invalidPageError() {
    SerenityRest.
      when().get("http://integration.peachfive.pac.pyramidchallenges.com/invalid").
      then().assertThat().statusCode(200);
  }
}
