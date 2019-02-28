package setup

import (
  "github.com/PyramidSystemsInc/go/commands"
  "github.com/PyramidSystemsInc/go/logger"
)

func LocalDynamoDb() {
  commands.Run("docker run --name pac-db-local -d -p 8001:8001 amazon/dynamodb-local -jar DynamoDBLocal.jar -sharedDb", "")
  logger.Info("Started local DynamoDB running in Docker")
}
