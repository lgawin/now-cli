now-cli
=======
A fast, lightweight CLI utility suite to shift system time and seamlessly synchronize root-enabled Android device clocks via `adb`.

Installation
------------
```shell
curl -sL "https://raw.githubusercontent.com/lgawin/now-cli/refs/heads/main/install.sh" | sh
```

Building `now`
------------

## Build instructions
```shell
GOOS=linux GOARCH=amd64 go build -o dist/now
cd dist
tar -czvf now-linux-x86_64.tar.gz now
```

## Running tests
```shell
go test -v -race ./...
```

License
-------

[MIT](LICENSE)
