package service

import (
	"github.com/PyramidSystemsInc/go/files"
)

func CreateDefaultServiceTs(filePath string) {
	const template = `import { Database } from '../database/db.interface';
	import { DynamoDB } from '../database/dynamo.db';
	import { Service } from './service.interface';
	export class DefaultService implements Service {
		constructor(public db = new DynamoDB()) { }
		get = async (query: any) => {
			try {
				let dbResponse = await this.db.getByQuery(query);
				return dbResponse.Items
			} catch (err) {
				throw err;
			}
		}
	
		getById = async (id: string) => {
			try {
				let dbResponse = await this.db.getById(id);
				return dbResponse.Items;
			} catch (err) {
				throw err;
			}
		}
	
		post = async (body) => {
			try {
				let dbResponse = await this.db.create(body);
				return dbResponse;
			} catch (err) {
				throw err;
			}
		}
	}`
	files.CreateFromTemplate(filePath, template, nil)
}
