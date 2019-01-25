import { Database } from './db.interface';
import { config, DynamoDB as Dynamo } from 'aws-sdk';
import { projectName, serviceName, awsConfig } from '../config'
import * as uuidv4 from 'uuid/v4'
import * as _ from 'lodash';
import * as AWS from 'aws-sdk';

const _tableName = `pac-${projectName}-i-${serviceName}`;
export class DynamoDB implements Database {
    /* DocumentClient API: https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/AWS/DynamoDB/DocumentClient.html */
    dbInstance: Dynamo.DocumentClient;
    constructor() {
        let env = process.env.ENVIRONMENT || "cloud";
        config.update(awsConfig[env]);
        this.dbInstance = new Dynamo.DocumentClient({ apiVersion: '2012-10-08' });
    }

    query = async (params) => {
        let whereClause;
        try {
            if (_.isEmpty(params)) {
                whereClause = this.buildGetParams(params);
                return await this.dbInstance.scan(whereClause).promise()
            } else {
                whereClause = this.buildQueryParams(params);
                return await this.dbInstance.query(whereClause).promise();
            }

        } catch (err) {
            throw err;
        }
    }

    update = async (params: any, object: any) => {
        const whereClause = this.buildUpdateParams(params);
        try {
            return await this.dbInstance.update(whereClause).promise();
        } catch (err) {
            throw err;
        }
    }

    delete = async (query) => {
        const whereClause = this.buildDeleteParams(query);
        try {
            return await this.dbInstance.delete(whereClause).promise();
        } catch (err) {
            throw err;
        }
    }

    create = async (object: any): Promise<AWS.DynamoDB.DocumentClient.PutItemOutput> => {
        const whereClause = this.buildCreateParams(object);
        try {
            return await this.dbInstance.put(whereClause).promise();
        } catch (err) {
            throw err;
        }
    }

    buildGetParams = (query): AWS.DynamoDB.DocumentClient.GetItemInput => {
        return {
            TableName: _tableName,
            Key: query
        }
    }

    /* TODO: only support single level json */
    buildQueryParams = (query): AWS.DynamoDB.DocumentClient.QueryInput => {
        const queryInput: AWS.DynamoDB.DocumentClient.QueryInput = {
            TableName: _tableName
        }
        queryInput.KeyConditionExpression = this.createKeyConditionExpression(query);
        queryInput.ExpressionAttributeValues = this.createExpressionAttributeValues(query);
        return queryInput;
    }

    buildCreateParams = (body): AWS.DynamoDB.DocumentClient.PutItemInput => {
        return {
            TableName: _tableName,
            Item: body
        }
    }

    buildDeleteParams = (query): AWS.DynamoDB.DocumentClient.DeleteItemInput => {
        return {
            TableName: _tableName,
            Key: query
        }
    }

    buildUpdateParams = (query): AWS.DynamoDB.DocumentClient.UpdateItemInput => {
        const params: AWS.DynamoDB.DocumentClient.UpdateItemInput = {
            TableName: _tableName,
            Key: query,
            ReturnValues: "UPDATED_NEW"
        }

        /* TODO: construct the string for update Expression and expressionAttributeValues */
        params.UpdateExpression = '';
        params.ExpressionAttributeValues = {};
        return params;
    }

    private createKeyConditionExpression = query => Object.keys(query).map(key => `${key} = :_${key}`).join(' and ');
    private createExpressionAttributeValues = query => Object.keys(query).reduce(
        (accumulator, key) => {
            accumulator[`:_${key}`] = query[key]
            return accumulator;
        }
        , {});
}