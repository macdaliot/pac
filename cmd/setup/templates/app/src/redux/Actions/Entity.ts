// Copy this file to a name that matches your entity, and uncomment. Replace
// every instance of "Entity" with the name of yours and correct the imports as necessary.

// import { createAction, ActionsUnion } from "@app/core/action";
// import { Entity } from "@pyramid-systems/domain";
// import { Dispatch } from "redux";
// import { EntityService } from "@app/services/entity/lib/entityService";
// import ApplicationTokenCredentials from "@app/core/application.token.credentials";
// import { UrlConfig } from "@app/config";
// import { GetState } from "../Reducers/Reducer";

// export const CALL_BEGUN = 'Entity_CALL_BEGUN'
// export const CALL_RESPONSE = 'Entity_CALL_RESPONSE'
// export const CALL_FAILED = 'Entity_CALL_FAILED'

// export const Actions = {
//     callBegun: () => createAction(CALL_BEGUN),
//     callFailed: () => createAction(CALL_FAILED),
//     response: (payload: Array<Entity>) => createAction(CALL_RESPONSE, payload),
// };
// export type ActionsType = ActionsUnion<typeof Actions>;

// export const Behavior = {
//     // set this true to bypass the actual api and return mock data instead
//     //
//     // Useful if the api's have not been deployed yet or are unreachable
//     mock: true,

//     // Set this true to ignore auth information and make calls with a dummy token
//     // 
//     // Useful if the authentication is not working or not set up
//     ignoreAuth: false
// }
// const mockData = {
//     getAll: [
//         { id: '1', buttons: 2, axles: 3 },
//         { id: '2', buttons: 2, axles: 6 },
//         { id: '3', buttons: 4, axles: 9 },
//         { id: '4', buttons: 6, axles: 12 },
//     ]
// };

// export const getAll = (dispatch: Dispatch, getState: GetState, serviceOverride?: EntityService) => {
//     callService(dispatch, getState, (service) => {
//         service.getAll().then(res => {
//             dispatch(Actions.response(res))
//         })
//     },
//     () => (mockData.getAll), serviceOverride)
// }
// export const getById = (id: string) => (dispatch: Dispatch, getState: GetState, serviceOverride?: EntityService) => {
//     callService(dispatch, getState, (service) => {
//         service.getById(id).then(res => {
//             dispatch(Actions.response([res]))
//         })
//     },
//     () => ([mockData.getAll.find(each => each.id == id)]), serviceOverride)
// }
// export const createService = (token: string) => new EntityService(ApplicationTokenCredentials(token), { baseUri: UrlConfig.apiUrl });
// export const callService = (dispatch: Dispatch, getState: GetState, body: (service: EntityService) => void, getMockData: () => Entity[], service?: EntityService) => {
//     dispatch(Actions.callBegun());
//     const token = Behavior.ignoreAuth ? "token" : getState().authentication.token;
//     if (token) {
//         if (Behavior.mock) {
//             return dispatch(Actions.response(getMockData()));
//         }
//         if (!service)
//             service = createService(token);
//         try {
//             body(service)
//             console.log("Call finished")
//         }
//         catch(err){
//             console.log(err);
//             dispatch(Actions.callFailed())
//         }
//         console.log("foo");
//     }
//     else {
//         console.log("No token");
//         dispatch(Actions.callFailed());
//     }
// }