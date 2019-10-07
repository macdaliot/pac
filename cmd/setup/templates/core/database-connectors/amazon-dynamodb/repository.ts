import {
    DataMapper,
    ScanOptions,
    ParallelScanWorkerOptions,
    PutOptions,
    DeleteOptions,
    GetOptions
} from '@aws/dynamodb-data-mapper';
import { Injectable } from '@pyramid-systems/core';

export type Newable<T> = { new(...args: any[]): T };

export interface IEntity {
    id: string;
}

type IMapper = DataMapper;

@Injectable()
export class Repository<TModel extends IEntity> {
    constructor(protected mapper: IMapper) {
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

    protected async put(
        id: string,
        value: TModel,
        model: Newable<TModel>,
        options?: PutOptions
    ) {
        const toUpdate = Object.assign(new model(), value, { id });
        return await this.mapper.put<TModel>(toUpdate, options);
    }

    protected async post(
        value: TModel,
        model: Newable<TModel>,
        options?: PutOptions
    ) {
        const toUpdate = Object.assign(new model(), value);
        return await this.mapper.put<TModel>(toUpdate, options);
    }

    protected async delete(
        id: string,
        model: Newable<TModel>,
        options?: DeleteOptions
    ) {
        const toDelete = Object.assign(new model(), { id });
        return await this.mapper.delete<TModel>(toDelete, options);
    }
}
