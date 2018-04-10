#!/bin/sh

mkdir -p $PWD/bin/darwin/amd64
mkdir -p $PWD/bin/darwin/386
mkdir -p $PWD/bin/windows/amd64
mkdir -p $PWD/bin/windows/386
mkdir -p $PWD/bin/linux/amd64
mkdir -p $PWD/bin/linux/386
mkdir -p $PWD/bin/linux/arm64

GOOS=darwin GOARCH=amd64 go build -o $PWD/bin/darwin/amd64/pdfcompress $PWD/pdf-compressor.go
GOOS=darwin GOARCH=386 go build -o $PWD/bin/darwin/386/pdfcompress $PWD/pdf-compressor.go
GOOS=windows GOARCH=amd64 go build -o $PWD/bin/windows/amd64/pdfcompress.exe $PWD/pdf-compressor.go
GOOS=windows GOARCH=386 go build -o $PWD/bin/windows/386/pdfcompress.exe $PWD/pdf-compressor.go
GOOS=linux GOARCH=amd64 go build -o $PWD/bin/linux/amd64/pdfcompress $PWD/pdf-compressor.go
GOOS=linux GOARCH=arm64 go build -o $PWD/bin/linux/arm64/pdfcompress $PWD/pdf-compressor.go
GOOS=linux GOARCH=386 go build -o $PWD/bin/linux/386/pdfcompress $PWD/pdf-compressor.go