package frontEndFiles

import (
	"github.com/PyramidSystemsInc/go/files"
)

func CreateNotFoundPageComponent(filePath string) {
  const template = `import React from 'react';
import { hot } from 'react-hot-loader';
import './not-found.scss';

class NotFound extends React.Component {
  render() {
    return (
      <div className="not-found-page-component" />
    );
  }
}

export default hot(module)(NotFound);
`
  files.CreateFromTemplate(filePath, template, nil)
}
