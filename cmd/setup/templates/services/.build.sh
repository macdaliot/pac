#! /bin/bash

DIRECTORIES_TO_IGNORE=("." "./terraform")

DIRECTORIES=$(find . -maxdepth 1 -type d)
for DIRECTORY in ${DIRECTORIES}; do
  if [[ ! "${DIRECTORIES_TO_IGNORE[@]}" =~ "${DIRECTORY}" ]]; then
    pushd $(basename ${DIRECTORY})
      npm i
      rm -rf dist/*
      rm function.zip
      npm run generate:templates:cloud
      npx tsc || return 2
      npm prune --production
      zip -r function dist/* lambda.js package.json node_modules tsoa-cloud.json >> /dev/null
      rm -Rf dist
    popd
  fi
done
