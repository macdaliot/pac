node {	
    try{
        def latestVersion = "";	
        def functions = [];	
        stage("Pull Artifacts") {	
            try {	
                // Clear the working directory to avoid using older build products previously unzipped here	
                sh 'rm -Rf *'	
                // Pull the release candidate zip file	
                sh 'aws s3 cp s3://builds.${{ .projectName }}_PROJECT_NAME.pac.pyramidchallenges.com/build-jenkins-release-through-e2e-testing-' + params.BUILD_NUMBER + '-rc.zip .'	
            } catch (e) {	
                sh 'echo "No release candidate exists for the build number specified"'	
                throw e	
            }	
        }	

        stage("Deploy ML Notebook") {	
            // Unzip the release candidate build products	
            sh 'unzip build-jenkins-release-through-e2e-testing-' + params.BUILD_NUMBER + '-rc.zip'	

            sh """	
                cd terraform/ml/main; 	
                terraform init; 	
                terraform apply -var-file ./terraform.tfvars --auto-approve; 	
            """	
        }	

        stage("Upload To S3") {		

            sh """		
                cd data-science; 		
                python3 -m venv s3; 		
                bash ./s3/bin/activate;  		
                pip3 install -r scrape_requirements.txt;		
                cd notebooks;		
                python3 upload_dataset.py; 				
            """		
        }

        stage("Invoke NLP Pipeline") {
            build job: "data-pipeline", parameters: [], wait: false
        }
    } finally {
        sh "rm -rf *"
    }
} 