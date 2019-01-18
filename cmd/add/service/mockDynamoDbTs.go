package service

import (
	"github.com/PyramidSystemsInc/go/files"
)

func CreateMockDynamoDbTs(filePath string) {
	const template = `import * as AWS from 'aws-sdk';
    import * as AWSMock from 'aws-sdk-mock';
    
    export class MockDynamoDB {
        dbInstance: AWS.DynamoDB;
        _mockParams = {
            TableName: 'mydynamo',
            Key: {
                id: { S: 'myid' },
            },
        };
        _mockScanInput: AWS.DynamoDB.ScanInput = {
            TableName: 'mock'
        };
        _mockWriteInput = {
            "RequestItems": {
                "pac-{{.projectName}}-i-{{.serviceName}}": [
                    {
                        "PutRequest": {
                            "Item": {
                                "id": {
                                    "S": 'asdfasdf'
                                }
                            }
                        }
                    }
                ]
            }
        }
    
    
        constructor() {
            const _mockDBResponse = mockDBResponse;
            AWSMock.mock('DynamoDB', 'getItem', function (params, callback) {
                callback(null, _mockDBResponse);
            });
            AWSMock.mock('DynamoDB', 'batchWriteItem', function (params, callback) {
                callback(null, _mockDBResponse);
            });
            this.dbInstance = new AWS.DynamoDB();
        }
        getById = async id => await this.dbInstance.getItem(this._mockParams).promise();
        getByQuery = async query => await this.dbInstance.getItem(this._mockParams).promise();
        buildGetByIdParams = (id) => this._mockScanInput;
        buildQueryStringParams = (query) => this._mockScanInput;
        create = async body => await this.dbInstance.batchWriteItem(this._mockWriteInput).promise();
        query;
        update;
        delete;
    }
    export const mockDBobject = { test: 'value' }
    export const mockDBResponse = { Items: mockDBobject }
    `
	files.CreateFromTemplate(filePath, template, nil)
}
