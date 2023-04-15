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

# TODOs / Ideas
- [ ] Connect-go service
- [ ] Buf lint and breaking detection
- [ ] Generate Tech Docs
- [x] Change image versioning to git tags
- [ ] Link packages and containers to releases
- [ ] Swagger spec generation
- [ ] GRPC clients in multiple languages
- [x] Setup golangci-lint
- [x] Setup release-please
- [x] Setup Bazel build and test

# Resources Used For Developing The Monorepo
* https://bazelbuild.github.io/rules_nodejs/stamping.html
* https://github.com/bazelbuild/rules_docker#stamping
* https://github.com/google-github-actions/release-please-action
* https://github.com/golangci/golangci-lint-action
* https://github.com/orgs/community/discussions/26686
* https://github.com/orgs/community/discussions/27028
* https://github.com/orgs/community/discussions/25617
* https://docs.github.com/en/actions/using-workflows/triggering-a-workflow#triggering-a-workflow-from-a-workflow