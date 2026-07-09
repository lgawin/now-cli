# now-cli
A fast, lightweight CLI utility suite to shift system time and seamlessly synchronize root-enabled Android device clocks via ADB


## Building (manual)
```shell
GOOS=linux GOARCH=amd64 go build -o build/now
cd build
tar -czvf now-linux-x86_64.tar.gz now
```
