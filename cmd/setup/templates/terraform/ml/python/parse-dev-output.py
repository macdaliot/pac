import json
import sys
import os

envs = {
    "qa"  : "dev",
    "e2e"  : "dev",
    "prod" : "prod"
}
configuration = open(sys.argv[1], "r").read()
parameters = json.loads(configuration)['Parameters']

environment = envs.get(parameters['Environment'])
modelName = parameters['ModelName']
modelData = parameters['ModelData']
sourceDir = parameters['SourceDirectory']

fil=open("main/terraform.tfvars", "a+")

fil.write("model_name  =  \"" + modelName + "\"\r\n")
fil.write("model_data_url  =  \"" + modelData + "\"\r\n")
fil.write("source_directory  =  \"" + sourceDir + "\"\r\n")

fil.close

print(modelData)