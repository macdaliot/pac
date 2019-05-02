module.exports = {
  preset: 'ts-jest',
  testEnvironment: 'node',
  moduleNameMapper: {
    "^@pyramid-systems/core(.*)$": "<rootDir>/../../core$1",
    '^@pyramid-systems/domain(.*)$': '<rootDir>/../../domain$1'
  },
  modulePathIgnorePatterns: ['docs', 'dist'],
  coveragePathIgnorePatterns: ['dist']
};