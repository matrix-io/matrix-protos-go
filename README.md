# MATRIX Protos Go
![](https://i.imgur.com/gZkI9Ap.png)

Golang Protocol Buffers for the MATRIX CORE Protocol

# Compiling Protocol Buffers

## Prerequsits
- Download & build the latest version of protobuf-all: https://github.com/protocolbuffers/protobuf/releases
- Install Golang: https://golang.org/doc/install
- Get the protoc compiling tool: `go get -u github.com/golang/protobuf/protoc-gen-go`
- Get the compiling tool from Sqaure: `go get -u github.com/square/goprotowrap/cmd/protowrap`
- Download & unzip the MATRIX-io protocol buffers to your desktop: https://github.com/matrix-io/protocol-buffers

## Compiling Command
In your terminal, move to your desktop and run the following command. This will generate a directory named `github.com` with the Golang protocol buffers inside.

`protowrap --go_out=. -I protocol-buffers protocol-buffers/matrix_io/**/*.proto protocol-buffers/matrix_io/**/**/*.proto`
