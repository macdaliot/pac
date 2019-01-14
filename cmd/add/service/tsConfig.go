package service

import (
  "github.com/PyramidSystemsInc/go/files"
)

// createTsConfig creates a default tsConfig.json based on passed in config
//TODO: configure sourcemap, mapRoot for debugger
func CreateTsConfig(filePath string, config map[string]string) {
  const template = `{
  "compilerOptions": {
    "module": "commonjs",
    "target": "es6",
    "noImplicitAny": false,
    "sourceMap": false,
    "outDir": "dist/",
    "allowJs": true,
    "types": [
      "node"
    ],
    "typeRoots": [
      "../node_modules/@types"
    ]
  },
  "include": [
    "src"
  ],
  "exclude": [
    "node_modules",
    "tests",
    "dist"
  ]
}
`
  files.CreateFromTemplate(filePath, template, config)
}
