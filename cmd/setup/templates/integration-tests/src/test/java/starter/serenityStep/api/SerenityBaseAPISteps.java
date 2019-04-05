package starter.serenityStep.api;

import java.util.Properties;

import org.junit.BeforeClass;
import org.openqa.selenium.WebDriver;

import com.fasterxml.jackson.databind.ObjectMapper;

import io.restassured.specification.RequestSpecification;
import net.thucydides.core.annotations.Managed;
import net.thucydides.core.util.EnvironmentVariables;
import net.thucydides.core.util.SystemEnvironmentVariables;

public class SerenityBaseAPISteps {

    @Managed
    WebDriver driver;

    protected static ObjectMapper objectMapper;
    protected static EnvironmentVariables variables;
    protected static String baseUri;
    protected static String baseUrl;
    protected static String baseApiUri;
    protected static String token;
    protected static Properties serenityProperties;
    protected static RequestSpecification requestSpecification;
    protected static String environmentToken;

    @BeforeClass
    public static void init() {
        objectMapper = new ObjectMapper();
        variables = SystemEnvironmentVariables.createEnvironmentVariables();
        baseUri = variables.getProperty("baseUri");
        baseUrl = variables.getProperty("baseUrl");
        token = variables.getProperty("token");
        environmentToken = System.getenv("PEACHFIVE_ACCESSTOKEN");
        baseApiUri = variables.getProperty("baseApiUri");
    }








}