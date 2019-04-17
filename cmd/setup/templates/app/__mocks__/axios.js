export default {
  defaults: {
    headers: {
      common: {}
    }
  },
  get: jest.fn(() => Promise.resolve({ data: {} })),
  post: jest.fn((data) => Promise.resolve({}))
}
