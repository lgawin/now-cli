package main

import (
	"flag"
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
	fs := flag.NewFlagSet("now", flag.ContinueOnError)
	fs.SetOutput(w)

	var format string
	fs.StringVar(&format, "f", "", "Output format (epoch, iso)")
	fs.StringVar(&format, "format", "", "Output format (epoch, iso)")

	var version bool
	fs.BoolVar(&version, "v", false, "Print tool version")
	fs.BoolVar(&version, "version", false, "Print tool version")

	if len(args) > 1 {
		if err := fs.Parse(args[1:]); err != nil {
			return
		}
	}

	if version {
		if _, err := fmt.Fprintf(w, "now-cli version: %s\n", Version); err != nil {
			panic(err)
		}
		return
	}

	if format == "" {
		format = "iso"
	}

	switch format {
	case "epoch":
		if _, err := fmt.Fprintln(w, nowFunc().Unix()); err != nil {
			panic(err)
		}
	case "iso":
		if _, err := fmt.Fprintln(w, nowFunc().Format(time.RFC3339)); err != nil {
			panic(err)
		}
	default:
		if _, err := fmt.Fprintf(w, "unknown format: %s\n", format); err != nil {
			panic(err)
		}
	}
}
