package starter.serenityStep;

import com.fasterxml.jackson.databind.ObjectMapper;
import io.restassured.specification.RequestSpecification;
import java.io.IOException;
import java.util.Properties;
import net.serenitybdd.rest.SerenityRest;
import net.thucydides.core.annotations.Managed;
import net.thucydides.core.annotations.Step;
import net.thucydides.core.util.EnvironmentVariables;
import net.thucydides.core.util.SystemEnvironmentVariables;
import org.openqa.selenium.WebDriver;
import starter.pageObject.LoginPage;
import starter.util.LocalStorage;

public class SerenityAPISteps {
	@Managed
	WebDriver driver;

  LoginPage page;
  private ObjectMapper objectMapper = new ObjectMapper();
  private EnvironmentVariables variables = SystemEnvironmentVariables.createEnvironmentVariables();
  private String baseFrontEndUri = variables.getProperty("baseFrontEndUri");
  private String baseApiUri = variables.getProperty("baseApiUri");
  private String tokenName = variables.getProperty("jwtTokenKey");
  private Properties serenityProperties;
  private static RequestSpecification requestSpecification;
  private LocalStorage localStorage = new LocalStorage(driver);

  @Step
  public void httpGet(String endpoint, int expectedStatusCode) throws IOException {
    SerenityRest
      .given()
        .header("Accept", "application/json")
        .header("Content-Type", "application/json")
        .baseUri(baseApiUri)
      .when().get(endpoint)
      .then().assertThat().statusCode(expectedStatusCode);
/*
    boolean tokenExists = localStorage.isItemPresentInLocalStorage(tokenName);
    if (tokenExists) {
      String token = localStorage.getItemFromLocalStorage(tokenName);
      SerenityRest
        .given()
          .header("Accept", "application/json")
          .header("Authorization", "Bearer " + token)
          .header("Content-Type", "application/json")
          .baseUri(baseApiUri)
        .when().get(endpoint)
        .then().assertThat().statusCode(expectedStatusCode);
    } else {
      SerenityRest
        .given()
          .header("Accept", "application/json")
          .header("Content-Type", "application/json")
          .baseUri(baseApiUri)
        .when().get(endpoint)
        .then().assertThat().statusCode(expectedStatusCode);
    }
*/
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

  public void get200Status(String endpoint) {
    SerenityRest
      .given().baseUri(baseApiUri)
      .when().get(endpoint)
      .then().assertThat().statusCode(200);
  }

  @Step
  public void GETApiTokenTest(String endpoint, int expectedStatusCode, String username, String password, WebDriver driver) {
      page.openAt(baseFrontEndUri);
      page.clickInitialLoginButton();
      page.typeUserNameInput(username);
      page.typePasswordInput(password);
      page.clickLogin();
      LocalStorage ls = new LocalStorage(driver);
      String token = ls.getItemFromLocalStorage(tokenName);
      System.out.println("the token is " + token);
      SerenityRest
        .given()
          .header("Accept", "application/json")
          .header("Authorization", "Bearer " + token)
          .header("Content-Type", "application/json")
          .baseUri(baseApiUri)
        .when().get(endpoint)
        .then().assertThat().statusCode(expectedStatusCode);
  }
}

