// Copy this file to a name that matches your entity, and uncomment. Replace
// every instane of "Entity" with the name of yours and correct the imports as necessary.

// import { createAction, ActionsUnion } from "@app/core/action";
// import { Entity } from "@pyramid-systems/domain";
// import { Dispatch } from "redux";
// import { EntityService } from "@app/services/entity/lib/entityService";
// import ApplicationTokenCredentials from "@app/core/application.token.credentials";
// import { UrlConfig } from "@app/config";
// import { GetState } from "../Reducers/Reducer";

// export const CALL_BEGUN = 'CALL_BEGUN'
// export const CALL_RESPONSE = 'CALL_RESPONSE'
// export const CALL_FAILED = 'CALL_FAILED'

// const Actions = {
//     callBegun: () => createAction(CALL_BEGUN),
//     callFailed: () => createAction(CALL_FAILED),
//     response: (payload: Array<Entity>) => createAction(CALL_RESPONSE, payload),
// };
// export type ActionsType = ActionsUnion<typeof Actions>;

// // set this true to bypass the actual api and return mock data instead
// //
// // Useful if the api's have not been deployed yet or are unreachable
// let mock = true;

// // Set this true to ignore auth information and make calls with a dummy token
// // 
// // Useful if the authentication is not working or not set up
// let ignoreAuth = false;

// const mockData = {
//     getAll: [
//         { id: '1', buttons: 2, axles: 3 },
//         { id: '2', buttons: 2, axles: 6 },
//         { id: '3', buttons: 4, axles: 9 },
//         { id: '4', buttons: 6, axles: 12 },
//     ]
// };

// export const getAll = (dispatch: Dispatch, getState: GetState) => {
//     callService(dispatch, getState, async (service) => {
//         service.getAll().then(res => {
//             dispatch(Actions.response(res))
//         })
//     },
//     () => (mockData.getAll))
// }
// export const getById = (id: string) => (dispatch: Dispatch, getState: GetState) => {
//     callService(dispatch, getState, async (service) => {
//         service.getById(id).then(res => {
//             dispatch(Actions.response([res]))
//         })
//     },
//     () => ([mockData.getAll.find(each => each.id == id)]))
// }

// export const callService = (dispatch: Dispatch, getState: GetState, body: (service:EntityService) => Promise<void>, getMockData: () => Entity[]) => {
//     dispatch(Actions.callBegun());
//     const token = ignoreAuth ? "token" : getState().authentication.token;
//     if (token) {
//         if (mock) {
//             return dispatch(Actions.response(getMockData()));
//         }
//         const service = new EntityService(ApplicationTokenCredentials(token), { baseUri: UrlConfig.apiUrl });

//         body(service).then(() => {
//             console.log("Call finished")
//         }).catch(err => {
//             console.log(err);
//             dispatch(Actions.callFailed())
//         });
//     }
//     else {
//         console.log("No token");
//         dispatch(Actions.callFailed());
//     }
// }