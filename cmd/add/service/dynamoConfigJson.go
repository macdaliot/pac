package service

import (
  "github.com/PyramidSystemsInc/go/files"
)

// CreateDynamoConfigJSON creates a dynamo config for the local instance
func CreateDynamoConfigJSON(filePath string, config map[string]string) {
  const template = `{
  "TableName": "pac-{{.projectName}}-i-{{.serviceName}}",
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
