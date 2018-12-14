package service

import (
	"github.com/PyramidSystemsInc/go/files"
)

func CreateServerTs(filePath string, config map[string]string) {
  const template = `const AWS: any = require('aws-sdk');
const express: any = require('express');
const bodyParser: any = require('body-parser');
const uuidv4: any = require('uuid/v4');
const port: number = 3000;
const serviceName: string = '{{.serviceName}}';

const app: any = express();
app.use(bodyParser.json());

AWS.config.update({
  region: 'local',
  endpoint: 'http://pac-db-local:8000'
});
const dynamo: any = new AWS.DynamoDB();

app.get('/api/{{.serviceName}}/:id', function(request, response) {
  let whereClause: any = buildGetByIdParams(request.params.id);
  dynamo.scan(whereClause, function(err, data) {
    if (!err) {
      response.send(data);
    } else {
      response.status(500).send(err);
    }
  });
});

app.get('/api/{{.serviceName}}', function(request, response) {
  let whereClause: any = buildQueryStringParams(request.query);
  dynamo.scan(whereClause, function(err, data) {
    if (!err) {
      response.send(data);
    } else {
      response.status(500).send(err);
    }
  });
});

app.post('/api/{{.serviceName}}', function(request, response) {
  let params: any = {
    "RequestItems": {
      "pac-{{.serviceName}}-dev": [
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
  dynamo.batchWriteItem(params, function(err, data) {
    if (!err) {
      response.send(data);
    } else {
      response.status(500).send(err);
    }
  });
});

let buildQueryStringParams = function(query) {
  let params: any = {
    TableName: 'pac-{{.serviceName}}-dev'
  };
  let queryKeys: any = Object.keys(query);
  if (queryKeys.length > 0) {
    params.ExpressionAttributeValues = {};
    params.FilterExpression = "";
    queryKeys.forEach(function(key, i) {
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

let buildGetByIdParams = function(id) {
  let params: any = {
    ExpressionAttributeValues: {},
    FilterExpression: 'id = :id',
    TableName: 'pac-{{.serviceName}}-dev'
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
