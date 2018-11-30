# MATRIX Protos Go
![](https://i.imgur.com/gZkI9Ap.png)

Golang Protocol Buffers for the MATRIX CORE Protocol

The MATRIX CORE protocol is an abstraction layer built on top of ZeroMQ that uses protocol buffers as the data exchange format. In a nutshell, CORE exposes 4 ports for every service, sensor, and component: 

**Base Port**: Send configurations (LED colors, sensor refresh rate, etc..).

**Error Port**: Receive any incoming errors.

**Keep Alive Port**: Send keep alive pings to keep receiving data.

**Data Update Port**: Receive relevant port data (Sensor data, LED count, etc..).

In the following examples, we will show how to send a config to the driver port and how to obtain data updates on the data update port.

# Installation
For guided installation instructions on MATRIX devices, Visit the [MATRIX CORE installation page](https://matrix-io.github.io/matrix-documentation/matrix-core/getting-started/).
```
go get github.com/matrix-io/matrix-protos-go
```

# Docs & Community
[MATRIX Documentation](https://matrix-io.github.io/matrix-documentation/): Read and learn about the MATRIX Ecosystem.

[MATRIX Community](https://community.matrix.one/): Post any questions or projects in our community form.

# Examples

<details close>
<summary>LED Example</summary>

### Initial Imports & Variables
```go
package main

import (
	"time"

	"github.com/golang/protobuf/proto"
	core "github.com/matrix-io/matrix-protos-go/matrix_io/malos/v1"
	zmq "github.com/pebbe/zmq4"
)

// Everloop data struct
var everloop = core.EverloopImage{}
// 35 = MATRIX Creator, 18 = MATRIX Voice
var ledCount = 35                  
```

### Base Port
For simplicity, the Everloop example will hard code the number of LEDs.

```go
func main() {
	// Connect ZMQ socket to MATRIX CORE
	pusher, _ := zmq.NewSocket(zmq.PUSH)    // Create a pusher socket
	pusher.Connect("tcp://127.0.0.1:20021") // Connect pusher to data update port

	// Set each LED color
	for i := 0; i < ledCount; i++ {
		led := core.LedValue{
			Red:   1,
			Blue:  1,
			Green: 0,
			White: 0,
		}
		// Add LED to everloop
		everloop.Led = append(everloop.Led, &led)
	}

	// Create driver configuration for everloop protocol
	configuration := core.DriverConfig{
		Image: &everloop,
	}

	// Encode protocol buffer
	var encodedConfiguration, _ = proto.Marshal(&configuration)
	// Send protocol buffer
	pusher.Send(string((encodedConfiguration)), 1)
	// Give pusher time to connect
	time.Sleep(1 * time.Millisecond)
}
```
</details>

# Compiling The Protocol Buffers (Optional)
Below are the instructions on how to generate our [MATRIX protocol buffers](https://github.com/matrix-io/protocol-buffers), if you want to build for a version not currently supported.

## Prerequisites
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
