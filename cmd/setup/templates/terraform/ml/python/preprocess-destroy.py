import json
import sys

modelName = sys.arg[1]
model_data_url = sys.arg[2]
source_directory = sys.argv[3] 

fil=open("./terraform/" + environment + "../terraform/ml/python/terraform.tfvars", "a+")

fil.write("model_name  =  \"" + modelName + "\"\r\n")
fil.write("model_data_url  =  \"" + modelData + "\"\r\n")
fil.write("source_directory  =  \"" + sourceDir + "\"\r\n")

fil.close