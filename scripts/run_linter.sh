#!/bin/bash

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

FOLDER_PATH=".."

golangci-lint run "$SCRIPT_DIR/$FOLDER_PATH/..."

read -p "Do you want to run the fix? (y/n) " run_fix

if [[ $run_fix == "y" ]]; then
    golangci-lint run --fix "$SCRIPT_DIR/$FOLDER_PATH/..."
fi