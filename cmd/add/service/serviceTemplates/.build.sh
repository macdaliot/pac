#! /bin/bash
npm i
#clean up old stuff
rm -rf dist/*
rm function.zip
#build
npx tsc -p tsconfig-build.json
#get rid of dev deps
#npm prune --production
#package
zip -r function dist/* lambda.js node_modules >> /dev/null