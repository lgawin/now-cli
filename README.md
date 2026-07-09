# now-cli
A fast, lightweight CLI utility suite to shift system time and seamlessly synchronize root-enabled Android device clocks via ADB

## Run tests
```shell
go test -v -race ./...
```

## Building (manual)
```shell
GOOS=linux GOARCH=amd64 go build -o build/now
cd build
tar -czvf now-linux-x86_64.tar.gz now
```

## Install
```shell
curl -sL "https://raw.githubusercontent.com/lgawin/now-cli/refs/heads/main/install.sh" | sh
```
