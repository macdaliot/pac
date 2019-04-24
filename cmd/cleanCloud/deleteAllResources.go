package cleanCloud

import (
	"os"

	"github.com/PyramidSystemsInc/go/aws"
	"github.com/PyramidSystemsInc/go/aws/kms"
	"github.com/PyramidSystemsInc/go/aws/resourcegroups"
	"github.com/PyramidSystemsInc/go/aws/s3"
	"github.com/PyramidSystemsInc/go/logger"
	"github.com/PyramidSystemsInc/go/str"
	"github.com/PyramidSystemsInc/go/terraform"
	"github.com/PyramidSystemsInc/pac/config"
)

// DeleteAllResources deletes the AWS resourced created by this application.
func DeleteAllResources() {
	os.Chdir(config.GetRootDirectory())

  // Initialize all Terraform template directories
  terraform.Initialize("terraform")
  terraform.Initialize("svc/terraform")

	// Destroy AWS resources managed by Terraform
   logger.Info("Terraform is destroying all AWS resources...")
   output := terraform.Destroy("services/terraform")
   logger.Info(output)

	output = terraform.Destroy("terraform")
	logger.Info(output)
	logger.Info("Terraform is finished destroying AWS resources")

	projectName := config.Get("projectName")
	region := "us-east-2"
	awsSession := aws.CreateAwsSession(region)

	s3.DeleteAllObjectVersions(config.Get("terraformS3Bucket"), awsSession)
	s3.DeleteAllDeleteMarkers(config.Get("terraformS3Bucket"), awsSession)
	s3.DeleteBucket(config.Get("terraformS3Bucket"), awsSession)

	resourcegroups.Create(projectName, "pac-project-name", projectName, awsSession)
	resourcegroups.DeleteAllResources(projectName, awsSession)

	// delete last so things encrypted with it are deleted first so nothing is orphaned and inaccessible
	kms.ScheduleEncryptionKeyDeletion(config.Get("encryptionKeyID"), awsSession)

	logger.Info(str.Concat("Finished deleting all AWS resources tagged as part of PAC project ", projectName))
}
