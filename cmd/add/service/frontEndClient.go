package service

import (
  "strings"
	"github.com/PyramidSystemsInc/go/files"
)

func CreateFrontEndClient(filePath string, config map[string]string) {
  config["serviceNameUpperCase"] = strings.Title(config["serviceName"])
	const template = `import axios from 'axios';

export class {{.serviceNameUpperCase}} {
  get() {
    return new Promise(function(resolve, reject) {
      var serviceUrl = "http://{{.serviceUrl}}";
      axios.get(serviceUrl + "/api/{{.serviceName}}", {}).then(function(response) {
        resolve(response);
      });
    })
  }

  post() {
    return new Promise(function(resolve, reject) {
      var serviceUrl = "http://{{.serviceUrl}}";
      axios.post(serviceUrl + "/api/{{.serviceName}}", {}).then(function(response) {
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
	files.CreateFromTemplate(filePath, template, config)
}
