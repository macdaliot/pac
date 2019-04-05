package starter;

import org.junit.runner.RunWith;

import cucumber.api.CucumberOptions;
import net.serenitybdd.cucumber.CucumberWithSerenity;

@RunWith(CucumberWithSerenity.class)
@CucumberOptions(
        plugin = {"pretty"},
        features = "src/test/resources/features",
        tags = {"~@API", "~@UseCaseThreeAPITODO", "~@UseCaseThree"}
)
public class CucumberTestRunner {

}
