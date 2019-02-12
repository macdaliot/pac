package setup

import (
  "io/ioutil"
  "bytes"
  "encoding/json"
  "net/http"
  "github.com/PyramidSystemsInc/go/errors"
  "github.com/PyramidSystemsInc/go/logger"
  "github.com/PyramidSystemsInc/go/str"
  "github.com/PyramidSystemsInc/pac/config"
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

func GitHubWebhook(projectName string) {
  hooksApiEndpoint := str.Concat("https://api.github.com/repos/PyramidSystemsInc/", config.Get("projectName"), "/hooks")
  basicAuth := str.Concat("Basic ", config.Get("gitAuth"))
  request, err := http.NewRequest("GET", hooksApiEndpoint, nil)
  errors.LogIfError(err)
  request.Header.Add("Authorization", basicAuth)
  client := &http.Client{}
  response, err := client.Do(request)
  errors.LogIfError(err)
  defer response.Body.Close()
  bodyData, err := ioutil.ReadAll(response.Body)
  if string(bodyData) == "[]" {
    webhookRequest := createWebhookRequestBody(config.Get("jenkinsUrl"))
    request, err = http.NewRequest("POST", hooksApiEndpoint, webhookRequest)
    errors.LogIfError(err)
    request.Header.Add("Authorization", basicAuth)
    response, err = client.Do(request)
    errors.LogIfError(err)
    defer response.Body.Close()
    logger.Info("Created webhook to Jenkins on GitHub.com")
  } else {
    logger.Info("Skipping GitHub webhook either because it already exists or there is no repository")
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
      Url: str.Concat("http://", jenkinsUrl, "/github-webhook/"),
      ContentType: "json",
    },
  })
  errors.LogIfError(err)
  return bytes.NewBuffer(webhookRequest)
}
