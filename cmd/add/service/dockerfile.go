package service

import (
  "github.com/PyramidSystemsInc/go/files"
)

// CreateDockerfile creates a docker file to run the service
func CreateDockerfile(filePath string) {
  const template = `FROM node

WORKDIR "/app"
COPY . /app/
RUN cd /app

EXPOSE 3000
CMD ["npm", "start"]
`
  files.CreateFromTemplate(filePath, template, nil)
}
