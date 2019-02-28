import { generateRandomString } from '../src/functions'

describe("generic functions", () => {
    it("should generate random strings", () => {
        let first = generateRandomString();
        let second = generateRandomString();
        expect(second == first).toBeFalsy();
    })
})