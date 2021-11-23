#!/usr/bin/env sh
set -e

mkdir .out
rm -f .out/*
GOOS=darwin GOARCH=amd64 go build -o .out/decrypter-$RELEASE_VERSION-darwin-amd64 cmd/decrypter/main.go
GOOS=linux GOARCH=amd64 go build -o .out/decrypter-$RELEASE_VERSION-linux-amd64 cmd/decrypter/main.go
GOOS=linux GOARCH=arm go build -o .out/decrypter-$RELEASE_VERSION-linux-arm cmd/decrypter/main.go
GOOS=windows GOARCH=amd64 go build -o .out/decrypter-$RELEASE_VERSION-windows-amd64 cmd/decrypter/main.go
