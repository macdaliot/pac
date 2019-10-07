package cmd_test

import (
	"os"
	"testing"

	"github.com/PyramidSystemsInc/pac/config"
)

func TestCopyConfig(t *testing.T) {
	projectName := "testdynesthree"
	projectDir := os.Getenv("GOPATH") + "\\src\\" + projectName
	dest := projectDir + "\\.pac.json"

	config.CopyConfig(projectName)

	_, err := os.Stat(dest)
	if err != nil {
		t.Error(err)
	}
}
