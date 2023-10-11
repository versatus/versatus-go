# versatus-go

## Overview

This package contains various structures and helper functions for building smart contracts for the Versatus network using Go.

The `cmd/contract/main.go` file contains an example.

Note that none of the inputs and outputs have been fully defined yet, and will change. Very soon.

## Installing Toolchain

In order to run this example on the Versatus network, it needs to be compiled to Web Assembly with WASI support. This is most-easily done using the TinyGo embedded Go compiler, but other options may work as long as there is a `_start` symbol and the module is statically linked.

Installing TinyGo may be as simple as an `apt install tinygo` on many Linux distributions, or using something like `brew tap tinygo-org/tools && brew install tinygo` on MacOS, or by otherwise installing one of the releases from the [TinyGo Releases Page](https://github.com/tinygo-org/tinygo/releases).

## Creating a Contract

In order to build a smart contract for the Versatus network, simply include the `versatus-go/contract` package:

```go
import "github.com/versatus/versatus-go/contract"
```

The `contract` object will then allow you to parse the contract input (received on stdin) to Go data structures, and to write the result on stdout. For example:

```go
_, inputs := contract.ParseInput()
```

and:

```go
var outputs contract.ComputeOutputs
_ = contract.SendOutput(outputs)
```

## Compiling a Contract

Simply use the TinyGo compiler command instead of the standard Go compiler command, and specify the target as WASI:

```
tinygo build --target=wasi \
    -scheduler=none -gc=conservative \
    -o contract.wasm ./...
```

This should result in a `contract.wasm` file that is the Web Assembly smart contract.

## Deploying a Contract to the Versatus Network

TBA

