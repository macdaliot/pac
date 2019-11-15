// This will eventually be `pac deploy`

package main

import (
  "fmt"
  "os"
  "github.com/PyramidSystemsInc/go/aws/sts"
  "github.com/PyramidSystemsInc/go/commands"
)

func main() {
  setAwsAccountEnvironmentVariables()
  setHardCodedEnvironmentVariables()
  createTerraformPlan()
}

func setAwsAccountEnvironmentVariables() {
  accessKeyId, _ := commands.Run("aws configure get aws_access_key_id", "")
  os.Setenv("TF_VAR_aws_access_key_id", accessKeyId)
  accountNumber := sts.GetAccountID("default")
  os.Setenv("TF_VAR_aws_account_number", accountNumber)
  secretAccessKey, _ := commands.Run("aws configure get aws_secret_access_key", "")
  os.Setenv("TF_VAR_aws_secret_access_key", secretAccessKey)
  psiAccessKeyId, _ := commands.Run("aws configure get aws_access_key_id --profile psi", "")
  os.Setenv("TF_VAR_psi_aws_access_key_id", psiAccessKeyId)
  psiAccountNumber := sts.GetAccountID("psi")
  os.Setenv("TF_VAR_psi_aws_account_number", psiAccountNumber)
  psiSecretAccessKey, _ := commands.Run("aws configure get aws_secret_access_key --profile psi", "")
  os.Setenv("TF_VAR_psi_aws_secret_access_key", psiSecretAccessKey)
  for _, variable := range os.Environ() {
    fmt.Println(variable)
  }
}

func setHardCodedEnvironmentVariables() {
  os.Setenv("TF_VAR_github_auth", "jdiederiks%40psi-it.com:Diedre%5E2018")
  os.Setenv("TF_VAR_github_password", "Diedre^2018")
  os.Setenv("TF_VAR_github_username", "jdiederiks@psi-it.com")
  os.Setenv("TF_VAR_iam_user_password", "psystems")
  os.Setenv("TF_VAR_jenkins_password", "systems")
  os.Setenv("TF_VAR_jenkins_username", "pyramid")
  os.Setenv("TF_VAR_jwt_issuer", "urn:pacAuth")
  os.Setenv("TF_VAR_jwt_secret", "4pWQUrx6RkgU6o2TC")
  os.Setenv("TF_VAR_postgres_password", "pyramid")
  os.Setenv("TF_VAR_sonar_jdbc_password", "sonar")
  os.Setenv("TF_VAR_sonar_jdbc_username", "sonar")
  os.Setenv("TF_VAR_sonar_password", "systems")
  os.Setenv("TF_VAR_sonar_secret", "849e67e97176a6ddf594d102c86d93d085c8d222")
  os.Setenv("TF_VAR_sonar_username", "pyramid")
}

func createTerraformPlan() {
  out, err := commands.Run("terraform plan -out=tfplan", "")
  if err != nil {
    fmt.Println(err.Error())
  }
  fmt.Println(out)
}
