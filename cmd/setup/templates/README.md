## {{.projectName}}



### Setting up the Environment Locally

1. Ensure you have one or more services and one authentication service
    * If you do not have services, just run steps #3 and #9 to simply run the front end

2. Run `npm i` at the root of the project to install JS/TS module dependencies for all services

3. Run `npm i` in the `app` directory
    * `npm i` at the root does not install dependencies in this directory

4. Copy and paste the `.env` files, obtained through a secure means, into each service's directory (i.e. `services/auth` or `services/gizmos`)

5. Run `npm run build` individually in the service directories you wish to run locally (i.e. `services/auth`, `services/posts`, `services/threads`)

6. Ensure nothing is running on ports 3000, 8001, or 8080
    * Example command to check if a port is used: `netstat -a | grep <port-number>` (no output means the port is not being used)

7. Use the dev-tool provided [here](https://github.com/PyramidSystemsInc/tech-challenge-tools) to deploy locally 
    * **If using Windows, run the tool using Windows Command Prompt**
    * In case you already have the tool installed, the command is:
        * `pac-devtool localDeploy` run from the root of the project
    * If you are looking to install the tool, click the link above to see installation instructions

8. Run `docker ps -a` to verify you have the appropriate Docker containers running
    * There should be one container for the database (DynamoDB)
    * There should be one container for the reverse proxy (HaProxy)
    * There should be one container for **each** service

9. Run `npm run dev-server` in the `app` directory to start the front-end on `http://localhost:8080`
    * This command eats up a terminal/terminal tab

10. Verify all is working:
    1. Verify the front end is up by loading `http://localhost:8080`
    2. Verify authentication is up by clicking on the "Login" button on the front end' home page and follow the login prompts
    3. Verify all other services are up by:
        * Either attempting a GET request at `http://localhost:3000/api/<service-name>` -OR-
        * Checking out a service's documentation at `http://localhost:3000/api/<service-name>/docs`

###### Top Troubleshooting Tip

**If step #7 above shows containers that have been created, but are not running **and** you are running Windows**

1. Open the `Docker for Windows` settings through little whale icon in the taskbar's system tray

2. OPTIONAL: Restart Docker

3. Click on the "Shared Drives" tab

4. Click "reset credentials" at the bottom

5. Click the checkmark to denote you wish to share `C:/`

6. Click "Apply" and insert your credentials

7. Run `docker run --rm -v c:/Users:/data -p 8081:8081 alpine ls /data`. A successful run (where the command outputs the user accounts on your Windows computer) will verify:
    * The Docker daemon is running
    * Docker is able to mount volumes
    * Docker can map ports

* Once running, recompiling a service (`npm run build`) will update the service without needing to re-run the dev tool command



### Setting up Authentication in the Cloud

1. Ensure you have created an authentication service

2. Deploy the app by running it through the Jenkins pipeline

3. Login to [Auth0's management console](https://manage.auth0.com/dashboard/us/pyramidsystems/)

4. Click "Applications" on the left

5. Click "ACME Login"

6. Scroll down to "Allowed Callback URLs"

7. Paste `https://<environment-name>.{{.projectName}}.pac.pyramidchallenges.com/api/auth/callback` for each environment as a new entry (entries are separated by commas) at the bottom

###### Top Troubleshooting Tip

**If the callback URL keeps spinning and the browser tab reads, "Working..."**

1. Login to the AWS console

2. Find the Lambda authentication function for your project

3. Check if the environment variables listed match the environment variables at `./services/auth/lambdaEnvironment.txt`



### Post-Setup Authentication Steps

* All services do not require authentication to begin with. In order to require a valid JWT token to use the service, uncomment the `@Security` annotation in the `controller.ts` file in the service's directory



### Helpful Links

* Front-End - `https://<environment>.{{.projectName}}.pac.pyramidchallenges.com`, where `<environment>` is one of the following:
    * integration
    * staging
    * production
* Example service endpoint - `https://<environment>.{{.projectName}}.pac.pyramidchallenges.com/api/<service-name>`, where `<environment>` is one of the following:
    * integration
    * staging
    * production
* [https://jenkins.{{.projectName}}.pac.pyramidchallenges.com](https://jenkins.{{.projectName}}.pac.pyramidchallenges.com)
* [https://sonarqube.{{.projectName}}.pac.pyramidchallenges.com](https://sonarqube.{{.projectName}}.pac.pyramidchallenges.com)
* [https://selenium.{{.projectName}}.pac.pyramidchallenges.com](https://selenium.{{.projectName}}.pac.pyramidchallenges.com)



### Description

{{.description}}
