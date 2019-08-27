package cleanCloud

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/PyramidSystemsInc/go/aws"
	"github.com/PyramidSystemsInc/go/aws/kms"
	"github.com/PyramidSystemsInc/go/aws/resourcegroups"
	"github.com/PyramidSystemsInc/go/aws/s3"
	"github.com/PyramidSystemsInc/go/files"
	"github.com/PyramidSystemsInc/go/logger"
	"github.com/PyramidSystemsInc/go/str"
	"github.com/PyramidSystemsInc/go/terraform"
	"github.com/PyramidSystemsInc/pac/config"
)

const terraformS3Bucket string = "terraformS3Bucket"

// DeleteAllResources deletes the AWS resourced created by this application.
func DeleteAllResources() {
	os.Chdir(config.GetRootDirectory())
	environmentNames := strings.Split(config.Get("environments"), ",")

	projectRoot := config.GetRootDirectory() + "/terraform/"

	for i, _ := range environmentNames {
		files.DirectoryFileStringReplace(projectRoot+environmentNames[i], "^lambda_", " source_code_hash", "#source_code_hash")
	}

	for _, environmentName := range environmentNames {
		logger.Info(str.Concat("Terraform is cleaning the ", environmentName, " VPC..."))
		terraformDir := filepath.Join("terraform/", environmentName)
		terraform.Initialize(terraformDir)
		terraform.Destroy(terraformDir)
		logger.Info(str.Concat("Terraform cleaned the ", environmentName, " VPC"))
	}

	// Initialize all remaining Terraform template directories
	terraform.Initialize("terraform/dns")
	terraform.Initialize("terraform/management")

	// Destroy the remaining AWS resources managed by Terraform
	logger.Info("Terraform is cleaning the management VPC...")
	output := terraform.Destroy("terraform/management")
	logger.Info(output)
	logger.Info("Terraform cleaned the management VPC")

	logger.Info("Terraform is cleaning the project's DNS records...")
	if config.Get("dnsPristine") == "true" {
		output = terraform.Destroy("terraform/dns_pristine")
	} else {
		output = terraform.Destroy("terraform/dns")
	}
	logger.Info(output)
	logger.Info("Terraform cleaned the project's DNS records")

	logger.Info("Terraform is finished destroying Terraform Managed AWS resources")

	projectName := config.Get("projectName")
	region := config.Get("region")
	awsSession := aws.CreateAwsSession(region)

	s3.DeleteAllObjectVersions(config.Get(terraformS3Bucket), awsSession)
	s3.DeleteAllDeleteMarkers(config.Get(terraformS3Bucket), awsSession)
	s3.DeleteBucket(config.Get(terraformS3Bucket), awsSession)

	resourcegroups.Create(projectName, "pac-project-name", projectName, awsSession)
	resourcegroups.DeleteAllResources(projectName, awsSession)

	// delete last so things encrypted with it are deleted first so nothing is orphaned and inaccessible
	kms.ScheduleEncryptionKeyDeletion(config.Get("encryptionKeyID"), awsSession)

	logger.Info(str.Concat("Finished deleting all AWS resources tagged as part of PAC project ", projectName))
}
