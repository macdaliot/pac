// Copy this file to a name that matches your entity, and uncomment. Replace
// every instane of "Entity" with the name of yours and correct the imports as necessary.

// import { Entity } from "@pyramid-systems/domain";
// import { ActionsType, CALL_BEGUN, CALL_FAILED, CALL_RESPONSE } from "../Actions/Entity";

// type EntityState = {
//     callInProgress?: boolean,
//     queryResult?: Array<Entity>
// }

// const initialApplicationState: EntityState = {};

// export function reducer(
//     state = initialApplicationState,
//     action: ActionsType
// ): EntityState {
//     switch (action.type) {
//         case CALL_BEGUN:
//             return { callInProgress: true };
//         case CALL_FAILED:
//             return initialApplicationState;
//         case CALL_RESPONSE:
//             return { queryResult: action.payload };
//         default:
//             return { ...state };
//     }
// };