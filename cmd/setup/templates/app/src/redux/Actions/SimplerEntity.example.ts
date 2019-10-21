// // Replace all Entity with your entity name/service name

// import { createAction, ActionsUnion } from "@app/core/action";
// import { Dispatch } from "redux";
// import { GetState } from "../Reducers/Reducer";
// import axios from "axios";
// import { Entity } from "@pyramid-systems/domain";

// export const GotData = 'Entity_GOT_DATA';
// export const LoadingData = 'Entity_LOADING_DATA';
// export const ErrorOccurred = 'Entity_FAILED';

// export const Actions = {
//     gotData: (data: Entity[]) => createAction(GotData, data),
//     loadingData: () => createAction(LoadingData),
//     failed: (err: Error) => createAction(ErrorOccurred, err)
// };
// export type Actions = ActionsUnion<typeof Actions>;

// export const getData = async (dispatch: Dispatch, getState: GetState) => {
//     try {
//         dispatch(Actions.loadingData());
//         const data = await axios.get<Entity[]>('http://localhost:3000/api/Entity');
//         dispatch(Actions.gotData(data.data));
//     }
//     catch (err) { 
//         dispatch(Actions.failed(err));
//     }
// }