package starter.serenityStep;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.net.HttpURLConnection;
import java.net.URL;
import java.util.List;
import java.util.Properties;

import com.jayway.restassured.http.ContentType;
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
import net.serenitybdd.core.Serenity;
import net.serenitybdd.rest.SerenityRest;
import net.thucydides.core.annotations.Managed;
import net.thucydides.core.annotations.Step;
import net.thucydides.core.util.EnvironmentVariables;
import net.thucydides.core.util.SystemEnvironmentVariables;

import static org.hamcrest.Matchers.equalTo;

public class SerenityFrontEndSteps {

    private ObjectMapper objectMapper = new ObjectMapper();
    private EnvironmentVariables variables = SystemEnvironmentVariables.createEnvironmentVariables();
    private String baseFrontEndUri = variables.getProperty("baseFrontEndUri");
    private Properties serenityProperties;
    private static RequestSpecification requestSpecification;

    @Managed
    WebDriver driver;

    @Step
    public void checkFrontEndStatus() throws IOException {

            SerenityRest
                    .given()
                    .baseUri(baseFrontEndUri)
                    .when()
                    .get()
                    .then()
                    .assertThat().statusCode(200);

    }
}