module.exports = {
  preset: 'ts-jest',
  testEnvironment: 'node',
  moduleNameMapper: {
    "^@pyramid-systems/core(.*)$": "<rootDir>/dist/core$1"
  },
  modulePathIgnorePatterns: ["docs"]
};