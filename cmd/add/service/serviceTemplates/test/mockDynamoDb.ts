import * as AWS from 'aws-sdk';
import * as AWSMock from 'aws-sdk-mock';
import { Database } from '../src/database/db.interface';
import { DynamoDB } from '../src/database/dynamo.db';

export class MockDynamoDB implements Database {
    create = async body => ({});
    query = async query => (mockDBResponse);
    update = async body => ({});
    delete = async body => ({});
}
export class BrokenDynamoDB implements Database {
    create = async body => {throw new Error()};
    query = async query => {throw new Error()};
    update = async body => {throw new Error()};
    delete = async body => {throw new Error()};
}
export interface MockOptions {
    queryResponse?: any;
    updateResponse?: any;
    deleteResponse?: any;
    putResponse?: any;
}
export const mockDBobject = { test: 'value' }
export const mockDBResponse = { Items: mockDBobject }
export const mockDbResponses = (options: MockOptions = {}) => {
    AWSMock.mock('DynamoDB.DocumentClient', 'query', function (params, callback) {
        const error = (options.queryResponse instanceof Error) ? options.queryResponse : null;
        if (error)
            callback(error, null);
        else
            callback(null, mockDBResponse);
    });
    AWSMock.mock('DynamoDB.DocumentClient', 'scan', function (params, callback) {
        const error = (options.queryResponse instanceof Error) ? options.queryResponse : null;
        if (error)
            callback(error, null);
        else
            callback(null, mockDBResponse);
    });
    AWSMock.mock('DynamoDB.DocumentClient', 'update', function (params, callback) {
        const error = (options.updateResponse instanceof Error) ? options.updateResponse : null;
        if (error)
            callback(error, null);
        else
            callback(null, mockDBResponse);
    });
    AWSMock.mock('DynamoDB.DocumentClient', 'delete', function (params, callback) {
        const error = (options.deleteResponse instanceof Error) ? options.deleteResponse : null;
        if (error)
            callback(error, null);
        else
            callback(null, {});
    });
    AWSMock.mock('DynamoDB.DocumentClient', 'put', function (params, callback) {
        const error = (options.putResponse instanceof Error) ? options.putResponse : null;
        if (error)
            callback(error, null);
        else
            callback(null, {});
    });
}