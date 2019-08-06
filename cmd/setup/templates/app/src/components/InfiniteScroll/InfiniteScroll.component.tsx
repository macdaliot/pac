// /**
//  * To Use this component you must
//  * 1. have a lazy loading api endpoint and a service client defined in the redux/action directory
//  * 2. define a reducer for your entity, this should be registered in the combined root reducer (replace ENTITY_STATE below)
//  * 3. create a custom element/component which will be used to render the data in the renderCell func
//  *
//  */
// import * as React from "react";
// import { hot } from 'react-hot-loader/root';
// import { InfiniteScroller } from "react-iscroller";
// import { ApplicationState } from "@app/redux/Reducers/Reducer";
// import { connect } from "react-redux";
// import { RetrieveData } from "@app/redux/Actions/Entity";
// import "./InfiniteScroll.css";
//
// export const mapStateToProps = (state: ApplicationState) => {
//     return {
//         isCallInProgress: state.ENTITY_STATE.callInProgress,
//         data: state.ENTITY_STATE.queryResult
//     }
// };
//
// export const mapDispatchToProps = (dispatch) => ({
//     fetchData: () => {
//       dispatch(RetrieveData)
//     }
//   });
//
// type ReduxProps = ReturnType<typeof mapStateToProps>;
// type ReduxActions = ReturnType<typeof mapDispatchToProps>;
// type InfiniteScrollerProps = ReduxProps & ReduxActions;
//
// export class Scroller extends React.Component<InfiniteScrollerProps> {
//
//     componentDidMount() {
//         this.props.fetchData();
//     }
//
//     render = () => {
//         if (this.props.data != null) {
//             return (
//                 <InfiniteScroller
//                     itemAverageHeight={100} //play around with these values until scroll looks good
//                     containerHeight={500} //play around with these values until scroll looks good
//                     items={this.props.data}
//                     itemKey="id"
//                     onRenderCell={this.renderCell}
//                     onEnd={this.onEnd}
//                 />
//             );
//         }
//         else {
//             return (
//                 <div>No data</div>
//             )
//         }
//     };
//     renderCell = (item: any, index: number) => {
//         return (
//             <div> <!-- your component goes here--></div>
//         );
//     };
//
//     onEnd = () => {
//         this.props.fetchData();
//     };
//
// }
//
// export default connect(
//     mapStateToProps,
//     mapDispatchToProps
// )(hot(Scroller));
