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

var CONFIG_FILE_NAME string = ".pac.json"

func GetRootDirectory() string {
	rootDirectory := files.FindUpTree(CONFIG_FILE_NAME)
	if rootDirectory != "" {
		return rootDirectory
	} else {
		errors.QuitIfError(errors.New("This is not a PAC project. Navigate to a PAC project and try again"))
	}
	return ""
}

func GoToRootProjectDirectory(projectName string) {
	workingDirectory := directories.GetWorking()
	projectDirectory := filepath.Join(workingDirectory, projectName)
	os.Chdir(projectDirectory)
}

func Read() map[string]*json.RawMessage {
	rootDirectory := GetRootDirectory()
	pacFileData, err := ioutil.ReadFile(path.Join(rootDirectory, CONFIG_FILE_NAME))
	errors.QuitIfError(err)
	var configData map[string]*json.RawMessage
	err = json.Unmarshal(pacFileData, &configData)
	errors.QuitIfError(err)
	return configData
}

func ReadAll() map[string]string {
	//fix scope, namespace etc.
	GoToRootProjectDirectory(Get("projectName"))

	// Open our jsonFile
	jsonFile, err := os.Open(".pac.json")

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]string
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

func Set(property string, value string) {
	configData := Read()
	newRawMessage := json.RawMessage(`"` + value + `"`)
	configData[property] = &newRawMessage
	newConfigData, err := json.Marshal(&configData)
	errors.QuitIfError(err)
	err = ioutil.WriteFile(path.Join(GetRootDirectory(), CONFIG_FILE_NAME), []byte(newConfigData), 0644)
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
