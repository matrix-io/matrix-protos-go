# MATRIX Protos Go
![](https://i.imgur.com/gZkI9Ap.png)

Golang Protocol Buffers for the MATRIX CORE Protocol

# Installation
```
go get github.com/matrix-io/matrix-protos-go
```

# Compiling The Protocol Buffers (Optional)
The instructions below will show you how to compile `matrix-protos-go`.

## Prerequsits
- Download & build the latest version of protobuf-all: https://github.com/protocolbuffers/protobuf/releases
- Install Golang: https://golang.org/doc/install
- Get the protoc compiling tool: `go get -u github.com/golang/protobuf/protoc-gen-go`
- Get the compiling tool from Sqaure: `go get -u github.com/square/goprotowrap/cmd/protowrap`
- Download & unzip the MATRIX-io protocol buffers to your desktop: https://github.com/matrix-io/protocol-buffers

## Compiling Command
**⚠️Awaiting matrix-io/protocol-buffers 
<a href="https://github.com/matrix-io/protocol-buffers/pull/92">Pull Request #92</a> 
Approval⚠️**

In your terminal, move into your desktop and run the following command:
```
protowrap --go_out=. -I protocol-buffers protocol-buffers/matrix_io/**/*.proto protocol-buffers/matrix_io/**/**/*.proto
```
This will generate a folder named `github.com` with the Golang protocol buffers inside.
