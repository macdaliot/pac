## PAC: Pyramid Application Constructor



The Pyramid App Constructor (PAC) is a toolkit to help jumpstart the application development process, specifically designed for compressed time-duration events like hackathons and tech challenges. PAC will generate scaffolding composed of reusable components, templates, and pipelines to help accelerate development velocity, while ensuring security and quality discipline, to achieve acceptable software hygiene. PAC is an evolving toolkit, and currently supports the MERN stack (MongoDB, Express, React, Node). It leverages Jenkins for pipelines, Auth0 for authentication, AWS as the cloud platform, and is supported by relevant open source libraries

***



## Running the PAC CLI 

PAC is a CLI tool that builds a local JS application development environment on a developer's workstation,  creates a github webhook and a dedicated Jenkins server with a master pipline provisioner, then constructs and deploys the various service components of the overall application to AWS -while maitaining the git webhook within the local Jenkins (and it's dynamically generated pipelines). These ordered parts of the application build and deploy workflow are executed as the following PAC commands. 

***
#### **pac** `setup`
***
This command is **only to be executed once** per application project in PAC. 
Run this command and it will perform the following actions:

- Creates the JS directory skeleton (for the ReactJS FE content)
- Performs the Jenkins image deployment
- Creates the local DynamoDB image in Docker
- Creates the Application Load Balancer and HTTP listener
- Creates the GitHub Repository for your project
- Constructs a webhook in Jenkins that is pointed to your GitHub project

***

#### **pac** `automate jenkins`
***
This command is **only to be executed once** per application project in PAC. 

`cd ${project directory}`

Run this command within your project root directory and it will perform the following actions:

- Creates the master pipeline provisioner in Jenkins (shell exec script)
- Stages the template layout for pipeline scaffolding 

***

#### **pac** `add service --name ${service_name}`
***
This command is to be **executed for every service deployment** of your application project in PAC. 

`cd ./svc/`

1] Run this command within your svc directory and it will perform the following actions:

- Creates  ${service_name} microservice files
- Constructs local DynamoDB table
- Deploys ${service_name} in a local docker container
- Stages the pipeline provisioner to create a ${service_name} pipeline:
 -- `	git push ..` will trigger the provisioner to create new pipeline
 
 2] Manually Run (in Jenkins) the new pipeline and it will perform the following actions:
 
 - Creates Application Load Balancer ${service_name}  Rules 
 - Creates ${service_name} Target Group  
- Deploys/Updates  ${service_name} code as a Lambda function
- Stages the pipeline provisioner to create a ${service_name} pipeline:
 - Creates DynamoDB table
 

 ***

#### **pac** `create pipeline --name ${project_name}`
***
DEPRECATED

***

#### **pac** `add stage  --name ${stage_name}`
***

DEPRECATED


