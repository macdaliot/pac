package deploy

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/PyramidSystemsInc/go/errors"
	"github.com/PyramidSystemsInc/go/logger"
	"github.com/PyramidSystemsInc/go/str"
)

type CreateWebhookRequest struct {
	Name   string                     `json:"name"`
	Active bool                       `json:"active"`
	Events []string                   `json:"events"`
	Config CreateWebhookRequestConfig `json:"config"`
}

type CreateWebhookRequestConfig struct {
	Url         string `json:"url"`
	ContentType string `json:"content_type"`
}

func GitHubWebhook(projectName string, gitAuth string, jenkinsURL string) {
	hooksApiEndpoint := str.Concat("https://api.github.com/repos/PyramidSystemsInc/", projectName, "/hooks")
	basicAuth := str.Concat("Basic ", gitAuth)
	request, err := http.NewRequest("GET", hooksApiEndpoint, nil)
	errors.LogIfError(err)
	request.Header.Add("Authorization", basicAuth)
	client := &http.Client{}
	response, err := client.Do(request)
	errors.LogIfError(err)
	defer response.Body.Close()
	bodyData, err := ioutil.ReadAll(response.Body)
	if string(bodyData) == "[]" {
		webhookRequest := createWebhookRequestBody(jenkinsURL)
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

func createWebhookRequestBody(jenkinsURL string) *bytes.Buffer {
	webhookRequest, err := json.Marshal(CreateWebhookRequest{
		Name:   "web",
		Active: true,
		Events: []string{
			"push",
		},
		Config: CreateWebhookRequestConfig{
			ContentType: "json",
			Url:         str.Concat("https://", jenkinsURL, "/github-webhook/"),
		},
	})
	errors.LogIfError(err)
	return bytes.NewBuffer(webhookRequest)
}
