package setup

import (
	"time"

	"github.com/PyramidSystemsInc/go/aws"
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
	createBucket("terraform", "private", projectFqdn, projectName, region, awsSession)
}

func createBucket(suiteName string, access string, projectFqdn string, projectName string, region string, awsSession *session.Session) {
	frontEndFqdn := str.Concat(suiteName, ".", projectFqdn)
	s3.MakeBucket(frontEndFqdn, access, region, awsSession)
	time.Sleep(time.Second * 1)
	tagKey := "pac-project-name"
	s3.TagBucket(frontEndFqdn, tagKey, projectName, awsSession)
	bucketFqdn := str.Concat(frontEndFqdn, ".s3-website-", region, ".amazonaws.com")
	logger.Info(str.Concat("The ", suiteName, " suite will be available after running the front-end Jenkins job at ", bucketFqdn))
	logger.Info(str.Concat("The ", suiteName, " suite will be available in approx 15 minutes at ", frontEndFqdn))
}
