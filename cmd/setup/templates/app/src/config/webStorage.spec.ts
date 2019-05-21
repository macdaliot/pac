import { webStorage } from "./webStorage";

describe('Unit Tests for webStorage', () => {
    it("Testing the getData function of webStorage", () => {
        webStorage.setItem("this", "that");
        webStorage.setItem("test", "me");
        const data = webStorage.getData();
        expect(data).toEqual({"this": "that", "test": "me"});
    }),
    it("Testing the getItem function of webStorage", () => {
        webStorage.setItem("this", "that");
        webStorage.setItem("test", "me");
        const item = webStorage.getItem("this");
        expect(item).toEqual("that");
    })
});
