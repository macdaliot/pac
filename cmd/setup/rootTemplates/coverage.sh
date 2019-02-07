#!/bin/sh
TARGET=$1

echo $TARGET
cd $TARGET
ls -la
npm i
npm run-script test-coverage