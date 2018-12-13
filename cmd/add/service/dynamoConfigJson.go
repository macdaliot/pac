package service

import (
	"github.com/PyramidSystemsInc/go/files"
)

func CreateDynamoConfigJson(filePath string, config map[string]string) {
  const template = `{
  "TableName": "pac-{{.serviceName}}-dev",
  "KeySchema": [
    {
      "AttributeName": "id",
      "KeyType": "HASH"
    }
  ],
  "AttributeDefinitions": [
    {
      "AttributeName": "id",
      "AttributeType": "S"
    }
  ],
  "ProvisionedThroughput": {
    "ReadCapacityUnits": 5,
    "WriteCapacityUnits": 5
  }
}
`
  files.CreateFromTemplate(filePath, template, config)
}
