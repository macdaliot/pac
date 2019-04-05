package starter.serenityStep.api;

import com.fasterxml.jackson.databind.ObjectMapper;
import io.restassured.response.Response;
import model.Item;
import net.serenitybdd.rest.SerenityRest;
import net.thucydides.core.annotations.Step;

import static org.assertj.core.api.Assertions.assertThat;
import static starter.serenityStep.api.SerenityBaseAPISteps.baseApiUri;
import static starter.serenityStep.api.SerenityBaseAPISteps.objectMapper;

public class SerenityReqAPISteps extends SerenityBaseAPISteps {

    @Step
    public void checkStatus(String endpoint) throws Exception {

        Item item = new Item();
        item.setId("1");
        item.setPrice(21);
        item.setStatus('A');
        item.setTitle("Case of copy paper");


        baseApiUri = variables.getProperty("baseApiUri");
        // Serialize the object
        objectMapper = new ObjectMapper();
        String json = objectMapper.writeValueAsString(item);
        json = json.replaceAll("”", "\"");
        // POST
        Response response = SerenityRest
                .given()
                .baseUri(baseApiUri)
                .body(json)
                .get("requirements");
        assertThat(response.getStatusCode()).isEqualTo(200);
    }
}
