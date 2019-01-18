package service

import (
	"github.com/PyramidSystemsInc/go/files"
)

func CreateFunctionsTs(filePath string) {
	const template = `import { Logger } from '../middleware/logger/logger';
	let _logger = new Logger();
	export const getError = error => {
		/* put error mapping here */
		return 500;
	}
	export const errorHandler = (err, req, res, next) => {
		_logger.error(` + "`" + `error: ${err}` + "`" + `);
		res.status(getError(err)).send(err);
	}
	`
	files.CreateFromTemplate(filePath, template, nil)
}
