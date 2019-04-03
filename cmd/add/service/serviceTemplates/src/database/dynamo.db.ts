import { Database } from './db.interface';
import { config, DynamoDB as Dynamo } from 'aws-sdk';
import { projectName, serviceName, awsConfig } from '../config'
import * as uuidv4 from 'uuid/v4'
import * as _ from 'lodash';
import * as AWS from 'aws-sdk';

type KeyValueDictionary = { [key: string]: any };
const _tableName = `pac-${projectName}-i-${serviceName}`;
export class DynamoDB implements Database {
    /* DocumentClient API: https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/AWS/DynamoDB/DocumentClient.html */
    dbInstance: Dynamo.DocumentClient;
    constructor() {
        let env = process.env.ENVIRONMENT || "cloud";
        config.update(awsConfig[env]);
        this.dbInstance = new Dynamo.DocumentClient({ apiVersion: '2012-10-08' });
    }

    query = async (params: KeyValueDictionary) => {
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
            console.log("error during query");
            console.log(err);
            throw err;
        }
    }

    update = async (params: KeyValueDictionary, object: KeyValueDictionary) => {
        const whereClause = this.buildUpdateParams(params);
        try {
            let response = await this.dbInstance.update(whereClause).promise();
            console.log(JSON.stringify(response));
            return response;
        } catch (err) {
            console.log("error during update");
            console.log(err);
            throw err;
        }
    }

    delete = async (query: KeyValueDictionary) => {
        const whereClause = this.buildDeleteParams(query);
        try {
            let response = await this.dbInstance.delete(whereClause).promise();
            console.log(JSON.stringify(response));
            return response;
        } catch (err) {
            console.log("error during delete");
            console.log(err);
            throw err;
        }
    }

    create = async (object: KeyValueDictionary): Promise<{}> => {
        try {
            const whereClause = this.buildCreateParams(object);
            return await this.dbInstance.put(whereClause).promise();
        }
        catch (err){
            console.log("error during create");
            console.log(err);
            throw err;
        }
    }

    buildGetParams = (query: KeyValueDictionary): AWS.DynamoDB.DocumentClient.GetItemInput => {
        return {
            TableName: _tableName,
            Key: query
        }
    }

    /* TODO: only support single level json */
    buildQueryParams = (query: KeyValueDictionary): AWS.DynamoDB.DocumentClient.QueryInput => {
        const queryInput: AWS.DynamoDB.DocumentClient.QueryInput = {
            TableName: _tableName
        }
        queryInput.KeyConditionExpression = this.createKeyConditionExpression(query);
        queryInput.ExpressionAttributeValues = this.createExpressionAttributeValues(query);
        return queryInput;
    }

    buildCreateParams = (body: KeyValueDictionary): AWS.DynamoDB.DocumentClient.PutItemInput => {
        return {
            TableName: _tableName,
            Item: body
        }
    }

    buildDeleteParams = (query: KeyValueDictionary): AWS.DynamoDB.DocumentClient.DeleteItemInput => {
        return ({
            TableName: _tableName,
            Key: query
        });
    }

    buildUpdateParams = (query: KeyValueDictionary): AWS.DynamoDB.DocumentClient.UpdateItemInput => {
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

    createKeyConditionExpression =
        (query: KeyValueDictionary) => Object.keys(query).map(key => `${key} = :_${key}`).join(' and ');
    createExpressionAttributeValues =
        (query: KeyValueDictionary) => Object.keys(query).reduce(
            (accumulator, key) => {
                accumulator[`:_${key}`] = query[key]
                return accumulator;
            }, {} as KeyValueDictionary);
}