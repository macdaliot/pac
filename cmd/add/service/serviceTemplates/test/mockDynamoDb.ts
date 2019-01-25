import * as AWS from 'aws-sdk';
import * as AWSMock from 'aws-sdk-mock';
import { Database } from '../src/database/db.interface';
import { DynamoDB } from '../src/database/dynamo.db';

export class MockDynamoDB extends DynamoDB {
    dbInstance: AWS.DynamoDB.DocumentClient;
    constructor() {
        super();
        mockDbResponses();
        this.dbInstance = new AWS.DynamoDB.DocumentClient();
    }
    create = async body => await this.dbInstance.put({ TableName: 'test', Item: {} }).promise();
    query = async query => await this.dbInstance.query({ TableName: 'test' }).promise();
    update;
    delete;
}

export const mockDBobject = { test: 'value' }
export const mockDBResponse = { Items: mockDBobject }
export const mockDbResponses = () => {
    AWSMock.mock('DynamoDB.DocumentClient', 'query', function (params, callback) {
        callback(null, mockDBResponse);
    });
    AWSMock.mock('DynamoDB.DocumentClient', 'update', function (params, callback) {
        callback(null, mockDBResponse);
    });
    AWSMock.mock('DynamoDB.DocumentClient', 'delete', function (params, callback) {
        callback(null, {});
    });
    AWSMock.mock('DynamoDB.DocumentClient', 'put', function (params, callback) {
        callback(null, mockDBResponse);
    });
}