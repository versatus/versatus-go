# versatus-go

## Overview

This package contains various structures and helper functions for building smart contracts for the Versatus network using Go. We recommend leveraging the general smart contract development guide on Versatus (https://github.com/versatus/versatus/blob/main/docs/DevelopingSmartContracts.md) while diving into this Go-specific documentation.

An example smart contract can be found in `cmd/contract/main.go`. Please note that the interface specifications are still evolving.

## Go Installation
Go development requires the Go programming language. To install Go, follow the instructions on the official Go website (https://golang.org/dl/).

## Dependencies

Go is a statically typed, compiled language known for its simplicity and efficiency, making it suitable for smart contract development. Go natively supports Web Assembly (WASM), making it ideal for compiling smart contracts. Ensure you have the latest version installed.

## TinyGo
TinyGo is a Go compiler intended for use in small places such as microcontrollers, WebAssembly (WASM), and command-line tools. It can be used to compile Go code to a more compact form of WASM, suitable for smart contracts. Installation instructions can be found on the TinyGo website (https://tinygo.org/getting-started/) and below.

## Installing Toolchain

In order to run the aforementioned example on the Versatus network, it needs to be compiled to Web Assembly with WASI support. This is most-easily done using the TinyGo embedded Go compiler, but other options may work as long as there is a `_start` symbol and the module is statically linked.

Installing TinyGo may be as simple as an `apt install tinygo` on many Linux distributions, or using something like `brew tap tinygo-org/tools && brew install tinygo` on MacOS, or by otherwise installing one of the releases from the [TinyGo Releases Page](https://github.com/tinygo-org/tinygo/releases).

## Creating a Contract

In order to build a smart contract for the Versatus network, simply include the `versatus-go/contract` package:

```go
import "github.com/versatus/versatus-go/contract"
```

The `contract` object will then allow you to parse the contract input (received on stdin) to Go data structures, and to send outputs on stdout. For example:

```go
_, inputs := contract.ParseInput()
```

and:

```go
var outputs contract.ComputeOutputs
_ = contract.SendOutput(outputs)
```

## Compiling a Contract

To compile with TinyGo to produce a WASM file, simply use the TinyGo compiler command instead of the standard Go compiler command, and specify the target as WASI:

```
tinygo build --target=wasi \
    -scheduler=none -gc=conservative \
    -o contract.wasm ./...
```

This should result in a `contract.wasm` file that is the Web Assembly smart contract.


##Test Your Contract
Testing your contract can be done using Go's built-in testing framework. Leverage Go's testing framework to write tests and execute them using `go test ./...`

## Deploying a Contract to the Versatus Network
After testing, deploy your contract to the Versatus network following the deployment procedures outlined in Versatus's documentation.
TBA

