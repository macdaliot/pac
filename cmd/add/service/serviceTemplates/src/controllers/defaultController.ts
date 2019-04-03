import * as express from 'express';
import { Controller } from './controller.interface';
import { Database } from '../database/db.interface';

export class DefaultController implements Controller {
	constructor(dbInstance: Database) { this.dbInstance = dbInstance; };
	private dbInstance: Database;
	get = async (req: express.Request, res: express.Response, next: express.NextFunction) => {
		try {
			// if authentication is enabled (in routes.ts) then req.user will be 
			// populated. req.user.name will be a username, req.user.sub will be
			// an email address. req.user.groups will be an array of strings with
			// role names or group names populated from the SAML (Auth0) token
			//console.log(JSON.stringify(req.user.sub));
			let data = await this.dbInstance.query(req.query);
			res.send(data.Items)
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
			let data = await this.dbInstance.query({ id: req.params.id });
			res.send(data.Items)
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
			let data = await this.dbInstance.create(req.body);
			res.send(data)
		} catch (err) {
			next(err);
		}
	}
	update = async (req: express.Request, res: express.Response, next: express.NextFunction) => {
		try {
			// if authentication is enabled (in routes.ts) then req.user will be 
			// populated. req.user.name will be a username, req.user.sub will be
			// an email address. req.user.groups will be an array of strings with
			// role names or group names populated from the SAML (Auth0) token
			//console.log(JSON.stringify(req.user));
            let dbResponse = await this.dbInstance.update(this.getKeys(req), req.body);
			res.send(dbResponse)
		} catch (err) {
			next(err);
		}
	}
	getKeys = (req: express.Request) : { id: any } => {
		if (req.params && req.params.id){
			return { id: req.params.id };
		}
		else return { id: req.body.id };
	}
	delete = async (req: express.Request, res: express.Response, next: express.NextFunction) => {
		try {
			// if authentication is enabled (in routes.ts) then req.user will be 
			// populated. req.user.name will be a username, req.user.sub will be
			// an email address. req.user.groups will be an array of strings with
			// role names or group names populated from the SAML (Auth0) token
			//console.log(JSON.stringify(req.user));
			let data = await this.dbInstance.delete(this.getKeys(req));
			res.send(data)
		} catch (err) {
			next(err);
		}
	}
}