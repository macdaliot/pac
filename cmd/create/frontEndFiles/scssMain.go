package frontEndFiles

import (
	"github.com/PyramidSystemsInc/go/files"
)

func CreateScssMain(filePath string) {
  const template = `@import 'variables';

html, body {
  height: 100%;
  font-family: $font-sans-serif;
  font-size: $font-size;
  margin: 0;
  padding: 0;
}

#container {
  display: flex;
  height: 100%;
}

.app {
  display: flex;
  flex-basis: 0;
  flex-direction: column;
  flex-grow: 1;
}

.main {
  display: flex;
  flex-grow: 1;
  min-height: 500px;
}

.content {
  flex-grow: 1;
  padding-left: 12px;
  padding-right: 12px;
  padding-top: 66px;
}

#left-sidebar {
  flex-basis: 288px;
  transition: 0.4s;
}

#right-sidebar {
  flex-basis: 0px;
  transition: 0.4s;
}

a {
  text-decoration: none;
}

button {
  font-family: $font-sans-serif;
}
`
  files.CreateFromTemplate(filePath, template, nil)
}
