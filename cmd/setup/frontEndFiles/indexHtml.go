package frontEndFiles

import (
  "github.com/PyramidSystemsInc/go/files"
)

func CreateIndexHtml(filePath string, config map[string]string) {
  const template = `<!DOCTYPE html>
<html lang="en">
  <head>
    <title>{{.projectName}}</title>
    <base href="/" />
    <meta charset='utf-8'>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <!-- Google Analytics here -->
    <link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Merriweather:400,700|Source+Sans+Pro:400,700">
    <link rel="stylesheet" href="https://fonts.googleapis.com/icon?family=Material+Icons">
    <link rel="stylesheet" href="https://use.fontawesome.com/releases/v5.0.9/css/all.css" integrity="sha384-5SOiIsAziJl6AWe0HWRKTXlfcSHKmYV4RBF18PPJ173Kzn7jzMyFuTtk8JA7QQG1" crossorigin="anonymous">
    <link rel="stylesheet" href="/styles.css">
    <link rel='shortcut icon' type='image/x-icon' href='/png/favicon.png' />
  </head>
  <body>
    <div id="container"></div>
    <script src="/bundle.js"></script>
  </body>
</html>
`
  files.CreateFromTemplate(filePath, template, config)
}
