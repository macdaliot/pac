package service

import (
	"github.com/PyramidSystemsInc/go/files"
)

func CreateDockerfile(filePath string) {
  const template = `FROM node

WORKDIR "/app"
COPY . /app/
RUN cd /app

EXPOSE 3000 8000
CMD ["npm", "start"]
`
  files.CreateFromTemplate(filePath, template, nil)
}
