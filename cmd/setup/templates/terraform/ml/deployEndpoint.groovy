node ("scikit-sagemaker") {
    try {
        stage ("SCM Checkout") {
            try{
                git branch: 'data-science-pipeline', credentialsId: 'jenkins-git', url: 'git@github.com:PyramidSystemsInc/bdso-fork.git'
                
            } catch (Exception ex) {
                print("Failed to checkout {{ .projectName }}")
                print(ex.toString())
            }
        }
        
        stage ("Deploy $env.ENVIRONMENT Endpoint") {
            sh """
                    cd terraform/ml; 
                    aws s3 cp s3://\$PIPELINE_DATA_BUCKET/dev/configuration_dev.json ./configuration_dev.json;
                    echo "environment_abbr = $ENVIRONMENT"
                    python3 python/parse-dev-output.py configuration_dev.json;
                    cd main/sagemaker/endpoint/$ENVIRONMENT;
                    terraform init;
                    terraform apply --var-file ../../../terraform.tfvars -var 'environment_abbr=$ENVIRONMENT' -var sagemaker_role=\$SAGEMAKER_ROLE --auto-approve
                """
        } 
    } finally {
        sh "rm -rf *"
    }  
}