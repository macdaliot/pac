node("scikit-sagemaker") {
    
    stage ("SCM Checkout") {
        try{
            git branch: 'data-science-pipeline', credentialsId: 'jenkins-git', url: 'git@github.com:PyramidSystemsInc/bdso-fork.git'
            
        } catch (Exception ex) {
            print("Failed to checkout {{ .projectName }}")
            print(ex.toString())
        }
    }
    RESOLVED_SOURCE_VERSION=sh(returnStdout: true, script: "git rev-parse HEAD")

    stage ("Test $ENVIRONMENT") {
        sh """
            aws s3 cp s3://\$PIPELINE_DATA_BUCKET/dev/configuration_dev.json ./configuration_dev.json;
            aws sagemaker list-endpoints > endpoints.json
            python3 data-science/nlp/test.py configuration_dev.json endpoints.json $ENVIRONMENT
        """
    }

}