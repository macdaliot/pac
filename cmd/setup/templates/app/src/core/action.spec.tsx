import {isActionOf, createAction, Action} from "@app/core/action";

describe('isActionOf  createAction', () => {
    it("should return true for matching action", () => {
        const actionType : string = "FOO_ACTION"
        const myAction  = createAction(actionType);
        const isAction : boolean = isActionOf(actionType, myAction);
        expect(isAction).toBe(true);
    })
    it("should return False for non-matching action", () => {
        const actionType : string = "FOO_ACTION"
        const myAction  = createAction(actionType);
        const isAction : boolean = isActionOf("BAR_ACTION", myAction);
        expect(isAction).toBe(false);
    })
    it("should return action with type field only", () => {
        const actionType : string = "FOO_ACTION"
        const myAction  = createAction(actionType);
        expect(myAction).toHaveProperty("type");
        expect(myAction.type).toEqual(actionType)
    })
    it("should return action with type field and payload field", () => {
        const actionType : string = "FOO_ACTION"
        const payload : any = {data:"test payload"}
        const myAction  = createAction(actionType, payload);
        expect(myAction).toHaveProperty("type");
        expect(myAction).toHaveProperty("payload");
        expect(myAction.type).toEqual(actionType)
        // @ts-ignore
        expect(myAction.payload).toEqual(payload)
    })
});