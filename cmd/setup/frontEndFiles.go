package setup

import (
  "github.com/PyramidSystemsInc/pac/cmd/setup/frontEndFiles"
  "github.com/PyramidSystemsInc/go/logger"
)

func FrontEndFiles(projectDirectory string, projectName string, description string) {
  config := createConfig(projectName, description)
  frontEndFiles.CreatePackageJson(projectDirectory + "/app/package.json", config)
  frontEndFiles.CreateWebpackConfig(projectDirectory + "/app/webpack.config.js")
  frontEndFiles.CreateIndexHtml(projectDirectory + "/app/src/index.html", config)
  frontEndFiles.CreateApplicationJs(projectDirectory + "/app/src/Application.js")
  //createPyramidFavicon(projectDirectory + "/app/src/assets/png/favicon.png")
  frontEndFiles.CreateScssVariables(projectDirectory + "/app/src/scss/variables.scss")
  frontEndFiles.CreateScssMain(projectDirectory + "/app/src/scss/main.scss")
  frontEndFiles.CreateRoutesJs(projectDirectory + "/app/src/routes/routes.js")
  frontEndFiles.CreateRoutesJson(projectDirectory + "/app/src/routes/routes.json")
  frontEndFiles.CreateHeaderComponent(projectDirectory + "/app/src/components/Header/Header.js", config)
  frontEndFiles.CreateHeaderStyles(projectDirectory + "/app/src/components/Header/header.scss")
  frontEndFiles.CreateSidebarComponent(projectDirectory + "/app/src/components/Sidebar/Sidebar.js")
  frontEndFiles.CreateSidebarStyles(projectDirectory + "/app/src/components/Sidebar/sidebar.scss")
  frontEndFiles.CreateSidebarButtonComponent(projectDirectory + "/app/src/components/Sidebar/parts/Button/Button.js")
  frontEndFiles.CreateSidebarButtonStyles(projectDirectory + "/app/src/components/Sidebar/parts/Button/button.scss")
  frontEndFiles.CreateNotFoundPageComponent(projectDirectory + "/app/src/components/pages/NotFound/NotFound.js")
  frontEndFiles.CreateNotFoundPageStyles(projectDirectory + "/app/src/components/pages/NotFound/not-found.scss")
  logger.Info("Created ReactJS front-end")
}

func createConfig(projectName string, description string) map[string]string {
  config := make(map[string]string)
  config["projectName"] = projectName
  config["description"] = description
  return config
}
