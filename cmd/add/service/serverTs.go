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
const serviceName: string = 'users';

const app: any = express();
app.use(bodyParser.json());

AWS.config.update({
  region: 'local',
  endpoint: 'http://localhost:8000'
});
const dynamo: any = new AWS.DynamoDB();

app.get('/api/users/:id', function(request, response) {
  let whereClause: any = buildGetByIdParams(request.params.id);
  dynamo.scan(whereClause, function(err, data) {
    if (!err) {
      response.send(data);
    } else {
      response.send("ERROR. See console");
      console.error(err);
    }
  });
});

app.get('/api/users', function(request, response) {
  let whereClause: any = buildQueryStringParams(request.query);
  dynamo.scan(whereClause, function(err, data) {
    if (!err) {
      response.send(data);
    } else {
      response.send("ERROR. See console");
      console.error(err);
    }
  });
});

app.get('/post/users', function(request, response) {
  let params: any = {
    "RequestItems": {
      "pac-users-dev": [
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
      response.send("ERROR. See console");
      console.error(err);
    }
  });
});

let buildQueryStringParams = function(query) {
  let params: any = {
    TableName: 'pac-users-dev'
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
    TableName: 'pac-users-dev'
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
