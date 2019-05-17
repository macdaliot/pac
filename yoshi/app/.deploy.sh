#!/bin/bash
aws s3 sync dist s3://integration.yoshi.pac.pyramidchallenges.com --acl public-read