package service

import (
	"path"
	"strings"

	"github.com/PyramidSystemsInc/go/directories"
	"github.com/PyramidSystemsInc/go/files"
)

func CreateFrontEndClient(fileName string, cfg map[string]string) {
	frontEndClientPath := "app/src/services"
	directories.Create(frontEndClientPath)
	cfg["serviceNameUpperCase"] = strings.Title(cfg["serviceName"])
	const template = `import axios from 'axios';
import { UrlConfig } from '../config';
export class {{.serviceNameUpperCase}} {
  async get() {
    return axios.get(UrlConfig.apiUrl + "{{.serviceName}}");
  }

  async post(data: any = {}) {
    return axios.post(UrlConfig.apiUrl + "{{.serviceName}}", data);
  }
}

/*
Sample Usage:

  const result = await new {{.serviceNameUpperCase}}().get();
  this.setState({
    devices: result
  });  
*/
`
	files.CreateFromTemplate(path.Join(frontEndClientPath, fileName), template, cfg)
}
