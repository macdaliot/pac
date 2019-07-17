package setup

import (
	"github.com/PyramidSystemsInc/go/aws"
	"github.com/PyramidSystemsInc/go/aws/ec2"
	"github.com/PyramidSystemsInc/pac/config"
)

func FindNextAvailableVpcCidrBlock() string {
	vpcCidrBlocksNeeded := 1
	awsRegion := config.Get("region")
	awsSession := aws.CreateAwsSession(awsRegion)
	return ec2.FindAvailableVpcCidrBlocks(vpcCidrBlocksNeeded, awsSession)[0]
}
