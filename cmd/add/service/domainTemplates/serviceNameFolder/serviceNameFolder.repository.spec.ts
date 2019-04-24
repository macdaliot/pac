import { {{.serviceNamePascal}} } from './{{.serviceName}}';
import {
  GetOptions,
  ScanOptions,
  ParallelScanWorkerOptions,
  ScanIterator,
  PutOptions,
  DeleteOptions
} from '@aws/dynamodb-data-mapper';
import {
  Repository,
  IEntity,
  Newable
} from '@pyramid-systems/core';

const items = [
  { id: '1', title: 'hello' },
  { id: '2', title: 'world' },
  { id: '3', title: 'universe' }
];

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
      };
    } else {
      return {
        done: true,
        value: null
      };
    }
  }

  [Symbol.iterator](): IterableIterator<TestEntity> {
    return this;
  }
  pointer = 0;

  constructor(public entities: TestEntity[]) {}
}

let entitiesIteratable = new Frame(items);

const DataMapper = jest.fn(() => ({
  get: jest.fn((item: IEntity, options?: GetOptions) => {
    return items.find(x => x.id == item.id);
  }),
  scan: jest.fn(
      (
          value: Newable<IEntity>,
          options?: ScanOptions | ParallelScanWorkerOptions
      ) => {
        return entitiesIteratable;
      }
  ),

  put: jest.fn((value: IEntity, options?: PutOptions) => {
    return value;
  }),
  delete: jest.fn((model: IEntity, options?: DeleteOptions) => {
    const entityDeleted = items.find(x => x.id === model.id);
    return entityDeleted;
  }),
  post: jest.fn((value: IEntity, options?: PutOptions) => {
    return value;
  })
}));

const mockObject = DataMapper() as any;

describe('{{.serviceNamePascal}} Repository', () => {
  it('Should call the repository getById once', async () => {
    const classToTest = new {{.serviceNamePascal}}Repository(mockObject);

    const returnValue = await classToTest.getById('one');

    expect(jest.isMockFunction(mockObject.get)).toBeTruthy();
    expect(mockObject.get).toBeCalledTimes(1);
  });

  it('Should call the repository getAll once', async () => {
    const classToTest = new {{.serviceNamePascal}}Repository(mockObject);

    const returnValue = await classToTest.getAll();

    expect(jest.isMockFunction(mockObject.scan)).toBeTruthy();
    expect(mockObject.scan).toBeCalledTimes(1);
  });

  it('Should call the repository add once on post', async () => {
    const temp = new Giraffe();
    const classToTest = new {{.serviceNamePascal}}Repository(mockObject);

    const returnValue = await classToTest.add(temp);

    expect(jest.isMockFunction(mockObject.put)).toBeTruthy();
    expect(mockObject.put).toBeCalledTimes(1);
  });
  it('Should call the repository put twice on update', async () => {
    const temp = new Giraffe();
    const classToTest = new {{.serviceNamePascal}}Repository(mockObject);

    const returnValue = await classToTest.update(temp.id, temp);

    expect(jest.isMockFunction(mockObject.put)).toBeTruthy();
    expect(mockObject.put).toBeCalledTimes(2);
  });

  it('Should call the repository deleteById once', async () => {
    const classToTest = new {{.serviceNamePascal}}Repository(mockObject);
    const returnValue = await classToTest.deleteById('test');

    expect(jest.isMockFunction(mockObject.delete)).toBeTruthy();
    expect(mockObject.delete).toBeCalledTimes(1);
  });
});
