package service

import (
	"github.com/PyramidSystemsInc/go/files"
)

// CreateServerTs creates a server.ts file based on the configuration passed in
func CreateServerTs(filePath string, config map[string]string) {
	const template = `import AWS = require('aws-sdk');
  import express = require('express');
  import bodyParser = require('body-parser');
  import uuidv4 = require('uuid/v4');
  const port: number = 3000;
  const serviceName: string = '{{.serviceName}}';
  
  const app = express();
  app.use(bodyParser.json());
  
  import awsSdkConfig = require('./awsSdkConfig');
  AWS.config.update(awsSdkConfig.local);
  const dynamo = new AWS.DynamoDB({ apiVersion: '2012-10-08' });
  
  app.get('/api/{{.serviceName}}/:id', (request, response) => {
    const whereClause = buildGetByIdParams(request.params.id);
    dynamo.scan(whereClause, function(err, data) {
      if (!err) {
        response.send(data);
      } else {
        response.status(500).send(err);
      }
    });
  });
  
  app.get('/api/{{.serviceName}}', (request, response) => {
    const whereClause = buildQueryStringParams(request.query);
    dynamo.scan(whereClause, function(err, data) {
      if (!err) {
        response.send(data);
      } else {
        response.status(500).send(err);
      }
    });
  });
  
  app.post('/api/{{.serviceName}}', (request, response) => {
    let params: any = {
      "RequestItems": {
        "pac-{{.projectName}}-i-{{.serviceName}}": [
          {
            "PutRequest": {
              "Item": {
                "id": {
                  "S": uuidv4()
                },
                "UserName": {
                  "S": "bobby"
                },
                "FirstName": {
                  "S": "Robert"
                },
                "LastName": {
                  "S": "Jones"
                }
              }
            }
          }
        ]
      }
    };
    dynamo.batchWriteItem(params, (err, data) => {
      if (!err) {
        response.send(data);
      } else {
        response.status(500).send(err);
      }
    });
  });
  
  let buildQueryStringParams = (query: any) => {
    let params: AWS.DynamoDB.ScanInput = {
      TableName: 'pac-{{.projectName}}-i-{{.serviceName}}'
    };
    let queryKeys: any = Object.keys(query);
    if (queryKeys.length > 0) {
      params.ExpressionAttributeValues = {};
      params.FilterExpression = "";
      queryKeys.forEach((key: string, i: number) => {
        params.ExpressionAttributeValues[":" + key] = {
          S: query[key]
        };
        if (i != 0) {
          params.FilterExpression += " and ";
        }
        let keyUpperCase: string = key.charAt(0).toUpperCase() + key.slice(1);
        params.FilterExpression += keyUpperCase + " = :" + key;
      });
    }
    return params;
  };
  
  let buildGetByIdParams = (id: string) => {
    let params: AWS.DynamoDB.ScanInput = {
      ExpressionAttributeValues: {},
      FilterExpression: 'id = :id',
      TableName: 'pac-{{.projectName}}-i-{{.serviceName}}'
    };
    params.ExpressionAttributeValues[":id"] = {
      S: id
    };
    return params;
  };
  
  app.listen(port);  
`
	files.CreateFromTemplate(filePath, template, config)
}
