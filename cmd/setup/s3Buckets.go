package setup

import (
  "github.com/PyramidSystemsInc/go/aws"
  "github.com/PyramidSystemsInc/go/aws/cloudfront"
  "github.com/PyramidSystemsInc/go/aws/route53"
  "github.com/PyramidSystemsInc/go/aws/s3"
  "github.com/PyramidSystemsInc/go/logger"
  "github.com/PyramidSystemsInc/go/str"
  "github.com/aws/aws-sdk-go/aws/session"
)

func S3Buckets(projectName string, projectFqdn string) {
  region := "us-east-2"
  awsSession := aws.CreateAwsSession(region)
  createBucket("integration", projectFqdn, region, awsSession)
  createBucket("demo", projectFqdn, region, awsSession)
}

func createBucket(suiteName string, projectFqdn string, region string, awsSession *session.Session) {
  frontEndFqdn := str.Concat(suiteName, ".", projectFqdn)
  s3.MakeBucket(frontEndFqdn, "public-read", region, awsSession)
  s3.EnableWebsiteHosting(frontEndFqdn, awsSession)
  cloudfrontFqdn := cloudfront.CreateDistributionFromS3Bucket(frontEndFqdn, awsSession)
  var ttl int64 = 300
  route53.ChangeRecord(projectFqdn, "CNAME", frontEndFqdn, []string{cloudfrontFqdn}, ttl, awsSession)
  bucketFqdn := str.Concat(frontEndFqdn, ".s3-website-", region, ".amazonaws.com")
  logger.Info(str.Concat("The ", suiteName, " suite will be available after running the front-end Jenkins job at ", bucketFqdn))
  logger.Info(str.Concat("The ", suiteName, " suite will be available in approx 15 minutes at ", frontEndFqdn))
}
