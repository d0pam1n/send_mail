#!/usr/bin/env bash

package='send_mail'
platforms=("windows/amd64" "windows/386")

for platform in "${platforms[@]}"
do
	platform_split=(${platform//\// })
	GOOS=${platform_split[0]}
	GOARCH=${platform_split[1]}
	output_name='./assembly/'$package'-'$GOOS'-'$GOARCH
	if [ $GOOS = "windows" ]; then
		output_name+='.exe'
	fi

	env GOOS=$GOOS GOARCH=$GOARCH go build -o $output_name
	if [ $? -ne 0 ]; then
		echo 'An error has occurred! Aborting the script execution...'
		exit 1
	fi
done
