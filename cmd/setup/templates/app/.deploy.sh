#!/bin/bash
aws s3 sync dist s3://integration.{{.projectName}}.pac.pyramidchallenges.com --acl public-read