package service

import (
	"path/filepath"
	"strings"

	"github.com/PyramidSystemsInc/go/files"
)

func CreateFrontEndClient(fullPath string, config map[string]string) {
	files.EnsurePath(filepath.Dir(fullPath))

	config["serviceNameUpperCase"] = strings.Title(config["serviceName"])
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
	files.CreateFromTemplate(fullPath, template, config)
}
