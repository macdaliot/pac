package frontEndFiles

import (
  "github.com/PyramidSystemsInc/go/files"
)

func CreateRoutesJson(filePath string) {
  const template = `[
]
`
  files.CreateFromTemplate(filePath, template, nil)
}
