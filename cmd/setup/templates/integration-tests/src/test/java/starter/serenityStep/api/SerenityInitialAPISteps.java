package starter.serenityStep.api;

import com.fasterxml.jackson.databind.ObjectMapper;
import io.restassured.RestAssured;
import model.User;
import model.User2;
import net.thucydides.core.annotations.Step;
import io.restassured.response.Response;
import net.serenitybdd.rest.SerenityRest;
import org.junit.Before;
import org.junit.BeforeClass;
import org.junit.Test;

import static org.assertj.core.api.Assertions.assertThat;


public class SerenityInitialAPISteps extends SerenityBaseAPISteps {

    @Step
    public void userInfoSent(String endpoint) throws Exception {
        // Create an object for serialization;

        baseApiUri = variables.getProperty("baseApiUri");

        User user = null;
            /*switch(userInput)
            {
                case "TPOOLE":
                    user = User.TPOOLE;
                    break;

                    default:
                        user = User.TPOOLE;
            }*/

        User2 user2 = new User2();
        user2.setEmail("test@test.com");
        user2.setFirst_name("ftest");
        user2.setId("26");
        user2.setLast_name("ltest");
        user2.setMiddle_name("mtest");
        user2.setPh_number("444-555-6666");

            // Serialize the object
            objectMapper = new ObjectMapper();
            String json = objectMapper.writeValueAsString(user2);
            json = json.replaceAll("”", "\"");
            // POST
            Response response = SerenityRest
                    .given()
                    .baseUri(baseApiUri)
                    .body(json)
                    .get("users");
            assertThat(response.getStatusCode()).isEqualTo(200);

    }
}