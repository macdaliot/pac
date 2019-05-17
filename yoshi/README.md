## yoshi

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
* [http://integration.yoshi.pac.pyramidchallenges.com.s3-website.us-east-2.amazonaws.com](http://integration.yoshi.pac.pyramidchallenges.com.s3-website.us-east-2.amazonaws.com)
* [https://api.yoshi.pac.pyramidchallenges.com/api/<service-name>](https://api.yoshi.pac.pyramidchallenges.com/api/<service-name>)
* [https://jenkins.yoshi.pac.pyramidchallenges.com](https://jenkins.yoshi.pac.pyramidchallenges.com)
* [https://sonarqube.yoshi.pac.pyramidchallenges.com](https://sonarqube.yoshi.pac.pyramidchallenges.com)
* [https://selenium.yoshi.pac.pyramidchallenges.com](https://selenium.yoshi.pac.pyramidchallenges.com)

### Description
Project created by PAC
