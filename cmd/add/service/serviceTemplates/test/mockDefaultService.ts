import { PromiseResult } from "aws-sdk/lib/request";

export class MockDefaultService {
    _mockResponseFromDB: PromiseResult<any, any> = {
        "key": "value"
    }
    db;
    get = async (query: any) => this._mockResponseFromDB;
    getById = async (id: string) => this._mockResponseFromDB;
    post = async (body) => this._mockResponseFromDB;
}