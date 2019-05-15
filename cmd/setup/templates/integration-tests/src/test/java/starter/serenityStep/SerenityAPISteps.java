package starter.serenityStep;

import java.io.IOException;
import java.util.Properties;

import com.fasterxml.jackson.databind.ObjectMapper;

import org.json.JSONObject;
import org.junit.Assert;
import org.openqa.selenium.WebDriver;

import io.restassured.response.Response;
import io.restassured.specification.RequestSpecification;
import io.restassured.*;
import net.serenitybdd.rest.SerenityRest;
import net.thucydides.core.annotations.Managed;
import net.thucydides.core.annotations.Step;
import net.thucydides.core.util.EnvironmentVariables;
import net.thucydides.core.util.SystemEnvironmentVariables;

import static org.hamcrest.Matchers.equalTo;

public class SerenityAPISteps {

  private ObjectMapper objectMapper = new ObjectMapper();
  private EnvironmentVariables variables = SystemEnvironmentVariables.createEnvironmentVariables();
  private String baseFrontEndUri = variables.getProperty("baseFrontEndUri");
  private String baseApiUri = variables.getProperty("baseApiUri");
  private String token = variables.getProperty("token");
  private Properties serenityProperties;
  private static RequestSpecification requestSpecification;

  @Managed
  WebDriver driver;

  @Step
  public void GETApiTest(String endpoint, int expectedStatusCode, boolean withToken) throws IOException {

      if (withToken) {
        SerenityRest
            .given()
            .header("Accept", "application/json")
            .header("Authorization", "Bearer " + token)
            .header("Content-Type", "application/json")
            .baseUri(baseApiUri)
            .when()
            .get(endpoint)
            .then()
            .assertThat().statusCode(expectedStatusCode);
      } else {
        SerenityRest
            .given()
            .header("Accept", "application/json")
            .header("Content-Type", "application/json")
            .baseUri(baseApiUri)
            .when()
            .get(endpoint)
            .then()
            .assertThat().statusCode(expectedStatusCode);
      }
    }

  @Step
  public void post200Status(String endpoint) throws IOException {
    SerenityRest
    .given()
    .header("Accept", "application/json")
    .header("Content-Type", "application/json")
    .body("{\"id\":\"test\",\"buttons\":0,\"axles\":0}")
    .baseUri(baseApiUri)
    .when()
    .post(endpoint)
    .then()
    .assertThat().statusCode(200);
  }
}
