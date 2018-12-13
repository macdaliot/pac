package create

import (
  "github.com/PyramidSystemsInc/go/commands"
  "github.com/PyramidSystemsInc/go/logger"
)

func DynamoDb(projectName string) {
  commands.Run("docker run --name pac-" + projectName + "-db-local -p 8000:8000 -d amazon/dynamodb-local -jar DynamoDBLocal.jar -sharedDb", "")
  logger.Info("Started local DynamoDB running in Docker")
}
