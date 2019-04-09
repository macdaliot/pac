import {
    DataMapper,
    ScanOptions,
    ParallelScanWorkerOptions,
    PutOptions,
    DeleteOptions,
    GetOptions
} from '@aws/dynamodb-data-mapper';
import { config, DynamoDB } from 'aws-sdk';
import { awsConfig } from './config';

export type Newable<T> = { new (...args: any[]): T };

export interface IEntity {
    id: string;
}

export class Repository<TModel extends IEntity> {
    mapper: DataMapper;
    constructor(protected dynamoDb: DynamoDB) {
        // TODO: Create a type for environment variables
        let env = process.env.ENVIRONMENT || 'cloud';
        config.update(awsConfig[env]);
        this.mapper = new DataMapper({
            client: new DynamoDB() // the SDK client used to execute operations
            // tableNamePrefix: 'dev_' // optionally, you can provide a table prefix to keep your dev and prod tables separate
        });
    }

    protected async get(
        id: string,
        model: Newable<TModel>,
        options?: GetOptions
    ): Promise<TModel> {
        const getRequest = new model();
        getRequest.id = id;
        return await this.mapper.get<TModel>(getRequest, options);
    }

    /*
      Retrieves all values in a table or index.
    */
    protected async scan(
        value: Newable<TModel>,
        options?: ScanOptions | ParallelScanWorkerOptions
    ) {
        return await this.mapper.scan<TModel>(value, options);
    }

    protected async put(value: TModel, options?: PutOptions) {
        return await this.mapper.put<TModel>(value, options);
    }

    protected async delete(
        id: string,
        model: Newable<TModel>,
        options?: DeleteOptions
    ) {
        const deleteRequest = new model();
        deleteRequest.id = id;
        return await this.mapper.delete<TModel>(deleteRequest, options);
    }
}
