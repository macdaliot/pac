@if "%DEBUG%" == "" @echo off
@rem ##########################################################################
@rem
@rem  startup script for Windows
@rem
@rem ##########################################################################

@rem Set local scope for the variables with windows NT shell
if "%OS%"=="Windows_NT" setlocal


set AWS_ACCESS_KEY_ID=aws
if not [%AWS_ACCESS_KEY_ID%] == [] goto AWSCheck
if not [%AWS_SECRET_ACCESS_KEY%] == [] goto AWSCheck

:AWSCheck
for /f "delims=" %%i in ('aws configure get aws_access_key_id') do set AWS_ACCESS_KEY_ID=%%i
for /f "delims=" %%u in ('aws configure get aws_secret_access_key') do set AWS_SECRET_ACCESS_KEY=%%u


if not exist node_modules npm i

docker stop pac-auth-service  || goto :error
docker stop pac-auth-service  || goto :error
docker rm pac-auth-service  || goto :error
  npx tsc server.ts  || goto :error
docker build -t pac-auth-service .  || goto :error
docker run --name pac-auth-service -p 3000:3000 --link pac-db-local -e AWS_ACCESS_KEY_ID=%AWS_ACCESS_KEY_ID% -e AWS_SECRET_ACCESS_KEY=%AWS_SECRET_ACCESS_KEY% -d pac-auth-service  || goto :error

echo "DONE! Launched microservice locally"

:error
echo ABORTED LAUNCH: Failed with error #%errorlevel%.
exit /b %errorlevel%
