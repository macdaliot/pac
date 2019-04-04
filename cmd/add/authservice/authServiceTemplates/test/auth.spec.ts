import * as express from 'express';
import { errorMiddleware, generateRandomString, HttpException } from '@pyramidlabs/core'

describe("generic functions", () => {
  it("should generate random strings", () => {
    let first = generateRandomString();
    let second = generateRandomString();
    expect(second == first).toBeFalsy();
  });
  it("should pass back a 500 code on any error", () => {
    const expectedStatus = 500;
    let statusResult: number = null;
    let statusFunc: any = (output: number) => {
      statusResult = output;
      return response;
    };
    let sendFunc: any = (output: Error) => { return output; }
    let error = new HttpException(500, "Things went wrong");
    let request = { body: {} } as any as express.Request;
    let response = { status: statusFunc, send: sendFunc } as any as express.Response;
    errorMiddleware(error, request, response, () => {});
    expect(statusResult).toBe(expectedStatus);
  });
});
