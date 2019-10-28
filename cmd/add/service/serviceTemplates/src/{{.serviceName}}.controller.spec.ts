import { {{.serviceNamePascal}}Controller } from './{{.serviceName}}.controller';
import { {{.serviceNamePascal}} } from '@pyramid-systems/domain';

describe('{{.serviceNamePascal}} Controller', () => {
  xit('Should not fail', () => {});
});

//
// const mockRepository = {
//   getById: jest.fn(() => {
//     return { id: 'one' };
//   }),
//   getAll: jest.fn(() => {
//     return {};
//   }),
//   add: jest.fn(test => {
//     return test;
//   }),
//   update: jest.fn(() => {
//     return {};
//   }),
//   deleteById: jest.fn(() => {
//     return { id: 'test' };
//   })
// } as any;
//
//   it('Should call the repository getById once', async () => {
//     const classToTest = new {{.serviceNamePascal}}Controller(mockRepository);
//
//     const returnValue = await classToTest.getById('one');
//
//     expect(jest.isMockFunction(mockRepository.getById)).toBeTruthy();
//     expect(mockRepository.getById).toBeCalledTimes(1);
//     expect(returnValue.id).toEqual('one');
//   });
//
//   it('Should call the repository getAll once', async () => {
//     const classToTest = new {{.serviceNamePascal}}Controller(mockRepository);
//
//     const returnValue = await classToTest.getAll();
//
//     expect(jest.isMockFunction(mockRepository.getAll)).toBeTruthy();
//     expect(mockRepository.getAll).toBeCalledTimes(1);
//   });
//
//   it('Should call the repository post once', async () => {
//     const temp = new {{.serviceNamePascal}}();
//     const classToTest = new {{.serviceNamePascal}}Controller(mockRepository);
//
//     const returnValue = await classToTest.post(temp);
//
//     expect(jest.isMockFunction(mockRepository.add)).toBeTruthy();
//     expect(mockRepository.add).toBeCalledTimes(1);
//   });
//
//   it('Should call the repository put once', async () => {
//     const temp = new {{.serviceNamePascal}}();
//     const classToTest = new {{.serviceNamePascal}}Controller(mockRepository);
//
//     const returnValue = await classToTest.put(temp.id, temp);
//
//     expect(jest.isMockFunction(mockRepository.update)).toBeTruthy();
//     expect(mockRepository.update).toBeCalledTimes(1);
//   });
//
//   it('Should call the repository deleteById once', async () => {
//     const classToTest = new {{.serviceNamePascal}}Controller(mockRepository);
//     const returnValue = await classToTest.deleteById('test');
//
//     expect(jest.isMockFunction(mockRepository.deleteById)).toBeTruthy();
//     expect(mockRepository.deleteById).toBeCalledTimes(1);
//   });
// });
