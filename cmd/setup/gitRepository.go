package setup

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/PyramidSystemsInc/go/commands"
	"github.com/PyramidSystemsInc/go/errors"
	"github.com/PyramidSystemsInc/go/logger"
	"github.com/PyramidSystemsInc/go/str"
	"github.com/PyramidSystemsInc/pac/config"
)

type CreateRepoRequest struct {
	Name        string `json:"name"`
	Private     bool   `json:"private"`
	Description string `json:"description"`
}

func GitRepository(projectName string) {
	repoConfig := createRepoConfig(projectName)
	postToGitHub(repoConfig)
	setupRepository(projectName)
	logger.Info("Created GitHub repository")
}

func createRepoConfig(projectName string) *bytes.Buffer {
	repoConfig, err := json.Marshal(CreateRepoRequest{
		Name:        projectName,
		Private:     true,
		Description: config.Get("description"),
	})
	errors.LogIfError(err)
	return bytes.NewBuffer(repoConfig)
}

func postToGitHub(repoConfig *bytes.Buffer) {
	request, err := http.NewRequest("POST", "https://api.github.com/orgs/PyramidSystemsInc/repos", repoConfig)
	errors.LogIfError(err)
	request.Header.Add("Authorization", str.Concat("Basic ", config.Get("gitAuth")))
	client := &http.Client{}
	response, err := client.Do(request)
	errors.LogIfError(err)
	defer response.Body.Close()
}

func setupRepository(projectName string) {
	commands.Run("git init", "")
	commands.Run(str.Concat("git remote add origin git@github.com:PyramidSystemsInc/", projectName, ".git"), "")
}
