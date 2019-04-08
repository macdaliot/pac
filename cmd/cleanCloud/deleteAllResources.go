package cleanCloud

import (
	"github.com/PyramidSystemsInc/go/aws"
	"github.com/PyramidSystemsInc/go/aws/dynamodb"
	"github.com/PyramidSystemsInc/go/aws/kms"
	"github.com/PyramidSystemsInc/go/aws/resourcegroups"
	"github.com/PyramidSystemsInc/go/logger"
	"github.com/PyramidSystemsInc/go/str"
	"github.com/PyramidSystemsInc/pac/cmd/setup"
	"github.com/PyramidSystemsInc/pac/config"
)

// DeleteAllResources deletes the AWS resourced created by this application.
func DeleteAllResources() {
	//destroy AWS resources managed by Terraform
	setup.TerraformDestroy()

	projectName := config.Get("projectName")
	region := "us-east-2"
	awsSession := aws.CreateAwsSession(region)

	go kms.ScheduleEncryptionKeyDeletion(config.Get("encryptionKeyID"), awsSession)
	go dynamodb.DeleteTable("pac-"+projectName+"-terraform-locking-table", awsSession)
	resourcegroups.Create(projectName, "pac-project-name", projectName, awsSession)
	resourcegroups.DeleteAllResources(projectName, awsSession)

	logger.Info(str.Concat("Finished deleting all AWS resources tagged as part of PAC project ", projectName))
}
