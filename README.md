# fun-with-monorepos

Test repo for fun with monorepos using Go and Bazel!

* Go packages reside at the root of the repo, and Bazel is used to build and test them.
* The `apps` directory contains Go applications that depend on the packages in the root of the repo.
  * Each of the applications is built and tested using Bazel and produces a Docker image.


## Setup

```shell
brew install bazelisk go
```

## Build

```shell
bazelisk build //...
```

## Test

```shell
bazelisk test //...
```

## Developing

### Adding a new Go package

```shell
# Install new package as normal
go get github.com/your/thing

# Update Bazel with new Go package
bazelisk run //:gazelle-update-repos
```

### Adding New Go Test, Package, or Go Code File

```shell
# Register new files with Bazel via Gazelle
bazelisk run //:gazelle
```