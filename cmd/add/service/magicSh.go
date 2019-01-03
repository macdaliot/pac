package service

import (
	"github.com/PyramidSystemsInc/go/commands"
	"github.com/PyramidSystemsInc/go/files"
)

func CreateMagicSh(filePath string, serviceName string) {
	const template = `#! /bin/bash

# Perform Setup
SERVICE_NAME=$(sed -e 's/.*\///g' <<< $(pwd))
FULL_SERVICE_NAME=pac-"$SERVICE_NAME"
LOAD_BALANCER_ARN=$(cat ../../.pac | jq '.loadBalancerArn')
LOAD_BALANCER_ARN=$(sed -e 's/^"//g' -e 's/"$//g' <<< $LOAD_BALANCER_ARN)
LISTENER_ARN=$(cat ../../.pac | jq '.listenerArn')
LISTENER_ARN=$(sed -e 's/^"//g' -e 's/"$//g' <<< $LISTENER_ARN)

# Create Target Group
TARGET_GROUP_ARN=$(aws elbv2 create-target-group --name "$FULL_SERVICE_NAME" --target-type lambda --region us-east-2 | jq '.TargetGroups[0].TargetGroupArn')
TARGET_GROUP_ARN=$(sed -e 's/^"//g' -e 's/"$//g' <<< $TARGET_GROUP_ARN)

# Create Lambda Function
npm i
npx tsc server.ts
echo $(sed -e 's/awsSdkConfig.local/awsSdkConfig.cloud/g' -e 's/app.listen(port);/module.exports = app;/g' server.js) > server.js
npx claudia generate-serverless-express-proxy --express-module server >> /dev/null
zip -r function awsSdkConfig.js lambda.js server.js node_modules >> /dev/null
LAMBDA_ARN=$(aws lambda create-function --function-name "$FULL_SERVICE_NAME" --runtime nodejs8.10 --role arn:aws:iam::118104210923:role/service-role/god --handler lambda.handler --zip-file fileb://function.zip --region us-east-2 | jq '.FunctionArn')
LAMBDA_ARN=$(sed -e 's/^"//g' -e 's/"$//g' <<< $LAMBDA_ARN)

# Adjust Lambda Permissions
aws lambda add-permission --function-name "$FULL_SERVICE_NAME" --statement-id "$FULL_SERVICE_NAME"-grant-elb-access --action lambda:* --principal elasticloadbalancing.amazonaws.com --region us-east-2 >> /dev/null

# Register Lambda to Target Group
aws elbv2 register-targets --target-group-arn "$TARGET_GROUP_ARN" --targets Id="$LAMBDA_ARN" --region us-east-2 >> /dev/null

# Create ELB Rule
NEW_PRIORITY=$(aws elbv2 describe-rules --listener-arn "$LISTENER_ARN" --region us-east-2 | jq '.Rules' | jq length)
aws elbv2 create-rule --region us-east-2 --cli-input-json '{ "ListenerArn": "'"$LISTENER_ARN"'", "Priority": '"$NEW_PRIORITY"', "Conditions": [ { "Field": "path-pattern", "Values": [ "/api/'"$SERVICE_NAME"'" ] } ], "Actions": [ { "TargetGroupArn": "'"$TARGET_GROUP_ARN"'", "Type": "forward", "Order": 1 } ] }'

# Create DynamoDB Table
aws dynamodb create-table --cli-input-json file://dynamoConfig.json --region us-east-2 >> /dev/null
`
	files.CreateFromTemplate(filePath, template, nil)
	commands.Run("chmod 755 " + serviceName + "/magic.sh", "")
}
