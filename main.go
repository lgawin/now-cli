package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

var Version = "development"

func main() {
	run(os.Stdout, os.Args, time.Now)
}

func run(
	w io.Writer,
	args []string,
	nowFunc func() time.Time,
) {
	_, _ = fmt.Fprintf(w, "now-cli version: %s\n", Version)
}
