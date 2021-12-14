[![Actions](https://github.com/working-group-two/wgtwoapis/workflows/Test%20all%20JDKs%20on%20all%20OSes/badge.svg)](https://github.com/working-group-two/wgtwoapis/actions)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Java gRPC version](https://img.shields.io/badge/gRPC%20Version-%201.38.0-blue.svg)](https://grpc.io/)


# Working Group Two APIs
This repository contains the public APIs of Working Group Two.

Working Group Two's APIs uses gRPC backed by Protocol Buffers. This repository contains both the
.proto files that can be used to generate clients and documentation, in addition to ready made
Maven packages to simplify integration.

## Documentation
See https://docs.wgtwo.com

## Release/deploy

1. `./mvnw build-helper:parse-version release:prepare -B`
2. `./mvnw release:perform`
