#!/bin/bash

platforms=("windows/amd64" "windows/386" "darwin/amd64" "darwin/386" "linux/amd64" "linux/386" "linux/arm64" "linux/arm")
buildDirectory=./build
echo "Creating executables for platforms ${platforms[*]}"

echo 'Purging previous build directory'
rm ${buildDirectory}/*

for platform in "${platforms[@]}"
do
    splitPlatform=(${platform//\// })
    GOOS=${splitPlatform[0]}
    GOARCH=${splitPlatform[1]}
    executableName='BeeMail-'$GOOS'-'$GOARCH
    if [ $GOOS = "windows" ]; then
        executableName+='.exe'
    fi  
    echo "Creating executable for $platform"

    env GOOS=$GOOS GOARCH=$GOARCH go build -o ${buildDirectory}/${executableName} 
    if [ $? -ne 0 ]; then
        >&2 echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi
done

echo 'Successfully created executables'
