package setup

import (
	"github.com/PyramidSystemsInc/go/logger"
	"github.com/PyramidSystemsInc/pac/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/aws/aws-sdk-go/service/s3"

	"fmt"
	"os"
)

// CreateEncryptionKey creates a customer managed key in the AWS Key Management Service
func CreateEncryptionKey() {
	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-2")},
	)

	// Create KMS service client
	svc := kms.New(sess)

	// Create the key
	result, err := svc.CreateKey(&kms.CreateKeyInput{
		Tags: []*kms.Tag{
			{
				TagKey:   aws.String("CreatedBy"),
				TagValue: aws.String("ExampleUser"),
			},
		},
	})

	if err != nil {
		fmt.Println("Got error creating key: ", err)
		os.Exit(1)
	}

	config.Set("encryptionKeyId", *result.KeyMetadata.KeyId)
}

// EncryptS3Bucket turns on encryption on the S3 bucket
func EncryptS3Bucket() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := s3.New(sess)

	bucket := config.Get("terraformS3Bucket")
	key := config.Get("encryptionKeyId")

	defEnc := &s3.ServerSideEncryptionByDefault{KMSMasterKeyID: aws.String(key), SSEAlgorithm: aws.String(s3.ServerSideEncryptionAwsKms)}
	rule := &s3.ServerSideEncryptionRule{ApplyServerSideEncryptionByDefault: defEnc}
	rules := []*s3.ServerSideEncryptionRule{rule}
	serverConfig := &s3.ServerSideEncryptionConfiguration{Rules: rules}
	input := &s3.PutBucketEncryptionInput{Bucket: aws.String(bucket), ServerSideEncryptionConfiguration: serverConfig}
	_, err := svc.PutBucketEncryption(input)
	if err != nil {
		fmt.Println("Got an error adding default KMS encryption to bucket", bucket)
		fmt.Println(err.Error())
		os.Exit(1)
	}

	logger.Info("Bucket " + bucket + " now has KMS encryption by default")
}
