# Release process

## Setup required to release/deploy from local machine:

See [our internal wiki](https://github.com/omnicate/loltel/wiki/Public-APIs#releasing-to-the-maven-central-repository)
for the required local tooling/setup for releasing to the maven central repository.

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

If you get an error about "gpg: signing failed: Inappropriate ioctl for device", run these commands in your shell and
try again:

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

## Build and release protobuf artifacts

```shell
# Build FileDescriptorSet and generate Go code
make buf

# Push to Buf Schema Registry
make buf_push
```

### Push to Buf Schema Registry

This requires that you have logged in to the Buf Schema Registry:
```shell
buf registry login buf.build --username 'your username'
```

## Required tooling

1. Building image.bin and generating go code requires [buf](https://docs.buf.build/installation).
2. Building the Java code requires a compatible Java version.

### `mise`

We recommend using [mise](https://github.com/jdx/mise) for managing these dependencies.

If `mise` is installed, the required tooling should be automatically picked up from [.mise.toml](.mise.toml).

## NPM package

The release for the Node.js Package Manager so far only contains the .proto files, since the other contents of the
origin repository are not relevant for this type of packaging.

The package is available both on
the [Github Registry](https://github.com/orgs/working-group-two/packages?repo_name=wgtwoapis) and
on [NPMJS](https://www.npmjs.com/package/@working-group-two/wgtwoapis).
