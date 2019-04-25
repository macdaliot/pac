#! /bin/bash

DIRECTORIES_TO_IGNORE=("." "./terraform")

DIRECTORIES=$(find . -maxdepth 1 -type d)
for DIRECTORY in ${DIRECTORIES}; do
  if [[ ! "${DIRECTORIES_TO_IGNORE[@]}" =~ "${DIRECTORY}" ]]; then
    pushd $(basename ${DIRECTORY})
      npm i
      rm -rf dist/*
      rm function.zip
      cp -R ../../core .
      cp -R ../../domain .
      npm run generate:routes
      npm run generate:swagger
      npx tsc -p tsconfig-build.json || return 2
      zip -r function dist/* lambda.js node_modules >> /dev/null
      rm -Rf core dist docs domain
    popd
  fi
done
