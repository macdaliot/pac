## {{.projectName}}

### Getting Started
1. Navigate to your Jenkins' IP address. Wait until Jenkins has a "Credentials" link on the left sidebar. When the link appears, run from your project's root directory:

* pac automate jenkins

2. Create one or more microservices with:

* cd svc
* pac add service --name <service-name>

3. Push to GitHub:

* git add --all
* git commit -m "Initial commit"
* git push origin master

The above workflow will set up your project and deploys it to AWS

### Helpful Links
* [https://integration.{{.projectName}}.pac.pyramidchallenges.com](https://integration.{{.projectName}}.pac.pyramidchallenges.com) (takes ~15 minutes to provision)
* [https://api.{{.projectName}}.pac.pyramidchallenges.com/api/<service-name>](https://api.{{.projectName}}.pac.pyramidchallenges.com/api/<service-name>)
* [https://jenkins.{{.projectName}}.pac.pyramidchallenges.com](https://jenkins.{{.projectName}}.pac.pyramidchallenges.com)
* [https://sonarqube.{{.projectName}}.pac.pyramidchallenges.com](https://sonarqube.{{.projectName}}.pac.pyramidchallenges.com)
* [https://selenium.{{.projectName}}.pac.pyramidchallenges.com](https://selenium.{{.projectName}}.pac.pyramidchallenges.com)

### Description
{{.description}}
