package frontEndFiles

import (
	"github.com/PyramidSystemsInc/go/files"
)

func CreateSidebarStyles(filePath string) {
  const template = `@import '../../scss/variables.scss';

.sidebar-component {
  background-color: white;
  box-shadow: $box-shadow;
  display: flex;
  flex-direction: column;
  overflow: hidden;
  padding-top: 66px;
  transition: flex-basis 0.7s, padding-left 1.0s, padding-right 1.0s;

  .sidebar-element {
    display: flex;
    flex-direction: column;
  }
}

.collapsed {
  flex-basis: 0px;
  padding-left: 0px;
  padding-right: 0px;
}

.expanded {
  flex-basis: 250px;
  padding-left: 6px;
  padding-right: 6px;
}
`
  files.CreateFromTemplate(filePath, template, nil)
}
