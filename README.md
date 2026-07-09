now-cli
=======
A fast, lightweight CLI utility suite to shift system time and seamlessly synchronize root-enabled Android device clocks via `adb`.

[![Build status](https://img.shields.io/github/actions/workflow/status/lgawin/now-cli/ci.yml?branch=main&label=ci&style=flat-square)](https://github.com/lgawin/now-cli/actions/workflows/ci.yml)
[![GitHub Release](https://img.shields.io/github/v/release/lgawin/now-cli?include_prereleases&label=version&style=flat-square&color=blue)](https://github.com/lgawin/now-cli/releases)
[![License](https://img.shields.io/github/license/lgawin/now-cli?label=license&style=flat-square&color=orange)](https://github.com/lgawin/now-cli/blob/main/LICENSE)

Installation
------------
Install the latest stable version of the `now` binary and orchestrator tools directly into `/usr/local/bin`:

```shell
curl -sL "https://raw.githubusercontent.com/lgawin/now-cli/main/install.sh" | sh
```

### Verify Installation
```shell
now --version
```

Development & Building
------------

### Prerequisites
Ensure you have Go installed. If you use `mise`, the environment toolchain will auto-configure upon entering the directory.

### Running tests

Execute the local unit and integration test suite with the race detector enabled:

```shell
go test -v -race ./...
```

### Manual Compilation

To cross-compile the release asset manually for Linux x86_64:

```shell
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o dist/now
cd dist
tar -czvf now-linux-x86_64.tar.gz now
```

Releasing
---------

Deployments are entirely automated via GitHub CD pipelines. To cut a new release, simply tag your commit and push it up:

```shell
VERSION=... # provide version, like v0.0.1 
git tag $VERSION && git push origin $VERSION
```

License
-------

[MIT](LICENSE)
