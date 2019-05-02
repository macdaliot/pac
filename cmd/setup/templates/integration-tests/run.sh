#! /bin/bash

# If the gradle-wrapper.jar does not exist, create it
if [ ! -f gradle/wrapper/gradle-wrapper.jar ]; then
  pushd gradle/wrapper; jar cf gradle-wrapper.jar *; popd
fi

chmod 755 gradlew
./gradlew clean test aggregate
