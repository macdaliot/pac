package setup

import (
	"github.com/PyramidSystemsInc/go/aws"
	"github.com/PyramidSystemsInc/go/aws/kms"
	"github.com/PyramidSystemsInc/go/logger"
	"github.com/PyramidSystemsInc/pac/config"
)

// CreateEncryptionKey - Creates an encryption key which will secure the Terraform state
func CreateEncryptionKey() string {
	projectName := config.Get("projectName")
	region := "us-east-2"
	awsSession := aws.CreateAwsSession(region)
	encryptionKeyID := kms.CreateEncryptionKey(awsSession, "pac-project", projectName)
	logger.Info("Created encryption key to secure future Terraform state")
	return encryptionKeyID
}
