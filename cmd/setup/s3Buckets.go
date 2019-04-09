package setup

import (
  "github.com/PyramidSystemsInc/go/aws"
  // "github.com/PyramidSystemsInc/go/aws/cloudfront"
  // "github.com/PyramidSystemsInc/go/aws/route53"
  "github.com/PyramidSystemsInc/go/aws/s3"
  "github.com/PyramidSystemsInc/go/logger"
  "github.com/PyramidSystemsInc/go/str"
  "github.com/PyramidSystemsInc/pac/config"
  "github.com/aws/aws-sdk-go/aws/session"
)

func S3Buckets(projectName string) {
  region := "us-east-2"
  awsSession := aws.CreateAwsSession(region)
  projectFqdn := config.Get("projectFqdn")
  createBucket("integration", projectFqdn, projectName, region, awsSession)
  createBucket("demo", projectFqdn, projectName, region, awsSession)
}

func createBucket(suiteName string, projectFqdn string, projectName string, region string, awsSession *session.Session) {
  frontEndFqdn := str.Concat(suiteName, ".", projectFqdn)
  s3.MakeBucket(frontEndFqdn, "public-read", region, awsSession)
  s3.EnableWebsiteHosting(frontEndFqdn, awsSession)
  tagKey := "pac-project-name"
  s3.TagBucket(frontEndFqdn, tagKey, projectName, awsSession)
  // cloudfrontFqdn := cloudfront.CreateDistributionFromS3Bucket(frontEndFqdn, awsSession)
  // cloudfront.TagDistribution(cloudfrontFqdn, tagKey, projectName, awsSession)
  // var ttl int64 = 300
  // route53.ChangeRecord(projectFqdn, "CNAME", frontEndFqdn, []string{cloudfrontFqdn}, ttl, awsSession)
  bucketFqdn := str.Concat(frontEndFqdn, ".s3-website-", region, ".amazonaws.com")
  logger.Info(str.Concat("The ", suiteName, " suite will be available after running the front-end Jenkins job at ", bucketFqdn))
  logger.Info(str.Concat("The ", suiteName, " suite will be available in approx 15 minutes at ", frontEndFqdn))
}
