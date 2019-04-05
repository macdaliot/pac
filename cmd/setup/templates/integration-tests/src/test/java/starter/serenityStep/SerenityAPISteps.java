package starter.serenityStep;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.net.HttpURLConnection;
import java.net.URL;
import java.util.List;
import java.util.Properties;

import org.apache.oltu.oauth2.client.OAuthClient;
import org.apache.oltu.oauth2.client.URLConnectionClient;
import org.apache.oltu.oauth2.client.request.OAuthClientRequest;
import org.apache.oltu.oauth2.client.response.OAuthJSONAccessTokenResponse;
import org.apache.oltu.oauth2.common.message.types.GrantType;
import org.junit.Assert;
import org.openqa.selenium.WebDriver;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.ObjectMapper;

import io.restassured.response.Response;
import io.restassured.specification.RequestSpecification;
import model.Employee;
import net.serenitybdd.core.Serenity;
import net.serenitybdd.rest.SerenityRest;
import net.thucydides.core.annotations.Managed;
import net.thucydides.core.annotations.Step;
import net.thucydides.core.util.EnvironmentVariables;
import net.thucydides.core.util.SystemEnvironmentVariables;

public class SerenityAPISteps {

	private ObjectMapper objectMapper = new ObjectMapper();
	private EnvironmentVariables variables = SystemEnvironmentVariables.createEnvironmentVariables();
	private String baseUri = variables.getProperty("baseUri");
	private String baseUrl = variables.getProperty("baseUrl");
	private String token = variables.getProperty("token");
	private Properties serenityProperties;
	private static RequestSpecification requestSpecification;
    String environmentToken = System.getenv("PEACHFIVE_ACCESSTOKEN");

	@Managed
	WebDriver driver;

	@Step
	public void get200Status(String endpoint) throws Throwable {

		/*
		 * let cookies = browser.getCookie(); cookies.forEach(function(cookie){ //fetch
		 * your properties console.log(cookie.domain); });
		 */
		//http://api.peachfive.pac.pyramidchallenges.com/api/expensereports
		Response response = SerenityRest.given().header("Authorization", getAuthToken()).baseUri("http://api.peachfive.pac.pyramidchallenges.com/").get("api/expensereports");
		Assert.assertTrue(response.getStatusCode() == 200);
	}

	@Step
	public void post200Status(String endpoint) throws Throwable {
		// Create an object for serialization
		Employee employee = Employee.API_TEST_USER;
		// Serialize the object
		String json = objectMapper.writeValueAsString(employee);
		// POST
		Response response = SerenityRest.given().header("Authorization", getAuthToken()).baseUri(baseUri).body(json)
				.post("/api/expensereports");

		System.out.println("the environment variable is: " + environmentToken);
		System.out.println("the response is: " + response.body());
		Assert.assertTrue(response.getStatusCode() == 200);
	}

	@Step
	public void deserializeGETObject(String endpoint) throws Throwable {
		// GET response
		Response response = SerenityRest.given().header("Authorization", getAuthToken()).baseUri(baseUri).get(endpoint);

		// Deserialize the response body
		// If single object use...
		// EmployeeModel employee =
		// objectMapper.readValue(response.asString(),EmployeeModel.class);
		// If list of objects use...
		List<Employee> employees = objectMapper.readValue(response.asString(), new TypeReference<List<Employee>>() {
		});

		// Store the deserialized object(s) as a Serenity session variable.
		Serenity.setSessionVariable("employees").to(employees);

		// This can be called elsewhere in the scenario with
		// List<EmployeeModel> employees = Serenity.sessionVariableCalled("employees");
	}

	@Step
	public void navigateToInvalidPage() throws JsonProcessingException {
		// System.out.println(RestAssured.get("http://integration.peachfive.pac.pyramidchallenges.com/invalid").body().toString());
		// Create an object for serialization
		String endpoint = "null";
		Employee employee = Employee.API_TEST_USER;
		// Serialize the object
		String json = objectMapper.writeValueAsString(employee);
		// POST
		// GET response

		Response response = SerenityRest.given().header("Authorization", getAuthToken()).baseUri(baseUri + "/null")
				.get("/api/expensereports");
		System.out.println("the response is: " + response);

		/*
		 * SerenityRest. given(). auth(). preemptive(). basic("test@psi-it.com","test").
		 * when().
		 * get("http://integration.peachfive.pac.pyramidchallenges.com/invalid").
		 * then(). assertThat(). body("authentication_result",equalTo("Success"));
		 */
	}

