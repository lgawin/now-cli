package main

import (
	"fmt"
	"io"
	"os"
)

var Version = "development"

func main() {
	run(os.Stdout)
}

func run(w io.Writer) {
	_, _ = fmt.Fprintf(w, "now-cli version: %s\n", Version)
}
