package service

import (
	"github.com/PyramidSystemsInc/go/files"
)

func CreateDockerfile(filePath string) {
  const template = `FROM node

WORKDIR "/app"
COPY package.json /app/
RUN cd /app
COPY . /app
RUN npm i

EXPOSE 3000
CMD ["npm", "start"]
`
  files.CreateFromTemplate(filePath, template, nil)
}
