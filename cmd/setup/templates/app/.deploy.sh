#!/bin/bash
bucket_found=$(aws s3 ls --region us-east-2 | grep integration.{{.projectName}}.pac.pyramidchallenges.com)
if [ ${#bucket_found} -eq 0 ]; then
    aws s3 mb s3://integration.{{.projectName}}.pac.pyramidchallenges.com --region us-east-2
    aws s3 website s3://integration.{{.projectName}}.pac.pyramidchallenges.com --index-document index.html --error-document index.html
fi
aws s3 sync dist s3://integration.{{.projectName}}.pac.pyramidchallenges.com --acl public-read
DISTRO_ID=$(aws cloudfront list-distributions --query "DistributionList.Items[].{Id: Id, OriginDomainName: Origins.Items[0].DomainName}[?contains(OriginDomainName, 'integration.{{.projectName}}.pac.pyramidchallenges.com.s3.amazonaws.com')] | [0]" | jq '.Id')
DISTRO_ID=$(sed -e 's/^"//g' -e 's/"$//g' <<< $(echo "$DISTRO_ID"))
aws cloudfront create-invalidation --distribution-id $DISTRO_ID --paths /bundle.js /styles.css