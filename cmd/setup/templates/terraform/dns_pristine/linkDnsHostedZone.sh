#!/bin/bash

echo "@Link DNS Hosted Zones"

# call to get output
#TODO NEED TO CD INTO DNS_PRISTINE DIR
TF_OUTPUT_NS=$(terraform output ns_records)
CLEAN_NS=$(echo $TF_OUTPUT_NS|sed 's/\(.*\),/\1/')

curl -X PUT -H "Content-Type: application/json" -d "{\"recordName\":\"${TF_VAR_project_name}.pac.pyramidchallenges.com\", \"recordValues\":$CLEAN_NS}" https://bdso.aws.pyramidchallenges.com
echo "Successfully updated DNS records to support certificate validation."
