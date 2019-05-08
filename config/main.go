package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/PyramidSystemsInc/go/directories"
	"github.com/PyramidSystemsInc/go/errors"
	"github.com/PyramidSystemsInc/go/files"
	"github.com/PyramidSystemsInc/go/str"
)

// ConfigFileName is a json formatted file used to store configuration variables
const ConfigFileName string = ".pac.json"

// CopyConfig copies the default .pac.json file to the project root
func CopyConfig(projectName string) {
	projectDir := os.Getenv("GOPATH") + "\\src\\" + projectName
	src := os.Getenv("GOPATH") + "\\src\\github.com\\PyramidSystemsInc\\pac\\" + ConfigFileName
	dest := projectDir + "\\" + ConfigFileName

	// Copy configuration file to templates directory
	input, err := ioutil.ReadFile(src)
	if err != nil {
		errors.LogIfError(err)
	}

	err = ioutil.WriteFile(dest, input, 0644)
	if err != nil {
		errors.LogIfError(err)
	}
}

// GetRootDirectory returns the path to the project root directory
func GetRootDirectory() string {
	rootDirectory := files.FindUpTree(ConfigFileName)

	if rootDirectory != "" {
		return rootDirectory
	} else {
		errors.QuitIfError(errors.New("This is not a PAC project. Navigate to a PAC project and try again"))
	}

	return ""
}

// GoToRootProjectDirectory changes the working directory to the project root directory
func GoToRootProjectDirectory(projectName string) {
	workingDirectory := directories.GetWorking()
	projectDirectory := filepath.Join(workingDirectory, projectName)
	os.Chdir(projectDirectory)
}

func Read() map[string]*json.RawMessage {
	rootDirectory := GetRootDirectory()

	GoToRootProjectDirectory(Get("projectName"))

	pacFileData, err := ioutil.ReadFile(path.Join(rootDirectory, ConfigFileName))
	errors.QuitIfError(err)

	var configData map[string]*json.RawMessage

	err = json.Unmarshal(pacFileData, &configData)
	errors.QuitIfError(err)

	return configData
}

// ReadAll changes the current working directory to the project root and returns the .pac.json file as a map[string]string
func ReadAll() map[string]string {
	// TODO: fix scope, namespace etc.
	path := GetRootDirectory()

	// Open our jsonFile
	jsonFile, err := os.Open(path + "\\.pac.json")

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	result := make(map[string]string)

	json.Unmarshal([]byte(byteValue), &result)

	return result
}

func Get(property string) string {
	configData := Read()
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

// Preset sets values in .pac.json in the project template prior to deploy
func Preset(property string, value string) {
	configData := ReadAll()
	configData[property] = value
	newConfigData, err := json.Marshal(&configData)
	errors.QuitIfError(err)

	destinationFile := os.Getenv("GOPATH") + "\\src\\github.com\\PyramidSystemsInc\\pac\\cmd\\setup\\templates\\.pac.json"

	err = ioutil.WriteFile(destinationFile, []byte(newConfigData), 0644)
	errors.QuitIfError(err)
}

func Set(property string, value string) {
	configData := Read()
	newRawMessage := json.RawMessage(`"` + value + `"`)
	configData[property] = &newRawMessage
	newConfigData, err := json.Marshal(&configData)
	errors.QuitIfError(err)

	rootDir := GetRootDirectory()

	err = ioutil.WriteFile(path.Join(rootDir, ConfigFileName), []byte(newConfigData), 0644)
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
