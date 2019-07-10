// Copy this file to a name that matches your entity, and uncomment. Replace
// every instance of "Entity" with the name of yours and correct the imports as necessary.


// import { Actions } from '../Actions/Entity'
// import { reducer } from './Entity'

// describe('Entity reducer', () => {
//     it('should handle call begun', () => {
//         const result = reducer({}, Actions.callBegun());
//         expect(result.callInProgress).toBeTruthy();
//     })
//     it('should handle call failed', () => {
//         const result = reducer({callInProgress: true}, Actions.callFailed());
//         expect(result.callInProgress).toBeFalsy();
//     })
//     it('should handle call results', () => {
//         const data = [{id: '4', buttons: 0, axles: 0}];
//         const result = reducer({callInProgress: true}, Actions.response(data));
//         expect(result.queryResult).toHaveLength(1);
//         expect(result.queryResult).toEqual(data);
//     })
// });