package config

import (
  "encoding/json"
  "io/ioutil"
  "path"
	"github.com/PyramidSystemsInc/go/errors"
	"github.com/PyramidSystemsInc/go/files"
	"github.com/PyramidSystemsInc/go/str"
)

var CONFIG_FILE_NAME string = ".pac"

func GetRootDirectory() string {
  rootDirectory := files.FindUpTree(CONFIG_FILE_NAME)
  if rootDirectory != "" {
    return rootDirectory
  } else {
    errors.QuitIfError(errors.New("This is not a PAC project. Navigate to a PAC project and try again"))
  }
  return ""
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
  if len(data) > 0 && data[len(data) - 1] == '"' {
    data = data[:len(data) - 1]
  }
  return data
}
