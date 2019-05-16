package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/PyramidSystemsInc/go/directories"
	"github.com/PyramidSystemsInc/go/errors"
	"github.com/PyramidSystemsInc/go/files"
	"github.com/PyramidSystemsInc/go/str"
)

const configFileName string = ".pac.json"

// Create creates the .pac.json file in the project root
func Create() {
	files.CreateBlank(configFileName)
	files.Write(configFileName, []byte("{}"))
}

// GetRootDirectory returns the path to the project root directory
func GetRootDirectory() string {
	rootDirectory := files.FindUpTree(configFileName)

	if rootDirectory == "" {
		errors.QuitIfError(errors.New("This is not a PAC project. Navigate to a PAC project and try again"))
	}

	return rootDirectory
}

// GoToRootProjectDirectory changes the working directory to the project root directory
func GoToRootProjectDirectory(projectName string) {
	workingDirectory := directories.GetWorking()
	projectDirectory := filepath.Join(workingDirectory, projectName)
	os.Chdir(projectDirectory)
}

func read() map[string]*json.RawMessage {
	rootDirectory := GetRootDirectory()
	pacFileData, err := ioutil.ReadFile(path.Join(rootDirectory, configFileName))
	errors.QuitIfError(err)

	var configData map[string]*json.RawMessage

	err = json.Unmarshal(pacFileData, &configData)
	errors.QuitIfError(err)

	return configData
}

// ReadAll changes the current working directory to the project root and returns the .pac.json file as a map[string]string
func ReadAll() map[string]string {
	path := GetRootDirectory()
	pacJSON := files.Read(filepath.Join(path, configFileName))
	result := make(map[string]string)
	json.Unmarshal([]byte(pacJSON), &result)
	return result
}

// Get reads a value from the .pac.json file in the project root
func Get(property string) string {
	configData := read()
	if _, ok := configData[property]; ok {
		propertyRawData := *configData[property]
		propertyData, err := propertyRawData.MarshalJSON()
		errors.LogIfError(err)
		return removeQuotes(string(propertyData))
	} else {
		errors.QuitIfError(errors.New(str.Concat("Property ", property, " does not exist")))
		return ""
	}
}

// Set value in .pac.json in project root
func Set(property string, value string) {
	configData := read()
	newRawMessage := json.RawMessage(`"` + value + `"`)
	configData[property] = &newRawMessage
	newConfigData, err := json.Marshal(&configData)
	errors.QuitIfError(err)

	rootDir := GetRootDirectory()

	err = ioutil.WriteFile(path.Join(rootDir, configFileName), []byte(newConfigData), 0644)
	errors.QuitIfError(err)
}

func removeQuotes(data string) string {
	if len(data) > 0 && data[0] == '"' {
		data = data[1:]
	}
	if len(data) > 0 && data[len(data)-1] == '"' {
		data = data[:len(data)-1]
	}
	return data
}
