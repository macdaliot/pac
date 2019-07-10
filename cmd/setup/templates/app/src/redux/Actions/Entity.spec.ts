// Copy this file to a name that matches your entity, and uncomment. Replace
// every instance of "Entity" with the name of yours and correct the imports as necessary.


// import { callService, Behavior, Actions, getAll, getById } from "./Entity";

// describe('entity actions', () => {
//     afterEach(()=> {
//         Behavior.ignoreAuth = false;
//         Behavior.mock = false;
//     })
//     describe('callService', () => {

//         it('should not call body without auth', () => {
//             const body = jest.fn(() => Promise.resolve(null));
//             const dispatch = jest.fn();
//             const getState = jest.fn(() => ({authentication: {}, entity: {}}));
//             const getMockData = jest.fn();
//             callService(dispatch, getState, body, getMockData);
//             expect(body).not.toHaveBeenCalled();
//         })
//         it('should call body without auth if skipped', () => {
//             const body = jest.fn(() => Promise.resolve(null));
//             const dispatch = jest.fn();
//             const getState = jest.fn(() => ({authentication: {}, entity: {}}));
//             const getMockData = jest.fn();
//             Behavior.ignoreAuth = true;
//             callService(dispatch, getState, body, getMockData);
//             expect(body).toHaveBeenCalled();
//         })
//         it('should call getMock if mocked', () => {
//             const body = jest.fn(() => Promise.resolve(null));
//             const dispatch = jest.fn();
//             const getState = jest.fn(() => ({authentication: {}, entity: {}}));
//             const getMockData = jest.fn();
//             Behavior.ignoreAuth = true;
//             Behavior.mock = true
//             callService(dispatch, getState, body, getMockData);
//             expect(body).not.toHaveBeenCalled();
//             expect(getMockData).toHaveBeenCalled();
//         })
//         it('should throw if error', () => {            
//             const body = jest.fn(() => {throw new Error()});
//             const dispatch = jest.fn((stuff) => {expect(stuff).toBeTruthy(); console.log(JSON.stringify(stuff))}) as any;
//             const getState = jest.fn(() => ({authentication: {}, entity: {}}));
//             const getMockData = jest.fn();
//             Behavior.ignoreAuth = true;
//             callService(dispatch, getState, body, getMockData);
//             expect(body).toHaveBeenCalled();
//             expect(dispatch).toHaveBeenNthCalledWith(1, Actions.callBegun())
//             expect(dispatch).toHaveBeenNthCalledWith(2, Actions.callFailed())
//         })
//     }) 
//     describe('thunk methods', () => {
//         it('getAll should call getAll', () => {
//             const dispatch = jest.fn();
//             const getState = jest.fn(() => ({ authentication: {}, entity: {}}));
//             const service = { getAll: jest.fn(() => { return Promise.resolve(null) }) } as any;
//             Behavior.ignoreAuth = true;
//             getAll(dispatch, getState, service);
//             expect(service.getAll).toHaveBeenCalled();
//         })
//         it('getById should call getById', () => {
//             const dispatch = jest.fn();
//             const getState = jest.fn(() => ({ authentication: {}, entity: {}}));
//             const service = { getById: jest.fn(() => { return Promise.resolve(null) }) } as any;
//             Behavior.ignoreAuth = true;
//             getById("1")(dispatch, getState, service);
//             expect(service.getById).toHaveBeenCalledWith("1");
//         })
//         it('getAll should call getMock if mock is true', () => {
//             const dispatch = jest.fn();
//             const getState = jest.fn(() => ({ authentication: {}, entity: {}}));
//             const service = { getAll: jest.fn(() => { return Promise.resolve(null) }) } as any;
//             Behavior.ignoreAuth = true;
//             Behavior.mock = true;
//             getAll(dispatch, getState, service);
//             expect(service.getAll).not.toHaveBeenCalled();
//         })
//         it('getById should call getMock if mock is true', () => {
//             const dispatch = jest.fn();
//             const getState = jest.fn(() => ({ authentication: {}, entity: {}}));
//             const service = { getById: jest.fn(() => { return Promise.resolve(null) }) } as any;
//             Behavior.ignoreAuth = true;
//             Behavior.mock = true;
//             getById("1")(dispatch, getState, service);
//             expect(service.getById).not.toHaveBeenCalled();
//         })
//     })
// })