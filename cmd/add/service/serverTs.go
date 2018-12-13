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
  endpoint: 'http://localhost:8000'
});
const dynamo: any = new AWS.DynamoDB();

app.get('/api/{{.serviceName}}', function(request, response) {
  let params: any = {
    ExpressionAttributeValues: {
      ":a": {
        S: "Person"
      }
    },
    FilterExpression: "Artist = :a",
    TableName: "pac-{{.serviceName}}-dev"
  };
  dynamo.scan(params, function(err, data) {
    if (!err) {
      response.send(data);
    } else {
      response.send("ERROR. See console");
      console.error(err);
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
              "AlbumTitle": {
                "S": "Something Cool"
              },
              "Artist": {
                "S": "Person"
              },
              "SongTitle": {
                "S": "14th Symphony in D-Minor"
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

app.listen(port);
`
  files.CreateFromTemplate(filePath, template, config)
}
