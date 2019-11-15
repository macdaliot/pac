#!/bin/bash

# Set configuration variables
BOOTSTRAP_GITHUB_REPO="PyramidSystemsInc/bdso-phase-two-prep"
ARCH="amd64"
OS="linux"
GOLANG_VERSION="1.13.1"
NODE_VERSION="10.16.3"
TF_VERSION="0.12.7"

# Install necessary programs with `yum` package manager
yum update -y
yum install git docker java-1.8.0-openjdk -y
systemctl start docker

# Create Bash function needed later
cat << EOF >> /root/.bashrc
export TF_VAR_aws_access_key_id=${aws_access_key_id}
export TF_VAR_aws_account_number=${aws_account_number}
export TF_VAR_aws_secret_access_key=${aws_secret_access_key}
export TF_VAR_github_auth=${github_auth}
export TF_VAR_github_password=${github_password}
export TF_VAR_github_username=${github_username}
export TF_VAR_iam_user_password=${iam_user_password}
export TF_VAR_jenkins_password=${jenkins_password}
export TF_VAR_jenkins_username=${jenkins_username}
export TF_VAR_jwt_issuer=${jwt_issuer}
export TF_VAR_jwt_secret=${jwt_secret}
export TF_VAR_postgres_password=${postgres_password}
export TF_VAR_project_name=${project_name}
export TF_VAR_psi_aws_access_key_id=${psi_aws_access_key_id}
export TF_VAR_psi_aws_account_number=${psi_aws_account_number}
export TF_VAR_psi_aws_secret_access_key=${psi_aws_secret_access_key}
export TF_VAR_sonar_jdbc_password=${sonar_jdbc_password}
export TF_VAR_sonar_jdbc_username=${sonar_jdbc_username}
export TF_VAR_sonar_password=${sonar_password}
export TF_VAR_sonar_secret=${sonar_secret}
export TF_VAR_sonar_username=${sonar_username}

alias terraformPlan='terraform init; terraform plan -out=tfplan'
alias terraformApply='terraform init; terraform apply tfplan'
EOF
source /root/.bashrc

# Configure the AWS credentials
mkdir -p /root/.aws
touch /root/.aws/credentials
cat << EOF >> /root/.aws/credentials
[default]
aws_access_key_id = ${aws_access_key_id}
aws_secret_access_key = ${aws_secret_access_key}
region = us-east-1

[psi]
aws_access_key_id = ${psi_aws_access_key_id}
aws_secret_access_key = ${psi_aws_secret_access_key}
region = us-east-1
EOF
cat << EOF >> /root/.aws/config
[default]
region = us-east-1
EOF

# Install NodeJS
mkdir -p /root/Programs/node-js
pushd /root/Programs/node-js
wget https://nodejs.org/dist/v"$NODE_VERSION"/node-v"$NODE_VERSION"-"$OS"-x64.tar.xz
tar -xf node-v"$NODE_VERSION"-"$OS"-x64.tar.xz
rm node-v"$NODE_VERSION"-"$OS"-x64.tar.xz
mv node-v"$NODE_VERSION"-"$OS"-x64 v"$NODE_VERSION"
echo -e "export NODE_JS_HOME=/root/Programs/node-js/v$NODE_VERSION" >> /root/.bashrc
echo -e "export PATH=\$NODE_JS_HOME/bin/:\$PATH" >> /root/.bashrc
popd

# Install Terraform
mkdir -p /root/Programs/terraform/v"$TF_VERSION"
pushd /root/Programs/terraform/v"$TF_VERSION"
wget https://releases.hashicorp.com/terraform/"$TF_VERSION"/terraform_"$TF_VERSION"_"$OS"_"$ARCH".zip
unzip terraform_"$TF_VERSION"_"$OS"_"$ARCH".zip
rm terraform_"$TF_VERSION"_"$OS"_"$ARCH".zip
popd
echo -e "export TERRAFORM_HOME=/root/Programs/terraform/v$TF_VERSION" >> /root/.bashrc
echo -e "export PATH=\$TERRAFORM_HOME:\$PATH" >> /root/.bashrc

# Install PAC
mkdir -p /root/go/bin
pushd /root/go/bin
wget https://pac-binary-downloads.s3.us-east-2.amazonaws.com/pac
chmod 755 pac
popd
echo -e "export GO_BIN_PATH=/root/go/bin" >> /root/.bashrc
echo -e "export PATH=\$GO_BIN_PATH:\$PATH" >> /root/.bashrc

# Clone the project bootstrapper GitHub repository
source /root/.bashrc
pushd /root
git clone https://"$TF_VAR_github_auth"@github.com/"$BOOTSTRAP_GITHUB_REPO".git
popd
