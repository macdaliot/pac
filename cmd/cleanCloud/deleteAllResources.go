package cleanCloud

import (
	"github.com/PyramidSystemsInc/go/aws"
	"github.com/PyramidSystemsInc/go/aws/kms"
	"github.com/PyramidSystemsInc/go/aws/resourcegroups"
	"github.com/PyramidSystemsInc/go/logger"
	"github.com/PyramidSystemsInc/go/str"
	"github.com/PyramidSystemsInc/go/terraform"
	"github.com/PyramidSystemsInc/pac/config"
)

// DeleteAllResources deletes the AWS resourced created by this application.
func DeleteAllResources() {
	// Destroy AWS resources managed by Terraform
  terraform.Destroy("svc/terraform")
  terraform.Destroy("terraform")

	projectName := config.Get("projectName")
	region := "us-east-2"
	awsSession := aws.CreateAwsSession(region)

	go kms.ScheduleEncryptionKeyDeletion(config.Get("encryptionKeyID"), awsSession)
	resourcegroups.Create(projectName, "pac-project-name", projectName, awsSession)
	resourcegroups.DeleteAllResources(projectName, awsSession)

	logger.Info(str.Concat("Finished deleting all AWS resources tagged as part of PAC project ", projectName))
}
