package deploy

import (
  "time"

  "github.com/PyramidSystemsInc/go/aws"
  "github.com/PyramidSystemsInc/go/aws/s3"
  "github.com/PyramidSystemsInc/go/logger"
  "github.com/PyramidSystemsInc/go/str"
  "github.com/aws/aws-sdk-go/aws/session"
)

// TerraformS3Bucket - Creates a Terraform S3 bucket
func TerraformS3Bucket(projectName string, encryptionKeyID string) (string, string) {
  region := "us-east-2"
  awsSession := aws.CreateAwsSession(region)

  projectFqdn := str.Concat(projectName, ".pac.pyramidchallenges.com")
  terraformS3Bucket := str.Concat("terraform.", projectFqdn)
  createBucket("terraform", "private", projectFqdn, projectName, region, awsSession)

  s3.EncryptBucket(terraformS3Bucket, encryptionKeyID)
  logger.Info("The S3 bucket for Terraform state has been created and encrypted")

  s3.EnableVersioning(terraformS3Bucket)
  logger.Info("Versioning has been enabled on the S3 bucket")

  return terraformS3Bucket, projectFqdn
}

func createBucket(suiteName string, access string, projectFqdn string, projectName string, region string, awsSession *session.Session) {
  frontEndFqdn := str.Concat(suiteName, ".", projectFqdn)
  s3.MakeBucket(frontEndFqdn, access, region, awsSession)
  time.Sleep(time.Second * 3)
  tagKey := "pac-project-name"
  s3.TagBucket(frontEndFqdn, tagKey, projectName, awsSession)
}
