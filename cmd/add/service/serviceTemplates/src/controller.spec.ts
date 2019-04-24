import { {{.serviceNamePascal}}Controller } from './{{.serviceName}}.controller';
import { {{.serviceNamePascal}}Repository, {{.serviceNamePascal}} } from '@pyramid-systems/domain';

jest.mock('@pyramid-systems/domain', () => {
  return {
    getById: jest.fn(() => {
      return { id: 'one' };
    }),
    getAll: jest.fn(() => {
      return {};
    }),
    add: jest.fn(test => {
      return test;
    }),
    update: jest.fn(() => {
      return {};
    }),
    deleteById: jest.fn(() => {
      return { id: 'test' };
    })
  };
});

const mockObject = require('@pyramid-systems/domain');

describe('{{.serviceNamePascal}} Controller', () => {
  it('Should call the repository getById once', async () => {
    const classToTest = new {{.serviceNamePascal}}Controller(mockObject);

    const returnValue = await classToTest.getById('one');

    expect(jest.isMockFunction(mockObject.getById)).toBeTruthy();
    expect(mockObject.getById).toBeCalledTimes(1);
    expect(returnValue.id).toEqual('one');
  });

  it('Should call the repository getAll once', async () => {
    const classToTest = new {{.serviceNamePascal}}Controller(mockObject);

    const returnValue = await classToTest.getAll();

    expect(jest.isMockFunction(mockObject.getAll)).toBeTruthy();
    expect(mockObject.getAll).toBeCalledTimes(1);
  });

  it('Should call the repository post once', async () => {
    const temp = new {{.serviceNamePascal}}();
    const classToTest = new {{.serviceNamePascal}}Controller(mockObject);

    const returnValue = await classToTest.post(temp);

    expect(jest.isMockFunction(mockObject.add)).toBeTruthy();
    expect(mockObject.add).toBeCalledTimes(1);
  });

  it('Should call the repository put once', async () => {
    const temp = new Giraffe();
    const classToTest = new {{.serviceNamePascal}}Controller(mockObject);

    const returnValue = await classToTest.put(temp.id, temp);

    expect(jest.isMockFunction(mockObject.update)).toBeTruthy();
    expect(mockObject.update).toBeCalledTimes(1);
  });

  it('Should call the repository deleteById once', async () => {
    const classToTest = new {{.serviceNamePascal}}Controller(mockObject);
    const returnValue = await classToTest.deleteById('test');

    expect(jest.isMockFunction(mockObject.deleteById)).toBeTruthy();
    expect(mockObject.deleteById).toBeCalledTimes(1);
  });
});
