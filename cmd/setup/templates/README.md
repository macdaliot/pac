## {{.projectName}}

### Setting up Authentication Locally

1. Ensure you have created one or more services and one authentication service

2. Navigate to `./services/<any-regular-service-name>` in any bash-like terminal

3. Change the `launch.sh` script to have executable permissions (`chmod 755 launch.sh`)

4. Ensure nothing is running on ports 3000 or 8001

5. Run `launch.sh`

6. If you did not create the project yourself, request the `.env` files that appear in each service folders from the person who created the project

###### Top Troubleshooting Tip

* If you are a Windows user and the `launch.sh` script failed:

1. Open up the script and replace all instances of `aws.cmd` to `aws`



### Setting up Authentication in the Cloud

1. Ensure you have created an authentication service

2. Deploy the app by running it through the Jenkins pipeline

3. Login to [Auth0's management console](https://manage.auth0.com/dashboard/us/pyramidsystems/)

4. Click "Applications" on the left

5. Click "ACME Login"

6. Scroll down to "Allowed Callback URLs"

7. Paste `https://api.{{.projectName}}.pac.pyramidchallenges.com/api/auth/callback` as a new entry (entries are separated by commas) at the bottom

###### Top Troubleshooting Tip

* If the callback URL keeps spinning and the browser tab reads, "Working...":

1. Login to the AWS console

2. Find the Lambda authentication function for your project

3. Check if the environment variables listed match the environment variables at `./services/auth/lambdaEnvironment.txt`



### Post-Setup Authentication Steps

* All services do not require authentication to begin with. In order to require a valid JWT token in order to use the service, uncomment the `@Security` annotation in the `controller.ts` file in the service's directory



### Helpful Links

* [http://integration.{{.projectName}}.pac.pyramidchallenges.com.s3-website.{{.region}}.amazonaws.com](http://integration.{{.projectName}}.pac.pyramidchallenges.com.s3-website.{{.region}}.amazonaws.com)
* [https://api.{{.projectName}}.pac.pyramidchallenges.com/api/<service-name>](https://api.{{.projectName}}.pac.pyramidchallenges.com/api/<service-name>)
* [https://jenkins.{{.projectName}}.pac.pyramidchallenges.com](https://jenkins.{{.projectName}}.pac.pyramidchallenges.com)
* [https://sonarqube.{{.projectName}}.pac.pyramidchallenges.com](https://sonarqube.{{.projectName}}.pac.pyramidchallenges.com)
* [https://selenium.{{.projectName}}.pac.pyramidchallenges.com](https://selenium.{{.projectName}}.pac.pyramidchallenges.com)



### Description

{{.description}}
