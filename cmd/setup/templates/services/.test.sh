#! /bin/bash

DIRECTORIES_TO_IGNORE=("." "./terraform")

DIRECTORIES=$(find . -maxdepth 1 -type d)
for DIRECTORY in ${DIRECTORIES}; do
  if [[ ! "${DIRECTORIES_TO_IGNORE[@]}" =~ "${DIRECTORY}" ]]; then
    pushd $(basename ${DIRECTORY})
      rm -Rf node_modules
      npm i
      echo "Running tests in $(basename ${DIRECTORY})"
      EXIT_CODE=$(npm run test-ci)
      if [ $(echo $?) -ne 0 ]; then
        exit 2
      fi
    popd
  fi
done
