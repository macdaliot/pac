import { DataMapper, ScanIterator } from '@aws/dynamodb-data-mapper';
import { Injectable } from '@pyramid-systems/core';
import { config, DynamoDB } from 'aws-sdk';
import { awsConfig } from './config';
import { Newable } from '../newable';
import { Repository } from '../repository';

interface IEntity {
  id: string
}

@Injectable()
export class DynamoRepository<TModel extends IEntity> implements Repository<TModel> {
  private mapper: DataMapper;

  constructor() {
    const env: string = process.env.ENVIRONMENT || 'cloud';
    config.update(awsConfig[env]);
    this.mapper = new DataMapper({
      client: new DynamoDB()
    });
  }

  async get(id: string, model: Newable<TModel>): Promise<TModel> {
    const getRequest: TModel = new model();
    getRequest.id = id;
    return await this.mapper.get<TModel>(getRequest);
  }

  // GET /api/fire?column=this&othercol=that
  // query || search (returns list of data rows that match on an arbitrary column)

  async scan(value: Newable<TModel>): Promise<Array<TModel>> {
    const results: Array<TModel> = [];
    const iterator: ScanIterator<TModel> = await this.mapper.scan<TModel>(value);
    for await (const record of iterator) {
      results.push(record);
    }
    return results;
  }

  async put(id: string, value: TModel, model: Newable<TModel>): Promise<TModel> {
    const toUpdate: TModel = Object.assign(new model(), value, { id });
    return await this.mapper.put<TModel>(toUpdate);
  }

  async post(value: TModel, model: Newable<TModel>): Promise<TModel> {
    const toCreate: TModel = Object.assign(new model(), value);
    return await this.mapper.put<TModel>(toCreate);
  }

  async delete(id: string, model: Newable<TModel>): Promise<TModel> {
    const toDelete: TModel = Object.assign(new model(), { id });
    return await this.mapper.delete<TModel>(toDelete);
  }
}