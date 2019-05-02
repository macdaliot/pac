import { isNullOrUndefined, intersection, generateRandomString } from '@pyramid-systems/core';

describe('generateRandomString', () => {
    it('should return a random string', () => {
        expect(generateRandomString()).toBeTruthy();
    })
    it('should have always have a length of 11', () => {
        expect(generateRandomString().length).toBeGreaterThan(8);;
    })
});

describe('isNullOrUndefined', () => {
    it('should return false if an object is not null or undefined', () => {
        const testValue = { payload: 5 };
        expect(isNullOrUndefined(testValue)).toBe(false);
        expect(isNullOrUndefined('testValue')).toBe(false);
        expect(isNullOrUndefined(5)).toBe(false);
        expect(isNullOrUndefined(true)).toBe(false);
    })

    it('should return true if an object is null or undefined', () => {
        expect(isNullOrUndefined(null)).toBe(true);
        expect(isNullOrUndefined(undefined)).toBe(true);
    })
})

describe('intersection', () => {

    it('should return the intersection between two number array that has intersection', () => {
        const firstArray = [1, 2, 3, 4, 5];
        const secondArray = [3, 4, 5];
        const output = intersection(firstArray, secondArray);

        expect(output).toEqual([3, 4, 5]);
    })

    it('should return the intersection between two string that has intersection', () => {
        const firstArray = ["1", "2", "3", "4", "5"];
        const secondArray = ["3", "4", "5"];
        const output = intersection(firstArray, secondArray);

        expect(output).toEqual(["3", "4", "5"]);
    })

    it('should return the intersection between two boolean array that has intersection', () => {
        const firstArray = [true, true, false];
        const secondArray = [false, true];
        const output = intersection(firstArray, secondArray);

        expect(output).toEqual([true, true, false]);
    })

    it('should return empty between two number array that has no intersection', () => {
        const firstArray = [1, 2, 3, 4, 5];
        const secondArray = [6, 7, 8];
        const output = intersection(firstArray, secondArray);

        expect(output).toEqual([]);
    })

    it('should return empty between two string array that has no intersection', () => {
        const firstArray = ["1", "2", "3", "4", "5"];
        const secondArray = ["6", "7", "8"];
        const output = intersection(firstArray, secondArray);

        expect(output).toEqual([]);
    })

    it('should return the intersection between two boolean array that has no intersection', () => {
        const firstArray = [true];
        const secondArray = [false];
        const output = intersection(firstArray, secondArray);

        expect(output).toEqual([]);
    })
})