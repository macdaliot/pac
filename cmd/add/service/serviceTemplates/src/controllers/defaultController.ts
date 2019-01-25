import { Request, Response, NextFunction } from 'express';
import { DefaultService } from '../services/defaultService';
import { DynamoDB } from '../database/dynamo.db';
import { Controller } from './controller.interface';
import { Service } from '../services/service.interface';

export class DefaultController implements Controller {
    constructor(public serviceInstance = new DefaultService()) { };
    get = async (req: Request, res: Response, next: NextFunction) => {
        try {
            let data = await this.serviceInstance.get(req.query);
            res.send(data)
        } catch (err) {
            next(err);
        }
    }

    getById = async (req: Request, res: Response, next: NextFunction) => {
        try {
            let data = await this.serviceInstance.getById(req.params.id);
            res.send(data)
        } catch (err) {
            next(err);
        }
    }

    post = async (req: Request, res: Response, next: NextFunction) => {
        try {
            let data = await this.serviceInstance.post(req.body);
            res.send(data)
        } catch (err) {
            next(err);
        }
    }
}