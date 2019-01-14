package frontEndFiles

import (
  "github.com/PyramidSystemsInc/go/files"
)

func CreateNotFoundPageStyles(filePath string) {
  const template = `@import '../../../scss/variables.scss';

.not-found-page-component {
  /* New selectors go here */
}
`
  files.CreateFromTemplate(filePath, template, nil)
}
