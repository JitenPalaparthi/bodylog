#!/bin/sh

cd covidAnalysis

git --version 2>&1 >/dev/null
GIT_IS_AVAILABLE=$?
if [ $GIT_IS_AVAILABLE -eq 0 ]; then 
git init
else
echo "git is not installled.Install git from the following url https://git-scm.com/book/en/v2/Getting-Started-Installing-Git"
fi

## generate protos for all models

 go mod init medicalResearch

 go mod tidy

 go build ./...