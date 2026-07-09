# now-cli
A fast, lightweight CLI utility suite to shift system time and seamlessly synchronize root-enabled Android device clocks via ADB


## Building (manual)
```shell
GOOS=linux GOARCH=amd64 go build -o build/now
cd build
tar -czvf now-linux-x86_64.tar.gz now
```

## Install
**FIXME**: fix file path before merging PR
```shell
curl -sL "https://raw.githubusercontent.com/lgawin/now-cli/refs/heads/task-1-4/install-script/install.sh" | sh
```
