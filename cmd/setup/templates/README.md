## {{.projectName}}

### Getting Started
1. Navigate to your Jenkins' IP address. Wait until Jenkins has a "Credentials" link on the left sidebar. When the link appears, run from your project's root directory:

* pac automate jenkins

2. Create one or more microservices with:

* cd services
* pac add service --name <service-name>

3. Push to GitHub:

* git add --all
* git commit -m "Initial commit"
* git push origin master

The above workflow will set up your project and deploys it to AWS

### Helpful Links
* [http://integration.{{.projectName}}.pac.pyramidchallenges.com.s3-website.us-east-2.amazonaws.com](http://integration.{{.projectName}}.pac.pyramidchallenges.com.s3-website.us-east-2.amazonaws.com)
* [http://api.{{.projectName}}.pac.pyramidchallenges.com/api/<service-name>](http://api.{{.projectName}}.pac.pyramidchallenges.com/api/<service-name>)
* [http://jenkins.{{.projectName}}.pac.pyramidchallenges.com](http://jenkins.{{.projectName}}.pac.pyramidchallenges.com)
* [http://sonarqube.{{.projectName}}.pac.pyramidchallenges.com](http://sonarqube.{{.projectName}}.pac.pyramidchallenges.com)
* [http://selenium.{{.projectName}}.pac.pyramidchallenges.com](http://selenium.{{.projectName}}.pac.pyramidchallenges.com)

### Description
{{.description}}
