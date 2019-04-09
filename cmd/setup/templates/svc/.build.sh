#! /bin/bash

DIRECTORIES_TO_IGNORE=("." "./terraform")

DIRECTORIES=$(find . -maxdepth 1 -type d)
for DIRECTORY in ${DIRECTORIES}; do
  if [[ ! "${DIRECTORIES_TO_IGNORE[@]}" =~ "${DIRECTORY}" ]]; then
    pushd $(basename ${DIRECTORY})
      npm i
      rm -rf dist/*
      rm function.zip
      npx tsc -p tsconfig-build.json
      zip -r function dist/* lambda.js node_modules >> /dev/null
    popd
  fi
done
