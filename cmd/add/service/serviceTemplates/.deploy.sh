#! /bin/bash
## Perform setup
SERVICE_NAME=$(sed -e 's/.*\///g' <<< $(pwd))

## Set Terraform environment variables
export TF_IN_AUTOMATION="NON_EMPTY_VALUE"

## Download Terraform
TF_VERSION="0.11.13"
wget https://releases.hashicorp.com/terraform/0.11.13/terraform_"$TF_VERSION"_linux_amd64.zip
unzip terraform_"$TF_VERSION"_linux_amd64.zip
ln -s /usr/bin terraform

## Navigate to Terraform directory
cd ./terraform

## Copy providers template to Terraform file
cp -a providers.tf.tpl providers.tf

## Replace ProjectName variable in providers.tf file
sed -i "s/{{ .ProjectName }}/{{.projectName}}/g" providers.tf

## Copy lambda template to Terraform file
cp -a lambda.tf.tpl lambda.tf

## Replace ProjectName variable in lambda.tf file
sed -i "s/{{ .ProjectName }}/{{.projectName}}/g" lambda.tf

## Replace ResourceName variable in lambda.tf file
sed -i "s/{{ .ResourceName }}/$SERVICE_NAME/g" lambda.tf

## Initialize terraform
terraform init -input=false

## Save tf plan
terraform plan -out=tfplan -input=false

## Apply tf plan
terraform apply -input=false tfplan

#################
#     OLD       #
#################
# # Perform setup
# SERVICE_NAME=$(sed -e 's/.*\///g' <<< $(pwd))
# FULL_SERVICE_NAME=pac-{{.projectName}}-i-"$SERVICE_NAME"
# # If AWS resources already exist...
# if aws elbv2 describe-target-groups --names $FULL_SERVICE_NAME --region us-east-2; then
#   echo "It appears the Lambda function exists. Updating the code in case there are any changes"
#   # Update the Lambda function code
#   aws lambda update-function-code --function-name $FULL_SERVICE_NAME --zip-file fileb://function.zip
# else
#   # Retrieve variables
#   LOAD_BALANCER_ARN=$(cat ../../.pac | jq '.loadBalancerArn')
#   LOAD_BALANCER_ARN=$(sed -e 's/^"//g' -e 's/"$//g' <<< $LOAD_BALANCER_ARN)
#   LISTENER_ARN=$(cat ../../.pac | jq '.listenerArn')
#   LISTENER_ARN=$(sed -e 's/^"//g' -e 's/"$//g' <<< $LISTENER_ARN)
  
#   # Create Target Group
#   TARGET_GROUP_ARN=$(aws elbv2 create-target-group --name "$FULL_SERVICE_NAME" --target-type lambda --region us-east-2 | jq '.TargetGroups[0].TargetGroupArn')
#   TARGET_GROUP_ARN=$(sed -e 's/^"//g' -e 's/"$//g' <<< $TARGET_GROUP_ARN)
#   echo "INFO (1/6 Completed): Created Target Group"

#   # Create Lambda Function
#   LAMBDA_ARN=$(aws lambda create-function --function-name "$FULL_SERVICE_NAME" --runtime nodejs8.10 --role arn:aws:iam::118104210923:role/service-role/god --handler lambda.handler --zip-file fileb://function.zip --region us-east-2 | jq '.FunctionArn')
#   LAMBDA_ARN=$(sed -e 's/^"//g' -e 's/"$//g' <<< $LAMBDA_ARN)
#   aws lambda tag-resource --resource $LAMBDA_ARN --tags pac-project-name={{.projectName}}
#   echo "INFO (2/6 Completed): Created Lambda Function"

#   # Adjust Lambda permissions
#   aws lambda add-permission --function-name "$FULL_SERVICE_NAME" --statement-id "$FULL_SERVICE_NAME"-grant-elb-access --action lambda:* --principal elasticloadbalancing.amazonaws.com --region us-east-2
#   echo "INFO (3/6 Completed): Added ELB permission to Lambda Function"

#   # Register Lambda to Target Group
#   aws elbv2 register-targets --target-group-arn "$TARGET_GROUP_ARN" --targets Id="$LAMBDA_ARN" --region us-east-2
#   echo "INFO (4/6 Completed): Registered the Lambda Function to the Target Group"

#   # Create ELB Rule
#   NEW_PRIORITY=$(aws elbv2 describe-rules --listener-arn "$LISTENER_ARN" --region us-east-2 | jq '.Rules' | jq length)
#   aws elbv2 create-rule --region us-east-2 --cli-input-json '{ "ListenerArn": "'"$LISTENER_ARN"'", "Priority": '"$NEW_PRIORITY"', "Conditions": [ { "Field": "path-pattern", "Values": [ "/api/'"$SERVICE_NAME"'" ] } ], "Actions": [ { "TargetGroupArn": "'"$TARGET_GROUP_ARN"'", "Type": "forward", "Order": 1 } ] }'
#   aws elbv2 create-rule --region us-east-2 --cli-input-json '{ "ListenerArn": "'"$LISTENER_ARN"'", "Priority": '"$((NEW_PRIORITY+1))"', "Conditions": [ { "Field": "path-pattern", "Values": [ "/api/'"$SERVICE_NAME"'/*" ] } ], "Actions": [ { "TargetGroupArn": "'"$TARGET_GROUP_ARN"'", "Type": "forward", "Order": 1 } ] }'
#   echo "INFO (5/6 Completed): Created Rule in ELB"

#   # Create DynamoDB Table
#   DYNAMO_TABLE_ARN=$(aws dynamodb create-table --cli-input-json file://dynamoConfig.json --region us-east-2 | jq '.TableDescription.TableArn')
#   DYNAMO_TABLE_ARN=$(sed -e 's/^"//g' -e 's/"$//g' <<< $DYNAMO_TABLE_ARN)
#   sleep 15
#   aws dynamodb tag-resource --resource-arn $DYNAMO_TABLE_ARN --tags Key=pac-project-name,Value={{.projectName}}
#   echo "INFO (6/6 Completed): Created DynamoDB table"
# fi