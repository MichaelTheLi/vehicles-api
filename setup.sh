#!/bin/bash

# Exit on error
set -e

echo "Setting up Go environment..."

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "Error: Go is not installed. Please install Go from https://go.dev/dl/"
    exit 1
fi

# Set Go module proxy settings
export GOPROXY=https://proxy.golang.org,direct
export GOSUMDB=sum.golang.org

# Set up environment variables
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

echo "Verifying Go installation..."
go version
echo "Current Go environment:"
go env

echo "Downloading dependencies..."
if ! go mod tidy; then
    echo "Error: Failed to download dependencies"
    exit 1
fi

echo "Installing gqlgen..."
if ! go install github.com/99designs/gqlgen@v0.17.73; then
    echo "Error: Failed to install gqlgen"
    exit 1
fi

echo "Generating GraphQL code..."
if ! go run github.com/99designs/gqlgen generate; then
    echo "Error: Failed to generate GraphQL code"
    exit 1
fi

echo "Setup completed successfully!" 
