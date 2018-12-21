package create

import (
  "io/ioutil"
  "bytes"
  "encoding/json"
  "net/http"
  "github.com/PyramidSystemsInc/go/errors"
  "github.com/PyramidSystemsInc/go/logger"
)

type CreateWebhookRequest struct {
  Name    string                      `json:"name"`
  Active  bool                        `json:"active"`
  Events  []string                    `json:"events"`
  Config  CreateWebhookRequestConfig  `json:"config"`
}

type CreateWebhookRequestConfig struct {
  Url          string  `json:"url"`
  ContentType  string  `json:"content_type"`
}

type PacFile struct {
  ProjectName  string  `json:"projectName"`
  GitAuth      string  `json:"gitAuth"`
  JenkinsUrl   string  `json:"jenkinsUrl"`
}

func GitHubWebhook() {
  pacFile := readPacFile()
  createWebhookIfDoesNotExist(pacFile)
/*
  if (gitUser != "" && gitPass != "") {
    repoConfig := createRepoConfig(projectName)
    postToGitHub(repoConfig, gitUser, gitPass)
    setupRepository(projectName, projectDirectory)
    logger.Info("Created GitHub repository")
  } else {
    logger.Warn("Skipping creation of GitHub repository due to one or more of the following being blank: --gitUser, --gitPass")
  }
*/
}

func readPacFile() PacFile {
  // TODO: Should run from anywhere
  // TODO: Should not depend on pacFile for git
  var pacFile PacFile
  pacFileData, err := ioutil.ReadFile(".pac")
  errors.QuitIfError(err)
  json.Unmarshal(pacFileData, &pacFile)
  return pacFile
}

func createWebhookIfDoesNotExist(pacFile PacFile) {
  hooksApiEndpoint := "https://api.github.com/repos/PyramidSystemsInc/" + pacFile.ProjectName + "/hooks"
  basicAuth := "Basic " + pacFile.GitAuth
  request, err := http.NewRequest("GET", hooksApiEndpoint, nil)
  errors.LogIfError(err)
  request.Header.Add("Authorization", basicAuth)
  client := &http.Client{}
  response, err := client.Do(request)
  errors.LogIfError(err)
  defer response.Body.Close()
  bodyData, err := ioutil.ReadAll(response.Body)
  if string(bodyData) == "[]" {
    webhookRequest := createWebhookRequestBody(pacFile.JenkinsUrl)
    request, err = http.NewRequest("POST", hooksApiEndpoint, webhookRequest)
    errors.LogIfError(err)
    request.Header.Add("Authorization", basicAuth)
    response, err = client.Do(request)
    errors.LogIfError(err)
    defer response.Body.Close()
    logger.Info("Created webhook to Jenkins on GitHub.com")
  } else {
    logger.Info("Skipping GitHub webhook because it already exists")
  }
}

func createWebhookRequestBody(jenkinsUrl string) *bytes.Buffer {
  webhookRequest, err := json.Marshal(CreateWebhookRequest{
    Name: "web",
    Active: true,
    Events: []string{
      "push",
    },
    Config: CreateWebhookRequestConfig{
      Url: "http://" + jenkinsUrl + "/github-webhook/",
      ContentType: "json",
    },
  })
  errors.LogIfError(err)
  return bytes.NewBuffer(webhookRequest)
}

/*
func setupRepository(projectName string, projectDirectory string) {
  commands.Run("git init", projectDirectory)
  commands.Run("git remote add origin git@github.com:PyramidSystemsInc/" + projectName + ".git", projectDirectory)
}
*/
