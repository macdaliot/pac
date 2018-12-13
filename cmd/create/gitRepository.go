package create

import (
  "github.com/PyramidSystemsInc/go/commands"
  "github.com/PyramidSystemsInc/go/errors"
  "github.com/PyramidSystemsInc/go/logger"
  "bytes"
  "encoding/base64"
  "encoding/json"
  "net/http"
)

type CreateRepoRequest struct {
  Name         string  `json:"name"`
  Private      bool    `json:"private"`
  Description  string  `json:"description"`
}

func GitRepository(projectName string, gitUser string, gitPass string, projectDirectory string) {
  if (gitUser != "" && gitPass != "") {
    repoConfig := createRepoConfig(projectName)
    postToGitHub(repoConfig, gitUser, gitPass)
    setupRepository(projectName, projectDirectory)
    logger.Info("Created GitHub repository")
  } else {
    logger.Warn("Skipping creation of GitHub repository due to one or more of the following being blank: --gitUser, --gitPass")
  }
}

func createRepoConfig(projectName string) *bytes.Buffer {
  repoConfig, err := json.Marshal(CreateRepoRequest{
    Name: projectName,
    Private: true,
    Description: projectName + " project (created by PAC)",
  })
  errors.LogIfError(err)
  return bytes.NewBuffer(repoConfig)
}

func postToGitHub(repoConfig *bytes.Buffer, gitUser string, gitPass string) {
  request, err := http.NewRequest("POST", "https://api.github.com/orgs/PyramidSystemsInc/repos", repoConfig)
  errors.LogIfError(err)
  basicAuth := base64.StdEncoding.EncodeToString([]byte(gitUser + ":" + gitPass))
  request.Header.Add("Authorization", "Basic " + basicAuth)
  client := &http.Client{}
  response, err := client.Do(request)
  errors.LogIfError(err)
  defer response.Body.Close()
}

func setupRepository(projectName string, projectDirectory string) {
  commands.Run("git init", projectDirectory)
  commands.Run("git remote add origin git@github.com:PyramidSystemsInc/" + projectName + ".git", projectDirectory)
}
