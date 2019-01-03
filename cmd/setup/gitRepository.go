package setup

import (
  "bytes"
  "encoding/json"
  "net/http"
  "github.com/PyramidSystemsInc/go/commands"
  "github.com/PyramidSystemsInc/go/errors"
  "github.com/PyramidSystemsInc/go/logger"
)

type CreateRepoRequest struct {
  Name         string  `json:"name"`
  Private      bool    `json:"private"`
  Description  string  `json:"description"`
}

func GitRepository(projectName string, gitAuth string, projectDirectory string) {
  repoConfig := createRepoConfig(projectName)
  postToGitHub(repoConfig, gitAuth)
  setupRepository(projectName, projectDirectory)
  logger.Info("Created GitHub repository")
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

func postToGitHub(repoConfig *bytes.Buffer, gitAuth string) {
  request, err := http.NewRequest("POST", "https://api.github.com/orgs/PyramidSystemsInc/repos", repoConfig)
  errors.LogIfError(err)
  request.Header.Add("Authorization", "Basic " + gitAuth)
  client := &http.Client{}
  response, err := client.Do(request)
  errors.LogIfError(err)
  defer response.Body.Close()
}

func setupRepository(projectName string, projectDirectory string) {
  commands.Run("git init", projectDirectory)
  commands.Run("git remote add origin git@github.com:PyramidSystemsInc/" + projectName + ".git", projectDirectory)
}
