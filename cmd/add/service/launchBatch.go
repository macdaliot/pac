package service

import (
  "runtime"
  "github.com/PyramidSystemsInc/go/files"
  "github.com/PyramidSystemsInc/go/str"
)

// CreateLaunchSh creates a templated shell script to launch the service
func CreateLaunchSh(filePath string, config map[string]string) {
  const template = `@if "%DEBUG%" == "" @echo off
@rem ##########################################################################
@rem
@rem  startup script for Windows
@rem
@rem ##########################################################################

@rem Set local scope for the variables with windows NT shell
if "%OS%"=="Windows_NT" setlocal


set AWS_ACCESS_KEY_ID=aws --profile default configure get aws_access_key_id
if not [%AWS_ACCESS_KEY_ID%] == [] goto AWSCheck
if not [%AWS_SECRET_ACCESS_KEY%] == [] goto AWSCheck

:AWSCheck
aws --profile default configure get aws_access_key_id > A1
set /p AWS_ACCESS_KEY_ID=<A1
aws --profile default configure get aws_secret_access_key > A1
set /p AWS_SECRET_ACCESS_KEY=<A1
del /q A1

if not exist node_modules npm i

docker stop pac-{{.serviceName}}-service  || goto :error
docker stop pac-{{.serviceName}}-service  || goto :error
docker rm pac-{{.serviceName}}-service  || goto :error
  npx tsc server.ts  || goto :error
docker build -t pac-{{.serviceName}}-service .  || goto :error
docker run --name pac-{{.serviceName}}-service -p 3000:3000 --link pac-db-local -e AWS_ACCESS_KEY_ID=%AWS_ACCESS_KEY_ID% -e AWS_SECRET_ACCESS_KEY=%AWS_SECRET_ACCESS_KEY% -d pac-{{.serviceName}}-service  || goto :error

echo "DONE! Launched microservice locally"

:error
echo ABORTED LAUNCH: Failed with error #%errorlevel%.
exit /b %errorlevel%
`
  files.CreateFromTemplate(filePath, template, config)
  if runtime.GOOS == "windows" {
    files.ChangePermissions(str.Concat(".\\", config["serviceName"], "\\launch.sh"), 0755)
  } else {
    files.ChangePermissions(str.Concat("./", config["serviceName"], "/launch.sh"), 0755)
  }
}
