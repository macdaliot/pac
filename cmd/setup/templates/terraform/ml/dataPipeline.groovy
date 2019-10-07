node ("scikit-sagemaker") {
    stage ("SCM Checkout") {
        try{
            git branch: 'data-science-pipeline', credentialsId: 'jenkins-git', url: 'git@github.com:PyramidSystemsInc/bdso-fork.git'
            
        } catch (Exception ex) {
            print("Failed to checkout {{ .projectName }}")
            print(ex.toString())
        }
    }
    RESOLVED_SOURCE_VERSION=sh(returnStdout: true, script: "git rev-parse HEAD")

    stage ("Scrape Data") {
        withCredentials([string(credentialsId: "TMDB_KEY", variable: "TMDB_KEY")]){
            sh """
                cd data-science;
                python3 -m venv scrape_env;
                sh ./scrape_env/bin/activate;
                pip3 install -r scrape_requirements.txt --user;
                cd notebooks; 
                python3 upload_dataset.py;
            """
        }
    }

    stage ("Deduplicate & Link Records") {
        sh """
            export ENVIRONMENT="$ENVIRONMENT"
            python3 -m venv scrape_env;
            bash ./scrape_env/bin/activate;
            cd data-science;
            pip3 install -r build-requirements.txt --user;
            pip3 install aws sagemaker --user;
            python3 deduplication/build.py
        """
    }
}