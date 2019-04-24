import { GetOptions, ScanOptions, ParallelScanWorkerOptions, ScanIterator, PutOptions, DeleteOptions } from '@aws/dynamodb-data-mapper';
import { IEntity, Newable } from 'database-connectors';

const items = [
    { id: "1", title: "hello" },
    { id: "2", title: "world" },
    { id: "3", title: "universe" }
]

interface Iterator<T> {
    next(value?: any): IteratorResult<T>;
    return?(value?: any): IteratorResult<T>;
    throw?(e?: any): IteratorResult<T>;
}

class TestEntity implements IEntity {
    id: string;

}

class Frame implements IterableIterator<TestEntity> {
    next(): IteratorResult<TestEntity> {
        if (this.pointer < this.entities.length) {
            return {
                done: false,
                value: this.entities[this.pointer++]
            }
        } else {
            return {
                done: true,
                value: null
            }
        }
    }

    [Symbol.iterator](): IterableIterator<TestEntity> {
        return this;
    }
    pointer = 0;

    constructor(public entities: TestEntity[]) {

    }
}


let entitiesIteratable = new Frame(items);

const DataMapper = jest.fn(() => ({
    get: jest.fn((item: IEntity, options?: GetOptions) => {
        return items.find(x => x.id == item.id);
    }),
    scan: jest.fn((value: Newable<IEntity>, options?: ScanOptions | ParallelScanWorkerOptions) => {
        return entitiesIteratable;
    }),

    put: jest.fn((value: IEntity, options?: PutOptions) => {
        return value;
    }),
    delete: jest.fn((
        model: IEntity,
        options?: DeleteOptions
    ) => {
        const entityDeleted = items.find(x => x.id === model.id);
        return entityDeleted;
    }),
    post: jest.fn((value: IEntity, options?: PutOptions) => {
        return value;
    })
}));

export { DataMapper };