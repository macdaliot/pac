package cleanCloud

import (
	"time"

	"github.com/PyramidSystemsInc/go/aws"
	"github.com/PyramidSystemsInc/go/aws/cloudfront"
	"github.com/PyramidSystemsInc/go/aws/resourcegroups"
	"github.com/PyramidSystemsInc/go/aws/route53"
	"github.com/PyramidSystemsInc/go/commands"
	"github.com/PyramidSystemsInc/go/logger"
	"github.com/PyramidSystemsInc/go/str"
	"github.com/PyramidSystemsInc/pac/cmd/setup"
	"github.com/PyramidSystemsInc/pac/config"
)

func DeleteAllResources() {
	start := time.Now()
	//destroy AWS resources managed by Terraform
	setup.TerraformDestroy()
	projectName := config.Get("projectName")
	region := "us-east-2"
	awsSession := aws.CreateAwsSession(region)
	parentHostedZone := "pac.pyramidchallenges.com"
	fqdn := str.Concat(projectName, ".pac.pyramidchallenges.com")
	integrationFqdn := str.Concat("integration.", fqdn)
	demoFqdn := str.Concat("demo.", fqdn)
	aliases := str.Concat(integrationFqdn, " ", demoFqdn)
	go deleteDistributions(aliases)
	resourcegroups.Create(projectName, "pac-project-name", projectName, awsSession)
	resourcegroups.DeleteAllResources(projectName, awsSession)
	logger.Info(str.Concat("Finished deleting all AWS resources tagged as part of PAC project ", projectName))
	// AWS Resource Groups does not have the ability to find certain types of
	// resources even if they are properly tagged. This is why Route53 and
	// Cloudfront are handled separately below
	route53.DeleteHostedZone(fqdn, awsSession)
	route53.DeleteRecord(parentHostedZone, fqdn, awsSession)
	logger.Info("Finished deleting Route53 resources")
	cloudfront.DisableDistribution(integrationFqdn, awsSession)
	cloudfront.DisableDistribution(demoFqdn, awsSession)
	logger.Info("Finished disabling the Cloudfront distributions")
	logger.Warn("    Cloudfront distributions have been disabled, not deleted")
	logger.Warn("    There is a Cloudfront origin access identity that PAC does not yet support deleting")
}

func deleteDistributions(aliases string) {
	commands.Run("go get -u github.com/PyramidSystemsInc/delete-cloudfront-resources", "")
	commands.Run(str.Concat("go run github.com/PyramidSystemsInc/delete-cloudfront-resources -- ", aliases), "")
}
