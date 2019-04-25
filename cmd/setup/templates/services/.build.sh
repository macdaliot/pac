#! /bin/bash

DIRECTORIES_TO_IGNORE=("." "./terraform")

DIRECTORIES=$(find . -maxdepth 1 -type d)
for DIRECTORY in ${DIRECTORIES}; do
  if [[ ! "${DIRECTORIES_TO_IGNORE[@]}" =~ "${DIRECTORY}" ]]; then
    pushd $(basename ${DIRECTORY})
      npm i
      rm -rf dist/*
      rm function.zip
      npm run generate:routes
      npm run generate:swagger
      npx tsc || return 2
      zip -r function dist/* lambda.js package.json node_modules >> /dev/null
      rm -Rf dist docs
    popd
  fi
done
