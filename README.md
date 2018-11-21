# MATRIX Protos Go
![](https://i.imgur.com/gZkI9Ap.png)

Golang Protocol Buffers for the MATRIX CORE Protocol

# Installation
```
go get github.com/matrix-io/matrix-protos-go
```

# Compiling The Protocol Buffers (Optional)
Below are the instructions on how to generate our [MATRIX protocol buffers](https://github.com/matrix-io/protocol-buffers), if you want to build for a version not currently supported.

## Prerequsits
- Download & build the latest version of **protobuf-all**: https://github.com/protocolbuffers/protobuf/releases
- Install Golang: https://golang.org/doc/install
- Get the protoc compiling tool: `go get -u github.com/golang/protobuf/protoc-gen-go`
- Get the goprotowrap compiling tool from Sqaure: `go get -u github.com/square/goprotowrap/cmd/protowrap`
- Download & unzip the MATRIX-io protocol buffers onto your desktop: https://github.com/matrix-io/protocol-buffers

## Compiling Command
In your terminal, `cd` into your desktop and run the following command:
```
protowrap --go_out=. -I protocol-buffers protocol-buffers/matrix_io/**/*.proto protocol-buffers/matrix_io/**/**/*.proto
```
This will generate a folder named `github.com` with the Golang protocol buffers inside.
