// import * as React from 'react';
// import { shallow } from 'enzyme';
// import {Scroller, mapStateToProps, mapDispatchToProps} from './InfiniteScroll.component';
// import { IUser } from '@pyramid-systems/domain';
//
// describe('header (unit/shallow)', () => {
//     // if you've exported the class by itself, you can do quick unit tests like the following test
//     // shallow() will not render child components
//
//     it('should show a no data message', () => {
//         const hdr = shallow(<Scroller
//             isCallInProgress={false}
//             data={null}
//             getInitialList={() => ({})}
//             getNextEntries={() => ({})}
//         />);
//         expect(
//             hdr.contains(
//                 <div>No Data Response</div>
//             )
//         ).toBe(true);
//     });
//
//     it('should show data container', () => {
//         const props = {
//             isCallInProgress: false,
//             data: {symbol: "testSymbol"}
//         };
//         // @ts-ignore
//         const hdr = shallow(<Scroller
//             {...props}
//             getInitialList={() => ({})}
//             getNextEntries={() => ({})}/>);
//         expect(hdr.html().includes(`<div class="" style="overflow:auto;-webkit-overflow-scrolling:touch;overflow-anchor:none;height:600px"><div style="height:0"></div><div style="height:0"></div></div>`)).toBe(true);
//     });
//
//     it('should map state appropriately', () => {
//         const fakeUser: IUser = {
//             name: 'testUser',
//             groups: [],
//             sub: ''
//         };
//         const inputState = {
//             authentication: {
//                 user: fakeUser
//             },
//             analyzer: {
//                 callInProgress:true,
//                 queryResult:{test:"not"}
//             }
//         };
//         const expectedProps = {
//             isCallInProgress: true,
//             data: {test:"not"}
//         };
//         // @ts-ignore
//         expect(mapStateToProps(inputState)).toEqual(expectedProps);
//     });
//
//     it('should dispatch correct actions', () => {
//         const dispatch = jest.fn();
//
//         mapDispatchToProps(dispatch).getNextEntries();
//         mapDispatchToProps(dispatch).getInitialList();
//         expect(dispatch.mock.calls[0][0].name).toEqual('getNext');
//         expect(dispatch.mock.calls[1][0].name).toEqual('getInitialList');
//     });
//
//     it('should render a cell with specified ID', () => {
//         // @ts-ignore
//         let renderedCell = new Scroller().renderCell({id:"myId"});
//         expect(renderedCell.key).toEqual("myId");
//     });
//
//     it('should render a cell with specified ID', () => {
//         let nxtEntryMock = jest.fn();
//         // @ts-ignore
//         new Scroller({getNextEntries:nxtEntryMock}).onEnd();
//         expect(nxtEntryMock).toBeCalledTimes(1);
//     });
//
// });
//
