import { Repository, IEntity, Newable } from '@pyramid-systems/core';
import {
  GetOptions,
  ScanOptions,
  ParallelScanWorkerOptions,
  ScanIterator,
  PutOptions,
  DeleteOptions,
  DataMapper
} from '@aws/dynamodb-data-mapper';

const mockDataMapper = <jest.Mock<DataMapper>>DataMapper;

class TestEntity implements IEntity {
  id: string;
}

class TestRepository extends Repository<TestEntity> {
  get(
    id: string,
    model: Newable<IEntity>,
    options?: GetOptions
  ): Promise<IEntity> {
    return super.get(id, model, options);
  }

  scan(
    value: Newable<IEntity>,
    options?: ScanOptions | ParallelScanWorkerOptions
  ): Promise<ScanIterator<IEntity>> {
    return super.scan(value, options);
  }

  put(
    id: string,
    value: IEntity,
    model: Newable<IEntity>,
    options?: PutOptions
  ) {
    return super.put(id, value, model, options);
  }

  delete(id: string, model: Newable<IEntity>, options?: DeleteOptions) {
    return super.delete(id, model, options);
  }

  post(value: IEntity, model: Newable<IEntity>, options?: PutOptions) {
    return super.post(value, model, options);
  }
}

describe('Repository', () => {
  beforeEach(() => {
    mockDataMapper.mockClear();
  });
  it('should create a new instance without any error', () => {
    const mapper = mockDataMapper();
    const instance = new TestRepository(mapper);

    expect(mockDataMapper).toBeCalledTimes(1);
    expect(instance).toBeTruthy();
  });

  it('post should return the entity', async () => {
    const mapper = mockDataMapper();
    const instance = new TestRepository(mapper);
    const entityToAdd = new TestEntity();
    entityToAdd.id = '1';
    const output = await instance.post(entityToAdd, TestEntity);

    expect(output).toBeTruthy();
    expect(jest.isMockFunction(mapper.put)).toBe(true);
    expect(mapper.put).toBeCalledTimes(1);
    expect(output).toEqual(entityToAdd);
  });

  it('get should return the entity given the id is found', async () => {
    const mapper = mockDataMapper();
    const instance = new TestRepository(mapper);
    const output = await instance.get('1', TestEntity);

    expect(output).toBeTruthy();
    expect(jest.isMockFunction(mapper.get)).toBe(true);
    expect(mapper.get).toBeCalledTimes(1);
  });

  it('get should return the entity given the id not found', async () => {
    const mapper = mockDataMapper();
    const instance = new TestRepository(mapper);
    const output = await instance.get('100000', TestEntity);

    expect(output).toBeUndefined();
    expect(jest.isMockFunction(mapper.get)).toBe(true);
    expect(mapper.get).toBeCalledTimes(1);
  });

  it('scan should return all entities', async () => {
    const mapper = mockDataMapper();
    const instance = new TestRepository(mapper);
    const output = await instance.scan(TestEntity);

    expect(output).toBeTruthy();
    expect(jest.isMockFunction(mapper.scan)).toBe(true);
    expect(mapper.scan).toBeCalledTimes(1);
  });

  it('put should add new entity', async () => {
    const mapper = mockDataMapper();
    const instance = new TestRepository(mapper);
    const entityToAddUpdate = new TestEntity();
    entityToAddUpdate.id = '100';
    const output = await instance.put(
      entityToAddUpdate.id,
      entityToAddUpdate,
      TestEntity
    );

    expect(output).toEqual(entityToAddUpdate);
    expect(jest.isMockFunction(mapper.put)).toBe(true);
    expect(mapper.put).toBeCalledTimes(1);
  });

  it('delete should return the deleted entity', async () => {
    const mapper = mockDataMapper();
    const instance = new TestRepository(mapper);
    const entityToDelete = new TestEntity();
    entityToDelete.id = '1';
    const output = await instance.delete(entityToDelete.id, TestEntity);

    expect(output).toBeTruthy();
    expect(jest.isMockFunction(mapper.delete)).toBe(true);
    expect(mapper.delete).toBeCalledTimes(1);
    expect(output.id).toEqual(entityToDelete.id);
  });
});
