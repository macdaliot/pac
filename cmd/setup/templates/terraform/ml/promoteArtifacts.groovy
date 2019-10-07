node {

    stage ("SCM Checkout") {
        try{
            git branch: 'data-science-pipeline', credentialsId: 'jenkins-git', url: 'git@github.com:PyramidSystemsInc/bdso-fork.git'
            
        } catch (Exception ex) {
            print("Failed to checkout {{ .projectName }}")
            print(ex.toString())
        }
    }

    stage ("Promote $ENVIRONMENT Artifacts") {
        nextEnv = getNextEnv();
        sh """ 
            cd terraform/ml;
            aws s3 cp s3://\$PIPELINE_DATA_BUCKET/$ENVIRONMENT/configuration_${ENVIRONMENT}.json ./configuration_${ENVIRONMENT}.json;
            aws s3 cp s3://\$PIPELINE_DATA_BUCKET/$ENVIRONMENT/configuration_${ENVIRONMENT}.json s3://\$PIPELINE_DATA_BUCKET/$nextEnv/configuration_${nextEnv}.json;
            """
            
        def model_data = sh returnStdout: true, script: "cd terraform/ml; python3 python/parse-dev-output.py configuration_${ENVIRONMENT}.json"
        def new_model = model_data.replaceAll(ENVIRONMENT, nextEnv);
        print model_data
        print new_model
        sh """
            OLD_DATA=$model_data
            MODEL_DATA=$new_model
            aws s3 cp \$OLD_DATA \$MODEL_DATA
        """
    }
    
}

def getNextEnv() {
    if("dev".equals(ENVIRONMENT)) {
        return "e2e";
    } else if("e2e".equals(ENVIRONMENT)) {
        return "prod"
    }
    throw new IllegalArgumentException("Uknown Environment Passed $ENVIRONMENT expecting 'dev' or 'e2e'")
}