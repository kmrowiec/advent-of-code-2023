#!/bin/bash

# Check if an argument is provided
if [ $# -eq 0 ]; then
    echo "Usage: $0 <integer>"
    exit 1
fi

destination_folder="day${1}"
source_folder_path="dayX"
destination_folder_path="./${destination_folder}"

mkdir ./$destination_folder
cp -r $source_folder_path/* $destination_folder_path
mv "$destination_folder_path/dayX.go.template" "$destination_folder_path/day${1}.go"
sed -i '' -e  "s/dayX/day$1/g" $destination_folder/day${1}.go
