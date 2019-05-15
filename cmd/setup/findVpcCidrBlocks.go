package setup

import (
	"github.com/PyramidSystemsInc/go/aws"
	"github.com/PyramidSystemsInc/go/aws/ec2"
	"github.com/PyramidSystemsInc/pac/config"
)

func FindAvailableVpcCidrBlocks() []string {
	vpcCidrBlocksNeeded := 2
	awsRegion := config.Get("region")
	awsSession := aws.CreateAwsSession(awsRegion)
	return ec2.FindAvailableVpcCidrBlocks(vpcCidrBlocksNeeded, awsSession)
}
