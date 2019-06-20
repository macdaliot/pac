import { WebStorage } from "./index";
jest.mock('./index');

let StorageMock: jest.Mocked<IWebStorage> = WebStorage as any;
describe('Unit Tests for webStorage', () => {
    it("Testing the getItem function of webStorage", () => {
        StorageMock.getItem.mockReturnValue("that");
        WebStorage.setItem("this", "that");
        const item = WebStorage.getItem("this");
        expect(item).toEqual("that");
    })
});
