package service

import (
	"github.com/PyramidSystemsInc/go/files"
)

func CreateDefaultControllerTs(filePath string) {
	const template = `import * as express from 'express';
	import { DefaultService } from '../services/defaultService';
	import { DynamoDB } from '../database/dynamo.db';
	import { Controller } from './controller.interface';
	import { Service } from '../services/service.interface';
	
	export class DefaultController implements Controller {
		constructor(public serviceInstance = new DefaultService()) { };
		get = async (req: express.Request, res: express.Response, next: express.NextFunction) => {
			try {
				let data = await this.serviceInstance.get(req.query);
				res.send(data)
			} catch (err) {
				next(err);
			}
		}
	
		getById = async (req: express.Request, res: express.Response, next) => {
			try {
				let data = await this.serviceInstance.getById(req.params.id);
				res.send(data)
			} catch (err) {
				next(err);
			}
		}
	
		post = async (req: express.Request, res: express.Response, next) => {
			try {
				let data = await this.serviceInstance.post(req.body);
				res.send(data)
			} catch (err) {
				next(err);
			}
		}
	}`
	files.CreateFromTemplate(filePath, template, nil)
}
