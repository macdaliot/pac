import { errorMiddleware } from './error.middleware';


const mockRequest = jest.fn(() => ({
    log: {
        error: jest.fn()
    }
}));


const message = "Hello";
const status = 500;

const mockHttpException = jest.fn(() => ({
    status: status,
    message: message,
    name: "Test"
}));

let mockResponseStatus = jest.fn();
let mockResponseSend = jest.fn();


let mockResponse = {
    status: mockResponseStatus,
    send: mockResponseSend
};
const mockNextFunction = jest.fn();


describe('HTTP Exception', () => {
    beforeEach(() => {
        mockResponseStatus = jest.fn();
        mockResponseSend = jest.fn();
        mockResponse = {
            status: mockResponseStatus,
            send: mockResponseSend
        };
        mockResponseStatus.mockImplementation(() => mockResponse);
        mockResponseSend.mockImplementation(() => mockResponse);
    })

    it('should send the response with the correct parameters given an http exception is given', () => {
        const httpException = mockHttpException();
        const request = mockRequest() as any;
        const response = mockResponse as any;
        const nextFunction = mockNextFunction();

        errorMiddleware(httpException, request, response, nextFunction);

        expect(request.log.error).toBeCalledTimes(1)
        expect(response.status).toBeCalledTimes(1)
        expect(response.status).toBeCalledWith(status)
        expect(response.status().send).toBeCalledWith({ message, status })
    })


    it('should send the response with the correct parameters given an http exception has no status or message', () => {
        const httpException = { status: null, message: null, name: null };
        const request = mockRequest() as any;
        const response = mockResponse as any;
        const nextFunction = mockNextFunction();

        errorMiddleware(httpException, request, response, nextFunction);

        expect(request.log.error).toBeCalledTimes(1)
        expect(response.status).toBeCalledTimes(1)
        expect(response.status).toBeCalledWith(status)
        expect(response.status().send).toBeCalledWith({ message: 'Something went wrong', status: 500 })
    })
})