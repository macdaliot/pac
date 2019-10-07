import sys

functionName = sys.argv[1]

os.system("aws ssm get-parameter --name '/pac/{{ .projectName }}/lambda/" + functionName + "/version' > tmdbkeyResponse")


responseFile = open("tmdbkeyResponse","r");

awsResponse = responseFile.read()
responseFile.close()
responseFile.__del__()

responseJson = json.loads(awsResponse)

secretString = responseJson["Parameter"]["Value"]
# json.loads(json.loads(json.dumps(responseJson["SecretString"], sort_keys=True, indent=4)))
print(str(secretString))


