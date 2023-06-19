#!/usr/bin/env sh
set -e

mkdir -p .out
rm -f .out/*
GOOS=darwin GOARCH=amd64 go build -o .out/decrypter-$RELEASE_VERSION-darwin-amd64 cmd/decrypter/main.go
GOOS=darwin GOARCH=arm64 go build -o .out/decrypter-$RELEASE_VERSION-darwin-arm64 cmd/decrypter/main.go
GOOS=linux GOARCH=amd64 go build -o .out/decrypter-$RELEASE_VERSION-linux-amd64 cmd/decrypter/main.go
GOOS=linux GOARCH=arm64 go build -o .out/decrypter-$RELEASE_VERSION-linux-arm64 cmd/decrypter/main.go
GOOS=windows GOARCH=amd64 go build -o .out/decrypter-$RELEASE_VERSION-windows-amd64 cmd/decrypter/main.go
