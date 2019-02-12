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
  get() {
    return new Promise(function(resolve, reject) {
      axios.get(UrlConfig.apiUrl + "/api/{{.serviceName}}", {}).then(function(response) {
        resolve(response);
      });
    })
  }

  post() {
    return new Promise(function(resolve, reject) {
      axios.post(UrlConfig.apiUrl + "/api/{{.serviceName}}", {}).then(function(response) {
        resolve(response);
      });
    })
  }
}

/*
Sample Usage:

  new {{.serviceNameUpperCase}}().get().then(function(result) {
    this.setState({
      {{.serviceName}}: result
    });
  }.bind(this));
*/
`
	files.CreateFromTemplate(path.Join(frontEndClientPath, fileName), template, cfg)
}
