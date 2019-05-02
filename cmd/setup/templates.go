package setup

import (
  "os"
  "path/filepath"

  "github.com/PyramidSystemsInc/go/aws/sts"
	"github.com/PyramidSystemsInc/go/directories"
	"github.com/PyramidSystemsInc/go/files"
	"github.com/PyramidSystemsInc/go/logger"
	"github.com/PyramidSystemsInc/go/terraform"
	"github.com/gobuffalo/packr"
)

// Templates - creates project directory, config files, and copies project files to project directory.
func Templates(projectName, description, gitAuth, awsRegion, encryptionKeyID, env string) string {
	createRootProjectDirectory(projectName)
  awsAccountID := sts.GetAccountID()
	cfg := createConfig(projectName, description, gitAuth, awsAccountID, awsRegion, encryptionKeyID, env)
	createProjectFiles(cfg)
	logger.Info("Created project structure")
  logger.Info("Installing node modules")
  NpmInstall()
  logger.Info("Finished installing node modules")
  return awsAccountID
}

func createRootProjectDirectory(projectName string) {
  workingDirectory := directories.GetWorking()
  projectDirectory := filepath.Join(workingDirectory, projectName)
  directories.Create(projectDirectory)
  os.Chdir(projectDirectory)
}

func createConfig(projectName, description, gitAuth, awsAccountID, awsRegion, encryptionKeyID, env string) map[string]string {
	cfg := make(map[string]string)
  cfg["awsID"] = awsAccountID
	cfg["projectName"] = projectName
	cfg["description"] = description
	cfg["encryptionKeyID"] = encryptionKeyID
	cfg["env"] = env
	cfg["gitAuth"] = gitAuth
	cfg["region"] = awsRegion
	cfg["terraformAWSVersion"] = terraform.AWSVersion
	cfg["terraformTemplateVersion"] = terraform.TemplateVersion
	return cfg
}

func createProjectFiles(cfg map[string]string) {
  box := packr.NewBox("./templates")
  for _, templatePath := range box.List() {
    logger.Info(templatePath)
    files.EnsurePath(filepath.Dir(templatePath))
    template, _ := box.FindString(templatePath)
    if templatePath != "core/templates/routes.template.ts" {
      files.CreateFromTemplate(templatePath, template, cfg)
    } else {
      files.Write(templatePath, []byte(template))
    }
  }
}
