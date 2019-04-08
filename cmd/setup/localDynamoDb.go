package setup

import (
	"github.com/PyramidSystemsInc/go/aws"
	"github.com/PyramidSystemsInc/go/aws/dynamodb"
	"github.com/PyramidSystemsInc/go/commands"
	"github.com/PyramidSystemsInc/go/logger"
	"github.com/PyramidSystemsInc/pac/config"
	awss "github.com/aws/aws-sdk-go/aws"
	ddb "github.com/aws/aws-sdk-go/service/dynamodb"
)

func LocalDynamoDb() {
	commands.Run("docker run --name pac-db-local -d -p 8001:8000 amazon/dynamodb-local -jar DynamoDBLocal.jar -sharedDb", "")
	logger.Info("Started local DynamoDB running in Docker")
}

// createDynamoDBTable creates a dynamodb table for the Terraform S3 backend to use for locking to prevent simultaneous acces to state
func CreateLockingTable() {
	awsSession := aws.CreateAwsSession("us-east-2")
	name := "pac-" + config.Get("projectName") + "-terraform-locking-table"

	input := &ddb.CreateTableInput{
		AttributeDefinitions: []*ddb.AttributeDefinition{
			{
				AttributeName: awss.String("LockID"),
				AttributeType: awss.String("S"),
			},
		},
		KeySchema: []*ddb.KeySchemaElement{
			{
				AttributeName: awss.String("LockID"),
				KeyType:       awss.String("HASH"),
			},
		},
		ProvisionedThroughput: &ddb.ProvisionedThroughput{
			ReadCapacityUnits:  awss.Int64(5),
			WriteCapacityUnits: awss.Int64(5),
		},
		TableName: awss.String(name),
	}

	dynamodb.CreateTable(input, awsSession)
}
