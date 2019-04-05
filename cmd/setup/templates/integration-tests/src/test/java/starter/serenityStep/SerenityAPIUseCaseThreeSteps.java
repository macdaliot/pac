package starter.serenityStep;

import static io.restassured.RestAssured.given;

import java.io.IOException;
import java.io.UnsupportedEncodingException;

import org.apache.http.HttpResponse;
import org.apache.http.client.ClientProtocolException;
import org.apache.http.client.HttpClient;
import org.apache.http.client.methods.HttpPost;
import org.apache.http.entity.StringEntity;
import org.apache.http.impl.client.HttpClientBuilder;
import java.io.BufferedInputStream;
import java.io.InputStream;
import java.io.OutputStream;
import java.net.HttpURLConnection;
import java.net.URL;
import org.apache.commons.io.IOUtils;
import org.json.JSONObject;
import com.google.gson.Gson;

import io.restassured.RestAssured;
import io.restassured.http.ContentType;
import io.restassured.parsing.Parser;
import model.Manager;
import net.thucydides.core.util.EnvironmentVariables;
import net.thucydides.core.util.SystemEnvironmentVariables;

public class SerenityAPIUseCaseThreeSteps {

	// TODO: get token
	private String token = "";
	private String enterPartnerInfoEndpoint = "";
	private String updatePartnerInfoEndpoint = "";
	private String getPartnerInfoEndpoint = "";

	private EnvironmentVariables variables = SystemEnvironmentVariables.createEnvironmentVariables();
	private String baseApiUri= variables.getProperty("baseUri");
	private String baseUrl = variables.getProperty("baseUrl");
	// private String token = variables.getProperty("token");

	// POST: enter acme manager
	public void enterAcmeManager() throws ClientProtocolException, IOException {
		String userEndpoint = baseApiUri + "users";
		String json = " {\r\n" + 
				"   \"id\": \"4FE2E67E-306B-883F-3135-B9FCEEC43EE1\",\r\n" + 
				"   \"first_name\": \"Kelsey\",\r\n" + 
				"   \"last_name\": \"Reese\",\r\n" + 
				"   \"middle_name\": \"Yvonne\",\r\n" + 
				"   \"email\": \"vestibulum.neque@sempercursusInteger.com\",\r\n" + 
				"   \"ph_number\": \"(024) 8164 8105\"\r\n" + 
				" }";
		try {
			URL url = new URL(userEndpoint);
			HttpURLConnection conn = (HttpURLConnection) url.openConnection();
			conn.setConnectTimeout(5000);
			conn.setRequestProperty("Content-Type", "application/json; charset=UTF-8");
			conn.setDoOutput(true);
			conn.setDoInput(true);
			conn.setRequestMethod("POST");
			OutputStream os = conn.getOutputStream();
			os.write(json.getBytes("UTF-8"));
			os.close();
			// read the response
			InputStream in = new BufferedInputStream(conn.getInputStream());
			String result = IOUtils.toString(in, "UTF-8");
			System.out.println(result);
			System.out.println("result after Reading JSON Response");
			JSONObject myResponse = new JSONObject(result);
			System.out.println("jsonrpc- " + myResponse.getString("jsonrpc"));
			System.out.println("id- " + myResponse.getInt("id"));
			System.out.println("result- " + myResponse.getString("result"));
			in.close();
			conn.disconnect();
		} catch (Exception e) {
			System.out.println(e);
		}
	}

	// POST: enter partner organization information
	public void enterPartnerInfo() {
		String userEndpoint = baseApiUri + "partners";
		String json = " {\r\n" + 
				"   \"id\": \"4FE2E67E-306B-883F-3135-B9FCEEC43EE1\",\r\n" + 
				"   \"first_name\": \"Kelsey\",\r\n" + 
				"   \"last_name\": \"Reese\",\r\n" + 
				"   \"middle_name\": \"Yvonne\",\r\n" + 
				"   \"email\": \"vestibulum.neque@sempercursusInteger.com\",\r\n" + 
				"   \"ph_number\": \"(024) 8164 8105\"\r\n" + 
				" }";
		try {
			URL url = new URL(userEndpoint);
			HttpURLConnection conn = (HttpURLConnection) url.openConnection();
			conn.setConnectTimeout(5000);
			conn.setRequestProperty("Content-Type", "application/json; charset=UTF-8");
			conn.setDoOutput(true);
			conn.setDoInput(true);
			conn.setRequestMethod("POST");
			OutputStream os = conn.getOutputStream();
			os.write(json.getBytes("UTF-8"));
			os.close();
			// read the response
			InputStream in = new BufferedInputStream(conn.getInputStream());
			String result = IOUtils.toString(in, "UTF-8");
			System.out.println(result);
			System.out.println("result after Reading JSON Response");
			JSONObject myResponse = new JSONObject(result);
			System.out.println("jsonrpc- " + myResponse.getString("jsonrpc"));
			System.out.println("id- " + myResponse.getInt("id"));
			System.out.println("result- " + myResponse.getString("result"));
			in.close();
			conn.disconnect();
		} catch (Exception e) {
			System.out.println(e);
		}
	}

	// PUT: update partner organization information
	public void updatePartnerInfo() {
		RestAssured.defaultParser = Parser.JSON;
		given().headers("Content-Type", ContentType.JSON, "Accept", ContentType.JSON).when()
				.put(updatePartnerInfoEndpoint).then().contentType(ContentType.JSON).extract().response();
	}

	// GET: get partner information
	public void getPartnerInfo() {
		RestAssured.defaultParser = Parser.JSON;
		given().headers("Content-Type", ContentType.JSON, "Accept", ContentType.JSON).when().get(getPartnerInfoEndpoint)
				.then().contentType(ContentType.JSON).extract().response();
	}
}
