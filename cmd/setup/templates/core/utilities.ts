export const generateRandomString = () => {
    return Math.random()
        .toString(36)
        .substring(2, 15);
};

export const isNullOrUndefined = (value: any): boolean => {
    return value === undefined || value === null;
};

export const intersection = (arr1: any[], arr2: any[]) =>
    arr1.filter(el => arr2.includes(el));
