package frontEndFiles

import (
	"github.com/PyramidSystemsInc/go/files"
)

func CreatePackageJson(filePath string, config map[string]string) {
  const template = `{
  "name": "{{.projectName}}",
  "version": "0.0.1",
  "description": "{{.description}}",
  "main": "src/Application.js",
  "scripts": {
    "build": "webpack --config webpack.config.js",
    "dev-server": "webpack-dev-server --hot --history-api-fallback"
  },
  "author": "Pyramid Systems, Inc.",
  "license": "ISC",
  "dependencies": {
    "@babel/core": "^7.2.0",
    "@babel/plugin-transform-runtime": "^7.2.0",
    "@babel/runtime": "^7.2.0",
    "@jeff.diederiks/pyramid-react": "0.0.28",
    "autoprefixer": "^9.4.2",
    "aws-sdk": "^2.373.0",
    "axios": "^0.18.0",
    "babel-cli": "^6.26.0",
    "babel-core": "^6.26.3",
    "babel-loader": "^7.1.3",
    "babel-plugin-syntax-dynamic-import": "^6.18.0",
    "babel-plugin-transform-es2015-arrow-functions": "^6.22.0",
    "babel-plugin-transform-es2015-classes": "^6.24.1",
    "babel-plugin-transform-es2015-template-literals": "^6.22.0",
    "babel-plugin-transform-es3-member-expression-literals": "^6.22.0",
    "babel-plugin-transform-es3-property-literals": "^6.22.0",
    "babel-plugin-transform-runtime": "^6.23.0",
    "babel-preset-env": "^1.7.0",
    "babel-preset-react": "^6.24.1",
    "css-loader": "^2.0.0",
    "node-sass": "^4.11.0",
    "postcss-cli": "^6.0.1",
    "postcss-loader": "^3.0.0",
    "promise-polyfill": "^8.1.0",
    "react": "^16.6.3",
    "react-dom": "^16.6.3",
    "react-router": "^4.3.1",
    "react-router-config": "^4.4.0-beta.6",
    "react-router-dom": "^4.3.1",
    "sass-loader": "^7.1.0",
    "style-loader": "^0.23.1",
    "url-loader": "^1.1.2",
    "webpack": "^4.27.1"
  },
  "devDependencies": {
    "copy-webpack-plugin": "^4.6.0",
    "css-hot-loader": "^1.4.2",
    "extract-text-webpack-plugin": "^4.0.0-beta.0",
    "react-hot-loader": "^4.3.12",
    "webpack-cli": "^3.1.2",
    "webpack-dev-server": "^3.1.10"
  }
}
`
  files.CreateFromTemplate(filePath, template, config)
}
