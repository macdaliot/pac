const pino = jest.fn(() => {
    return {
        child: jest.fn()
    }
});

export = pino;