package setup

import (
  "time"

  "github.com/PyramidSystemsInc/go/aws"
  "github.com/PyramidSystemsInc/go/aws/s3"
  "github.com/PyramidSystemsInc/go/aws/kms"
  "github.com/PyramidSystemsInc/go/logger"
  "github.com/PyramidSystemsInc/go/str"
  "github.com/aws/aws-sdk-go/aws/session"
)

func TerraformS3Bucket(projectName string) (string, string) {
  region := "us-east-2"
  awsSession := aws.CreateAwsSession(region)
  projectFqdn := str.Concat(projectName, ".pac.pyramidchallenges.com")
  terraformS3Bucket := str.Concat("terraform.", projectFqdn)
  createBucket("terraform", "private", projectFqdn, projectName, region, awsSession)
  encryptionKeyID := kms.CreateEncryptionKey(awsSession, "pac-project", projectName)
  s3.EncryptBucket(terraformS3Bucket, encryptionKeyID)
  logger.Info("The S3 bucket for Terraform state has been created and encrypted")
  return projectFqdn, encryptionKeyID
}

func createBucket(suiteName string, access string, projectFqdn string, projectName string, region string, awsSession *session.Session) {
  frontEndFqdn := str.Concat(suiteName, ".", projectFqdn)
  s3.MakeBucket(frontEndFqdn, access, region, awsSession)
  time.Sleep(time.Second * 3)
  tagKey := "pac-project-name"
  s3.TagBucket(frontEndFqdn, tagKey, projectName, awsSession)
}
