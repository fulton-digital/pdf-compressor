#!/bin/sh

mkdir -p bin/darwin/amd64
mkdir -p bin/darwin/386
mkdir -p bin/windows/amd64
mkdir -p bin/windows/386
mkdir -p bin/linux/amd64
mkdir -p bin/linux/386
mkdir -p bin/linux/arm64

GOOS=darwin GOARCH=amd64 go build -o bin/darwin/amd64/pdfcompress pdf-compressor.go
GOOS=darwin GOARCH=386 go build -o bin/darwin/386/pdfcompress pdf-compressor.go
GOOS=windows GOARCH=amd64 go build -o bin/windows/amd64/pdfcompress.exe pdf-compressor.go
GOOS=windows GOARCH=386 go build -o bin/windows/386/pdfcompress.exe pdf-compressor.go
GOOS=linux GOARCH=amd64 go build -o bin/linux/amd64/pdfcompress pdf-compressor.go
GOOS=linux GOARCH=arm64 go build -o bin/linux/arm64/pdfcompress pdf-compressor.go
GOOS=linux GOARCH=386 go build -o bin/linux/386/pdfcompress pdf-compressor.go