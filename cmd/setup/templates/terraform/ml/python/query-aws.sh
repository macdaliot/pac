#!/bin/bash
env=$1
project_name=$2
model_name="nlp"

aws sagemaker describe-endpoint --endpoint-name "$env-$project_name-$model_name" > endpoint-details.json
endpoint_config_name=`python3 getEndpointConfig.py endpoint-details.json`
aws sagemaker describe-endpoint-config --endpoint-config-name $endpoint_config_name > endpoint-config-details.json
modelName=`python3 getModelName.py endpoint-config-details.json`
aws sagemaker describe-model --model-name $modelName > model-details.json
submit_directory=`python3 getSubmitDirectory.py model-details.json`
model_data_url=`python3 getModelDataUrl.py model-details.json`

python3 preproccess-destroy.py model_name modelDataUrl submit_directory

rm endpoint-details.json endpoint-config-details.json model-details.json