import json
import sys

configuration = open(sys.argv[1], "r").read()
parameters = json.loads(configuration)

print(parameters["PrimaryContainer"]["ModelDataUrl"])