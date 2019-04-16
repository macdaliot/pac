package setup

import (
  "strconv"

  "github.com/PyramidSystemsInc/go/aws"
  "github.com/PyramidSystemsInc/go/aws/ec2"
  "github.com/PyramidSystemsInc/go/errors"
  "github.com/PyramidSystemsInc/go/logger"
  "github.com/PyramidSystemsInc/go/terraform"
)

func Infrastructure() {
  awsRegion := "us-east-2"
  awsSession := aws.CreateAwsSession(awsRegion)
  usedVpcCidrBlocks := ec2.GetAllVpcCidrBlocks(awsSession)
  freeVpcCidrBlocks := findFirstAvailableVpcCidrBlocks(usedVpcCidrBlocks, 2)
  terraformDirectory := "terraform"
  initOutput := terraform.Initialize(terraformDirectory)
  logger.Info(initOutput)
  TerraformPlan(freeVpcCidrBlocks)
  TerraformApply()
}

func findFirstAvailableVpcCidrBlocks(usedCidrBlocks []string, numberToFind int) []string {
  var freeVpcCidrBlocks []string
  var secondPartDigits []string
  for i := 0; i < numberToFind; i++ {
    cidrBlockError := "The following error occurred while attempting to find a free CIDR block for a VPC: "
    if i == 0 {
      secondPartDigits = append(secondPartDigits, "1")
    } else {
      lastValue, err := strconv.Atoi(secondPartDigits[i - 1])
      if err != nil {
        errors.LogAndQuit(cidrBlockError + err.Error())
      }
      secondPartDigits = append(secondPartDigits, strconv.Itoa(lastValue + 1))
    }
    digitFound := true
    for digitFound {
      digitFound = false
      out: for _, usedCidrBlock := range usedCidrBlocks {
        testCidrBlock := "10."+secondPartDigits[i]+".0.0/16"
        if usedCidrBlock == testCidrBlock {
          numberDigit, err := strconv.Atoi(secondPartDigits[i])
          if err != nil {
            errors.LogAndQuit(cidrBlockError + err.Error())
          }
          numberDigit++
          secondPartDigits[i] = strconv.Itoa(numberDigit)
          digitFound = true
          break out
        }
      }
    }
    freeVpcCidrBlocks = append(freeVpcCidrBlocks, "10."+secondPartDigits[i]+".0.0/16")
  }
  return freeVpcCidrBlocks
}
