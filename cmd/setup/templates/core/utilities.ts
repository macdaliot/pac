export const generateRandomString = () => {
    return Math.random()
        .toString(36)
        .substring(2, 15);
};

export const isNullOrUndefined = (value: any): boolean => {
    return value === undefined || value === null;
};

export const intersection = (firstArray: (string | number | boolean)[], secondArray: (string | number | boolean)[]) => {
    return firstArray.filter(Set.prototype.has, new Set(secondArray));
}
