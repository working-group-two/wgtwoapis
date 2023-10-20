[![Actions](https://github.com/working-group-two/wgtwoapis/workflows/Test%20all%20JDKs%20on%20all%20OSes/badge.svg)](https://github.com/working-group-two/wgtwoapis/actions)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Java gRPC version](https://img.shields.io/badge/gRPC%20Version-%201.38.0-blue.svg)](https://grpc.io/)

# Working Group Two APIs

This repository contains the public APIs of Working Group Two.

Working Group Two's APIs uses gRPC backed by Protocol Buffers. This repository contains both the
.proto files that can be used to generate clients and documentation, in addition to ready made
Maven packages to simplify integration.

**Contents:**

- [Java/Maven](#javamaven-artifacts)
- [Protobuf FileDescriptorSet release](#protobuf-filedescriptorset-release)
- [Go release](#go-release)
- [NPM package](#npm-package)

## Documentation

See https://docs.wgtwo.com

## Setup required to release/deploy from local machine:

See [our internal wiki](https://github.com/omnicate/loltel/wiki/Public-APIs#releasing-to-the-maven-central-repository)
for the required local tooling/setup for releasing to the maven central repository.

### Languages

## Java/Maven

The Java artifacts are released to maven central. We use semver, where major
version corresponds to the version set as part of the protobuf package.

## Go

Generated go code is included in this repo, next to their corresponding `.proto` files

## Node.js

> **Note:** This is currently experimental, so changes may occur.

We are publishing our .proto files to the Buf Schema Registry at [buf.build/wgtwo/wgtwoapis](https://buf.build/wgtwo/wgtwoapis).

This includes npm support.

## Other languages

[Buf Schema Registry](https://buf.build/wgtwo/wgtwoapis) offers support for:
- Go
- Node.js
- Maven
- Swift

We are not currently publishing artifacts for other languages, so please reach out if you have any needs.

However, gRPC has support for a variety of languages as displayed in[gRPC's documentation](https://grpc.io/docs/languages/).
Any of these may be used by using our protobuf files directly by syncing them from this repository.

That is, e.g. using the `dotnet-grpc` cli for .NET.
