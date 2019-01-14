package frontEndFiles

import (
  "github.com/PyramidSystemsInc/go/files"
)

func CreateHeaderStyles(filePath string) {
  const template = `@import '../../scss/variables.scss';

.header-component {
  align-items: center;
  box-shadow: $box-shadow;
  display: flex;
  flex-basis: 56px;
  font-size: 18px;
  justify-content: space-around;
  left: 0px;
  min-height: 56px;
  position: absolute;
  top: 0px;
  transition: background-color 1.5s;
  user-select: none;
  width: 100%;

  .section {
    align-items: center;
    display: flex;
    flex-basis: 0;
    flex-grow: 1;
  }

  .section.left {
    justify-content: flex-start;
    margin-left: 14px;
  }

  .section.center {
    justify-content: center;
  }

  .section.right {
    justify-content: flex-end;
    margin-right: 14px;
  }

  .menu-button {
    border-radius: 24px;
    color: #eee;
    cursor: pointer;
    font-size: 24px;
    height: 24px;
    padding-top: 1px;
    padding-left: 8px;
    padding-right: 12px;
    width: 24px;
  }

  .menu-button:hover {
    color: white;
  }

  .page-title {
    color: white;
    font-weight: 700;
    line-height: 40px;
    margin-left: 5px;
  }

  .application-title {
    color: white;
    font-weight: 700;
  }

  .profile-button {
    border-radius: 5px;
    color: white;
    cursor: pointer;
    padding-left: 10px;
    padding-right: 10px;
  }

  .profile-button:hover {
    background-color: $hover-highlight;
  }

  .user-name {
    font-weight: 700;
    line-height: 40px;
    margin-left: 12px;
  }
}
`
  files.CreateFromTemplate(filePath, template, nil)
}
