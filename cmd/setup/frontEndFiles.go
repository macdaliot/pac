package setup

import (
  "github.com/PyramidSystemsInc/pac/cmd/setup/frontEndFiles"
  "github.com/PyramidSystemsInc/go/logger"
  "github.com/PyramidSystemsInc/go/str"
)

func FrontEndFiles(projectDirectory string, projectName string, description string) {
  config := createConfig(projectName, description)
  frontEndFiles.CreatePackageJson(str.Concat(projectDirectory, "/app/package.json"), config)
  frontEndFiles.CreateWebpackConfig(str.Concat(projectDirectory, "/app/webpack.config.js"))
  frontEndFiles.CreateIndexHtml(str.Concat(projectDirectory, "/app/src/index.html"), config)
  frontEndFiles.CreateApplicationJs(str.Concat(projectDirectory, "/app/src/Application.js"))
  //createPyramidFavicon(str.Concat(projectDirectory, "/app/src/assets/png/favicon.png"))
  frontEndFiles.CreateScssVariables(str.Concat(projectDirectory, "/app/src/scss/variables.scss"))
  frontEndFiles.CreateScssMain(str.Concat(projectDirectory, "/app/src/scss/main.scss"))
  frontEndFiles.CreateRoutesJs(str.Concat(projectDirectory, "/app/src/routes/routes.js"))
  frontEndFiles.CreateRoutesJson(str.Concat(projectDirectory, "/app/src/routes/routes.json"))
  frontEndFiles.CreateHeaderComponent(str.Concat(projectDirectory, "/app/src/components/Header/Header.js"), config)
  frontEndFiles.CreateHeaderStyles(str.Concat(projectDirectory, "/app/src/components/Header/header.scss"))
  frontEndFiles.CreateSidebarComponent(str.Concat(projectDirectory, "/app/src/components/Sidebar/Sidebar.js"))
  frontEndFiles.CreateSidebarStyles(str.Concat(projectDirectory, "/app/src/components/Sidebar/sidebar.scss"))
  frontEndFiles.CreateSidebarButtonComponent(str.Concat(projectDirectory, "/app/src/components/Sidebar/parts/Button/Button.js"))
  frontEndFiles.CreateSidebarButtonStyles(str.Concat(projectDirectory, "/app/src/components/Sidebar/parts/Button/button.scss"))
  frontEndFiles.CreateNotFoundPageComponent(str.Concat(projectDirectory, "/app/src/components/pages/NotFound/NotFound.js"))
  frontEndFiles.CreateNotFoundPageStyles(str.Concat(projectDirectory, "/app/src/components/pages/NotFound/not-found.scss"))
  logger.Info("Created ReactJS front-end")
}

func createConfig(projectName string, description string) map[string]string {
  config := make(map[string]string)
  config["projectName"] = projectName
  config["description"] = description
  return config
}
