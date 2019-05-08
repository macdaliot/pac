package config

import (
	"encoding/json"
	"path/filepath"
	"io/ioutil"
	"path"
	"github.com/PyramidSystemsInc/go/errors"
	"github.com/PyramidSystemsInc/go/files"
	"github.com/PyramidSystemsInc/go/str"
)

const ConfigFileName string = ".pac.json"

func Create() {
	files.CreateBlank(ConfigFileName)
	files.Write(ConfigFileName, []byte("{}"))
}

func GetRootDirectory() string {
	rootDirectory := files.FindUpTree(ConfigFileName)
	if rootDirectory != "" {
		return rootDirectory
	} else {
		errors.QuitIfError(errors.New("This is not a PAC project. Navigate to a PAC project and try again"))
	}
	return ""
}

func Read() map[string]*json.RawMessage {
	rootDirectory := GetRootDirectory()
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
	pacJson := files.Read(filepath.Join(path, ".pac.json"))
  result := make(map[string]string)
  json.Unmarshal([]byte(pacJson), &result)
  return result

  // Open our jsonFile
  // jsonFile, err := os.Open(path + "\\.pac.json")

  // defer the closing of our jsonFile so that we can parse it later on
  //defer jsonFile.Close()

  // byteValue, _ := ioutil.ReadAll(jsonFile)

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
	err = ioutil.WriteFile(path.Join(GetRootDirectory(), ConfigFileName), []byte(newConfigData), 0644)
	errors.QuitIfError(err)
}

func removeQuotes(data string) string {
	if len(data) > 0 && data[0] == '"' {
		data = data[1:]
	}
	if len(data) > 0 && data[len(data) - 1] == '"' {
		data = data[:len(data) - 1]
	}
	return data
}