	@Step
	public void invalidPageError() {
		SerenityRest.when().get("http://integration.peachfive.pac.pyramidchallenges.com/invalid").then().assertThat()
				.statusCode(200);

	}

	/*
	 * public String getAuthToken() { loadSerenityProperties();
	 * 
	 * String clientId = "test@psi-it.com"; String clientSecret = "test"; String
	 * oauthTokenUri =
	 * "http://integration.peachfive.pac.pyramidchallenges.com/invalid";
	 * //getAPIUrl(); // RestAssured.baseURI = baseURI2; //requestSpecification =
	 * restWithCert(serenityProperties) // .param("username",
	 * internalUsername).param("password", internalPassword) // .param("grant_type",
	 * "password").auth().preemptive().basic(myUscisClientId, myUscisClientSecret);
	 * To address HTTPS SSL certification; works in Dev and Test
	 * //RequestSpecification requestSpecification = new RequestSpecification();
	 * 
	 * restWithCert(serenityProperties)
	 * .given().config(RestAssured.config().sslConfig(new
	 * SSLConfig().relaxedHTTPSValidation("TLS"))) .param("grant_type",
	 * "client_credentials").auth().preemptive().basic(clientId, clientSecret);
	 * 
	 * requestSpecification.baseUri(oauthTokenUri); Response response =
	 * requestSpecification.post(); JsonPath jsonPath = new
	 * JsonPath(response.asString()); String authToken =
	 * jsonPath.getString("token_type"); authToken = authToken+" " +
	 * jsonPath.getString("access_token"); return authToken; }
	 */

	private String getAuthToken() {
 
		String token = "";
		 String TOKEN_REQUEST_URL =
		 "http://integration.peachfive.pac.pyramidchallenges.com/api/expensereports"; 
		 String CLIENT_ID ="test@psi-it.com"; 
		 String CLIENT_SECRET = "test"; 
		 try {
	            OAuthClient client = new OAuthClient(new URLConnectionClient());

	            OAuthClientRequest request =
	                    OAuthClientRequest.tokenLocation(TOKEN_REQUEST_URL)
	                    .setGrantType(GrantType.CLIENT_CREDENTIALS)
	                    .setClientId(CLIENT_ID)
	                    .setClientSecret(CLIENT_SECRET)
	                    // .setScope() here if you want to set the token scope
	                    .buildQueryMessage();

	            token =
	                    client.accessToken(request, OAuthJSONAccessTokenResponse.class)
	                    .getAccessToken();

	            HttpURLConnection resource_cxn =
	                    (HttpURLConnection)(new URL(TOKEN_REQUEST_URL).openConnection());
	            resource_cxn.addRequestProperty("Authorization", "Bearer " + token);

	            InputStream resource = resource_cxn.getInputStream();

	            // Do whatever you want to do with the contents of resource at this point.

	            BufferedReader r = new BufferedReader(new InputStreamReader(resource, "UTF-8"));
	            String line = null;
	            while ((line = r.readLine()) != null) {
	                System.out.println(line);
	            }
	        } catch (Exception exn) {
	            exn.printStackTrace();
	        }
		 return token;
	    }
	

	@Step 
	public void initialGET(String endpoint) throws IOException {
		  SerenityRest.
          when().get(baseUri + endpoint).
          then().assertThat().statusCode(200);
  }
	


	
}
/*
 * public void loadSerenityProperties() { serenityProperties = new Properties();
 * FileInputStream in; try { in = new FileInputStream("serenity.properties");
 * serenityProperties.load(in); in.close(); } catch (IOException e) {
 * e.printStackTrace();
 * System.out.println("Error reading serenity.properties file"); } }
 */
