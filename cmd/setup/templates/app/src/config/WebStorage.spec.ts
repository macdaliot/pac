import {IWebStorage, WebStorageImpl, IStorageBackend} from "./WebStorage";
jest.mock('./index');

describe('webStorage local storage wrapper', () => {
    let cut: IWebStorage;
    let storageMock: IStorageBackend;
    beforeEach(() => {
        storageMock = { getItem: jest.fn(), setItem: jest.fn() };
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
    it("should use notice that storage isn't supported", () => {
        cut = new WebStorageImpl({} as IStorageBackend)
        expect(cut.isSupported()).toBeFalsy();
    })
});
