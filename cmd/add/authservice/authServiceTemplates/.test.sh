#! /bin/bash

npm run test
if [[ $(echo $?) -eq 0 ]]; then
  echo "All unit tests passed"
else
  echo "Stopping pipeline due to unit test failures"
  exit 2
fi