import * as express from 'express';
import { DefaultService } from '../services/defaultService';
import { DynamoDB } from '../database/dynamo.db';
import { Controller } from './controller.interface';
import { Service } from '../services/service.interface';

export class DefaultController implements Controller {
	constructor(public serviceInstance = new DefaultService()) { };
	get = async (req: express.Request, res: express.Response, next: express.NextFunction) => {
		try {
			// if authentication is enabled (in routes.ts) then req.user will be 
			// populated. req.user.name will be a username, req.user.sub will be
			// an email address. req.user.groups will be an array of strings with
			// role names or group names populated from the SAML (Auth0) token
			//console.log(JSON.stringify(req.user.sub));
			let data = await this.serviceInstance.get(req.query);
			res.send(data)
		} catch (err) {
			next(err);
		}
	}

	getById = async (req: express.Request, res: express.Response, next: express.NextFunction) => {
		try {
			// if authentication is enabled (in routes.ts) then req.user will be 
			// populated. req.user.name will be a username, req.user.sub will be
			// an email address. req.user.groups will be an array of strings with
			// role names or group names populated from the SAML (Auth0) token
			//console.log(JSON.stringify(req.user));
			let data = await this.serviceInstance.getById(req.params.id);
			res.send(data)
		} catch (err) {
			next(err);
		}
	}

	post = async (req: express.Request, res: express.Response, next: express.NextFunction) => {
		try {
			// if authentication is enabled (in routes.ts) then req.user will be 
			// populated. req.user.name will be a username, req.user.sub will be
			// an email address. req.user.groups will be an array of strings with
			// role names or group names populated from the SAML (Auth0) token
			//console.log(JSON.stringify(req.user));
			let data = await this.serviceInstance.post(req.body);
			res.send(data)
		} catch (err) {
			next(err);
		}
	}
}