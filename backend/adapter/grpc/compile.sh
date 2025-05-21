#!/bin/bash

# Exit the script if any command fails
set -e

# Define the directory containing your proto files
PROTO_DIR="./proto"

# Define the output directory for the compiled Go files
OUTPUT_DIR="."

# Ensure Buf is installed
if ! [ -x "$(command -v buf)" ]; then
  echo "Error: Buf is not installed. Install it from https://docs.buf.build/installation."
  exit 1
fi

# Compile the proto files using Buf
echo "Compiling proto files with Buf..."
buf generate $PROTO_DIR --output $OUTPUT_DIR

echo "Compilation complete. Files generated in $OUTPUT_DIR."

