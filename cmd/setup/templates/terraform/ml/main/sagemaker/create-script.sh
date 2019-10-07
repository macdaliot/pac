#!/bin/bash
cat << EOF > ./getGithubCreds
import sys
import json
import urllib.parse
functionName = sys.argv[1]
fil = open(sys.argv[1], "r")
awsResponse =fil.read()
fil.close()
fil.__del__()

responseJson = json.loads(awsResponse)

secretString = responseJson['Parameter']['Value']
print(urllib.parse.quote(str(secretString)))
EOF
pip3 install awscli urllib3
aws ssm get-parameter --with-decryption --name /pac/github/username > username.json
USERNAME=`python3 getGithubCreds username.json`; rm username.json
aws ssm get-parameter --with-decryption --name /pac/github/password > password.json
PASSWORD=`python3 getGithubCreds password.json`; rm password.json
cd /home/ec2-user/SageMaker
git clone https://$USERNAME:$PASSWORD@github.com/PyramidSystemsInc/{{ .projectName }}.git
chown -R ec2-user {{ .projectName }}