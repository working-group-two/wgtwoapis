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

## Setup required to release/deploy from local machine:

See [our internal wiki](https://github.com/omnicate/loltel/wiki/Public-APIs#releasing-to-the-maven-central-repository) for the required local tooling/setup for releasing to the maven central repository.

## Java/maven artifacts
### Release/deploy v0

1. `./mvnw build-helper:parse-version release:prepare -B --file wgtwo/pom-v0.xml`
2. `./mvnw release:perform --file wgtwo/pom-v0.xml`

### Release/deploy v1

1. `./mvnw build-helper:parse-version release:prepare -B --file wgtwo/pom-v1.xml`
2. `./mvnw release:perform --file wgtwo/pom-v1.xml`

### Troubleshooting

If you get an error about "gpg: signing failed: Inappropriate ioctl for device", run these commands in your shell and try again:

```
GPG_TTY=$(tty)
export GPG_TTY
```

#### rolling back a `release:prepare` that went horribly wrong

Maybe you forgot to set the intended version number in the pom.xml files before running the
command. Maybe you didn't know that it would create new commits. Anyway those commits have
now been pushed to github so you'll have to edit the history and do a force push.

1. delete the tag created by maven locally: `git tag -d $TAG`
2. delete the tag on the remote: `git push --delete $REMOTE $TAG`
3. roll back the commits: `git reset --hard $COMMIT`
(find the first commit prior to the [RELEASE] commits)
4. do a force push: `git push --force`

## Protobuf/buf/go release

1. `buf build -o image.bin`

### Troubleshooting

#### No `buf` command in PATH

Ask @gi on slack for assistance.
