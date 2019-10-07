

node ("scikit-sagemaker") {
    try {

        stage ("Data Pipeline") {
            build job: "data-pipeline", parameters: []    
        }

        stage ("SCM Checkout") {
            try{
                git branch: 'data-science-pipeline', credentialsId: 'jenkins-git', url: 'git@github.com:PyramidSystemsInc/bdso-fork.git'
                
            } catch (Exception ex) {
                print("Failed to checkout {{ .projectName }}")
                print(ex.toString())
            }
        }
        RESOLVED_SOURCE_VERSION=sh(returnStdout: true, script: "git rev-parse HEAD")

        stage ("Build Models") {
            //export necessary here since these are resolved at runtime and may change throughout build process
            sh """
                export RESOLVED_COMMIT_ID="$RESOLVED_SOURCE_VERSION"
                export ENVIRONMENT="dev"
                python3 -m venv build; 	
                bash ./build/bin/activate;
                cd data-science; 
                pip3 install -r  build-requirements.txt --user; 
                pip3 install aws sagemaker --user;
                python3 nlp/build.py
                aws s3 cp ./configuration_qa.json s3://\$PIPELINE_DATA_BUCKET/dev/configuration_dev.json;                
            """
        }

        

        stage ("Deploy Dev Endpoint") {
           build job: "deploy-endpoint", parameters: [
                [$class: 'StringParameterValue', name: 'ENVIRONMENT', value: "dev"]
            ]
        }

        stage ("Test Dev Endpoint") {
            build job: "test-endpoint", parameters: [
                [$class: 'StringParameterValue', name: 'ENVIRONMENT', value: "dev"]
            ]
        }

        stage ("Promote Artifacts") {
            build job: "promote-artifacts", parameters: [
                [$class: 'StringParameterValue', name: 'ENVIRONMENT', value: "dev"]
            ]
        }

        stage ("Deploy E2E Endpoint") {
           build job: "deploy-endpoint", parameters: [
                [$class: 'StringParameterValue', name: 'ENVIRONMENT', value: "e2e"]
            ]
        }

        stage ("Test E2E Endpoint") {
            build job: "test-endpoint", parameters: [
                [$class: 'StringParameterValue', name: 'ENVIRONMENT', value: "e2e"]
            ]
        }

        stage ("Promote Artifacts") {
            build job: "promote-artifacts", parameters: [
                [$class: 'StringParameterValue', name: 'ENVIRONMENT', value: "e2e"]
            ]
        }

        stage ("Deploy Pro Endpoint") {
           build job: "deploy-endpoint", parameters: [
                [$class: 'StringParameterValue', name: 'ENVIRONMENT', value: "pro"]
            ]
        }

        stage ("Test Pro Endpoint") {
            build job: "test-endpoint", parameters: [
                [$class: 'StringParameterValue', name: 'ENVIRONMENT', value: "pro"]
            ]
        }

    } finally {
        sh "rm -rf *"
    }

}
