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

See [our internal wiki](https://github.com/omnicate/loltel/wiki/Public-APIs#releasing-to-the-maven-central-repository) for the required local tooling/setup for releasing to the maven central repository.

## Java/Maven artifacts

The java artifacts are released to maven central. We use semver, but the first number/major
version iteration represents the "v0" or "v1" part of the api.
That is, all releases of version 1 of the sms api should have `1` as their major number.
Running `release:prepare` will increase the patch version by one (`major.minor.patch`).
If you need to increase the major or minor version, increment this manually in the various
pom.xml files before running the commands below.
E.g. if the current version is `x.y.z-SNAPSHOT`, version `x.y.z` is the version that will be released next.
To instead release version `x.b.0`, edit the pom.xml files and set the version property
to `x.b.0-SNAPSHOT`.

Note that maven needs to be run from within the wgtwo directory because of reasons only maven
knows.

### Bump v1/v0 to specific version

Useful for when we add a new module or for any other reason want to explicitly set the
next version before a release:

```bash
./mvnw build-helper:parse-version versions:set -DnewVersion=ASPARAGUS-SNAPSHOT versions:commit --file wgtwo/pom-vX.xml
```
Replace "`ASPARAGUS`" with the desired x.y.z version number.
Also replace "vX" with v0 or v1 depending on which version of the api you need to target.

### Release/deploy v0


1. `./mvnw build-helper:parse-version release:prepare -B --file wgtwo/pom-v0.xml`
2. `./mvnw release:perform --file wgtwo/pom-v0.xml`
3. Access [Sonatype's repo manager](https://s01.oss.sonatype.org/#welcome) and confirm the
   artifacts for release.

### Release/deploy v1

1. `./mvnw build-helper:parse-version release:prepare -B --file wgtwo/pom-v1.xml`
2. `./mvnw release:perform --file wgtwo/pom-v1.xml`
3. Access [Sonatype's repo manager](https://s01.oss.sonatype.org/#welcome) and confirm the
   artifacts for release.

### Troubleshooting

If you get an error about "gpg: signing failed: Inappropriate ioctl for device", run these commands in your shell and try again:

```
GPG_TTY=$(tty)
export GPG_TTY
```

#### Rolling back a `release:prepare` that went horribly wrong

Maybe you forgot to set the intended version number in the pom.xml files before running the
command. Maybe you didn't know that it would create new commits. Anyway those commits have
now been pushed to github so you'll have to edit the history and do a force push.

1. run `./mvnw release:rollback --file wgtwo/pom-vX.xml`
1. delete the tag created by maven locally: `git tag -d $TAG`
2. delete the tag on the remote: `git push --delete $REMOTE $TAG`

## Protobuf FileDescriptorSet release

This requires that you have `buf` installed. See note below.

1. `buf build -o image.bin`

## Go release

This requires that you have `buf` and the Go protoc plugins installed. See note below.

We do include generated go code as part of this repo, which is generated by Buf

1. `buf generate`

## Buf and other required tooling

Building the FileDescriptorSet and Go code requires some tooling:

1. `buf` (https://docs.buf.build/installation)
2. `protoc-gen-go` (https://pkg.go.dev/google.golang.org/protobuf)
2. `protoc-gen-go-grpc` (https://pkg.go.dev/google.golang.org/grpc/cmd/protoc-gen-go-grpc)

We have [Hermit](https://github.com/cashapp/hermit) support, which will provide the required packages.

## NPM package

The release for the Node.js Package Manager so far only contains the .proto files, since the other contents of the origin repository are not relevant for this type of packaging.

The package is available both on the [Github Registry](https://github.com/orgs/working-group-two/packages?repo_name=wgtwoapis) and on [NPMJS](https://www.npmjs.com/package/@working-group-two/wgtwoapis).
