import {IWebStorage, WebStorageImpl, IStorageBackend} from "./WebStorage";
jest.mock('./index');

describe('webStorage local storage wrapper', () => {
    let cut: IWebStorage;
    let storageMock: IStorageBackend;
    beforeEach(() => {
        storageMock = { getItem: jest.fn(), setItem: jest.fn(), removeItem: jest.fn() };
        cut = new WebStorageImpl(storageMock);
    })
    it("should wrap the getItem function", () => {
        cut.getItem("foo");
        expect(storageMock.getItem).toBeCalledTimes(1);
    })
    it("should use the getItem function in hasItem", () => {
        cut.hasItem("foo");
        expect(storageMock.getItem).toBeCalledTimes(1);
    })
    it("should use the setItem function in setItem", () => {
        cut.setItem("foo", "bar");
        expect(storageMock.setItem).toBeCalledTimes(1);
    })
    it("should use the removeItem function in removeItem", () => {
        cut.removeItem("foo");
        expect(storageMock.removeItem).toBeCalledTimes(1);
    })
    it("should notice that storage isn't supported", () => {
        cut = new WebStorageImpl({} as IStorageBackend)
        expect(cut.isSupported()).toBeFalsy();
    })
});