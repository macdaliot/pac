package service

import (
	"github.com/PyramidSystemsInc/go/files"
)

func CreateDynamoDbTs(filePath string, config map[string]string) {
	const template = `import { Database } from './db.interface';
	import AWS = require('aws-sdk');
	import awsSdkConfig = require('../config/awsSdkConfig');
	import * as uuidv4 from 'uuid/v4'
	
	/* need configMap here */
	/* TODO need to refactor all the dynamodb */
	
	export class DynamoDB implements Database {
		dbInstance: AWS.DynamoDB;
		constructor() {
			AWS.config.update(awsSdkConfig.local);
			this.dbInstance = new AWS.DynamoDB({ apiVersion: '2012-10-08' });
		}
	
		/* TODO: need to create a generalize query function 
			this function should replace getById and getByParams
		*/
		/* rewrite to use documentClient */
		query(params) {
		}
	
		update(params: any, object: any) {
		}
	
		delete(id) {
		}
	
		create = async (object: any) => {
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
			try {
				return await this.dbInstance.batchWriteItem(params).promise();
			} catch (err) {
				throw err;
			}
		}
	
		async getById(id: string) {
			const whereClause = this.buildGetByIdParams(id);
			try {
				return await this.dbInstance.scan(whereClause).promise();
			} catch (err) {
				throw err;
			}
	
		}
	
		async getByQuery(query: any) {
			const whereClause = this.buildQueryStringParams(query);
			try {
				return await this.dbInstance.scan(whereClause).promise();
			} catch (err) {
				throw err;
			}
		}
	
		buildGetByIdParams = (id: string) => {
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
	
		buildQueryStringParams = (query: any) => {
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
	}`
	files.CreateFromTemplate(filePath, template, config)
}
