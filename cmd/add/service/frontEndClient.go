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
  template := `import axios from 'axios';
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
  template = `
import * as React from 'react';
import mockAxios from 'axios';
import { {{.serviceNameUpperCase}} } from './{{.serviceNameUpperCase}}';

describe('{{.serviceNameUpperCase}} service', () => {
  it('should call axios.get() when its own get() is called', () => {
    new {{.serviceNameUpperCase}}().get();
    expect(mockAxios.get).toBeCalledTimes(1);
  });
  it('should call axios.post() when its own post() is called', () => {
    new {{.serviceNameUpperCase}}().post({});
    expect(mockAxios.post).toBeCalledTimes(1);
  });
});
`
  fileName = strings.Replace(fileName, ".ts", ".spec.tsx", -1)
  files.CreateFromTemplate(path.Join(frontEndClientPath, fileName), template, cfg)
}
