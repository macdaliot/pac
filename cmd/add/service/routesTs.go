package service

import (
	"github.com/PyramidSystemsInc/go/files"
)

func CreateRoutesTS(filePath string, config map[string]string) {
	const template = `import { DefaultController } from '../controllers/defaultController';
	import { errorHandler } from '../utility';
	import * as express from 'express';
	let apiRouter = express.Router();
	let defaultController = new DefaultController();
	apiRouter
		.get('/{{.serviceName}}', defaultController.get)
		.get('/{{.serviceName}}/:id', defaultController.getById)
		.post('/{{.serviceName}}', defaultController.post)
	
		.use('*', errorHandler)
	export default apiRouter;`
	files.CreateFromTemplate(filePath, template, config)
}
