package setup

import (
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/PyramidSystemsInc/pac/config"
)

// CopyBinaries copies files that aren't handled well by packr to the Terraform directories
func CopyBinaries() {
	pacRoot := os.Getenv("GOPATH") + "/src/github.com/PyramidSystemsInc/pac/cmd/setup"
	projectRoot := config.GetRootDirectory()

	// Copy cmd/setup/binaries/cwl2eslambda.zip to <project>/services/terraform
	copy(
		filepath.Join(pacRoot, "binaries/services/terraform/cwl2eslambda.zip"),
		filepath.Join(projectRoot, "services/terraform/cwl2eslambda.zip"))

	// Copy cmd/setup/binaries/dynamoDbToElasticSearch.zip to <project>/services/terraform
	copy(
		filepath.Join(pacRoot, "binaries/services/terraform/dynamoDbToElasticSearch.zip"),
		filepath.Join(projectRoot, "services/terraform/dynamoDbToElasticSearch.zip"))
}

func copy(src, dst string) {
	from, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
	}
	defer from.Close()

	to, err := os.OpenFile(dst, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer to.Close()

	_, err = io.Copy(to, from)
	if err != nil {
		log.Fatal(err)
	}
}
