import { Database } from '../database/db.interface';
import { DynamoDB } from '../database/dynamo.db';
import { Service } from './service.interface';
export class DefaultService implements Service {
    constructor(public db = new DynamoDB()) { }
    get = async (query: any) => {
        try {
            let dbResponse = await this.db.query(query);
            return dbResponse.Items
        } catch (err) {
            throw err;
        }
    }

    getById = async (id: string) => {
        try {
            let dbResponse = await this.db.query({ id });
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
}