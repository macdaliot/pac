#! /bin/bash
# Perform setup
SERVICE_NAME=$(sed -e 's/.*\///g' <<< $(pwd))
FULL_SERVICE_NAME=pac-{{.projectName}}-i-"$SERVICE_NAME"
# If AWS resources already exist...
if aws elbv2 describe-target-groups --names $FULL_SERVICE_NAME --region {{ .region }}; then
  echo "It appears the Lambda function exists. Updating the code in case there are any changes"
  # Update the Lambda function code
  aws lambda update-function-code --function-name $FULL_SERVICE_NAME --zip-file fileb://function.zip
else
  # Retrieve variables
  LOAD_BALANCER_ARN=$(cat ../../.pac | jq '.loadBalancerArn')
  LOAD_BALANCER_ARN=$(sed -e 's/^"//g' -e 's/"$//g' <<< $LOAD_BALANCER_ARN)
  LISTENER_ARN=$(cat ../../.pac | jq '.listenerArn')
  LISTENER_ARN=$(sed -e 's/^"//g' -e 's/"$//g' <<< $LISTENER_ARN)

  # Create Target Group
  TARGET_GROUP_ARN=$(aws elbv2 create-target-group --name "$FULL_SERVICE_NAME" --target-type lambda --region {{ .region }} | jq '.TargetGroups[0].TargetGroupArn')
  TARGET_GROUP_ARN=$(sed -e 's/^"//g' -e 's/"$//g' <<< $TARGET_GROUP_ARN)
  echo "INFO (1/5 Completed): Created Target Group"

  # Create Lambda Function
  LAMBDA_ARN=$(aws lambda create-function --function-name "$FULL_SERVICE_NAME" --runtime nodejs8.10 --role arn:aws:iam::118104210923:role/service-role/god --handler lambda.handler --zip-file fileb://function.zip --region {{ .region }} | jq '.FunctionArn')
  LAMBDA_ARN=$(sed -e 's/^"//g' -e 's/"$//g' <<< $LAMBDA_ARN)
  aws lambda tag-resource --resource $LAMBDA_ARN --tags pac-project-name=testa
  echo "INFO (2/5 Completed): Created Lambda Function"

  # Adjust Lambda permissions
  aws lambda add-permission --function-name "$FULL_SERVICE_NAME" --statement-id "$FULL_SERVICE_NAME"-grant-elb-access --action lambda:* --principal elasticloadbalancing.amazonaws.com --region {{ .region }}
  echo "INFO (3/5 Completed): Added ELB permission to Lambda Function"

  # Register Lambda to Target Group
  aws elbv2 register-targets --target-group-arn "$TARGET_GROUP_ARN" --targets Id="$LAMBDA_ARN" --region {{ .region }}
  echo "INFO (4/5 Completed): Registered the Lambda Function to the Target Group"

  # Create ELB Rule
  NEW_PRIORITY=$(aws elbv2 describe-rules --listener-arn "$LISTENER_ARN" --region {{ .region }} | jq '.Rules' | jq length)
  aws elbv2 create-rule --region {{ .region }} --cli-input-json '{ "ListenerArn": "'"$LISTENER_ARN"'", "Priority": '"$NEW_PRIORITY"', "Conditions": [ { "Field": "path-pattern", "Values": [ "/api/'"$SERVICE_NAME"'" ] } ], "Actions": [ { "TargetGroupArn": "'"$TARGET_GROUP_ARN"'", "Type": "forward", "Order": 1 } ] }'
  aws elbv2 create-rule --region {{ .region }} --cli-input-json '{ "ListenerArn": "'"$LISTENER_ARN"'", "Priority": '"$((NEW_PRIORITY+1))"', "Conditions": [ { "Field": "path-pattern", "Values": [ "/api/'"$SERVICE_NAME"'/*" ] } ], "Actions": [ { "TargetGroupArn": "'"$TARGET_GROUP_ARN"'", "Type": "forward", "Order": 1 } ] }'
  echo "INFO (5/5 Completed): Created Rule in ELB"
fi
