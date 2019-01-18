package service

import (
	"github.com/PyramidSystemsInc/go/files"
)

// CreateServerTs creates a server.ts file based on the configuration passed in
func CreateServerTs(filePath string, config map[string]string) {
	const template = `import * as express from 'express';
  import apiRouter from './routes/routes'
  import * as cors from 'cors';
  import * as loggerMiddleware from './middleware/logger/loggerMiddleware';
  const app = express();
  const port = 3000;
  /* TODO: update error handling */
  /* need configMap here */
  app
      .use(cors())
  
      /* parse middleware */
      /* https://expressjs.com/en/api.html#express.json */
      .use(express.json())
      .use(express.urlencoded({ extended: true }))
  
      /* logging middleware, order matters */
      .use('', loggerMiddleware._loggers)
  
      /* routes */
      .use('/api', apiRouter)
  
  app.listen(port, () => console.log(` + "`" + `{{.serviceName}} is running on port ${port}!` + "`" + `))
`
	files.CreateFromTemplate(filePath, template, config)
}
