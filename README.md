# BMS gRPC API

This repository contains the gRPC API through which monitors report their measurements to the BMS server.

Currently this API only consists of a storage service (see [`storage.proto`](storage.proto)).
Commonly usable types are separated into [`common.proto`](common.proto).


## Generated Go Interface

Besides the gRPC API specification this repository also contains automatically derived Go code (in [`gen/`](gen/)) which clients and servers use as a package to implement.

This code has to be generated when making changes to the `*.proto` files by running:
```
$ buf generate
```
(Make sure that [`buf`](https://buf.build/docs/installation) is installed.)

When changing the `*.proto` files, please make sure that you're not adding any lints.
You can check your changes by running:
```
$ buf lint
```
If you're using VSCode, you can also install the [Buf extension](https://buf.build/docs/editor-integration#visual-studio-code-visual-studio-marketplace-downloads) which automatically lints your Protobuf files.

## API Reference
Running `buf generate` also generates an API reference which can be found in the [`docs/`](docs/) folder.

## Usage
To implement a new client that uses this API, you have to import this package:
```go
import bmsapi "github.com/botnet-monitoring/grpc-api/gen"
```
For how to get started with the actual implementation see e.g. [this tutorial](https://grpc.io/docs/languages/go/basics/#client) or consult our [client implementation](https://github.com/botnet-monitoring/grpc-client).

If you just want to send data to BMS, **you probably want to use our more ergonomic [gRPC client](https://github.com/botnet-monitoring/grpc-client) package**.
