package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	run(os.Stdout)
}

func run(w io.Writer) {
	_, _ = fmt.Fprintln(w, "now-cli skeleton v0.0.1")
}
