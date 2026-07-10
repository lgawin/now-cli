package main

import (
	"errors"
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
	out io.Writer,
	args []string,
	nowFunc func() time.Time,
) int {
	fs := flag.NewFlagSet("now", flag.ContinueOnError)
	fs.SetOutput(out)
	fs.Usage = func() {
		_, _ = fmt.Fprintf(out, "Usage of %s:\n", args[0])
		_, _ = fmt.Fprintf(out, "  A fast CLI tool to generate the current time in various formats.\n\n")
		_, _ = fmt.Fprintf(out, "Options:\n")
		_, _ = fmt.Fprintf(out, "  -f, --format <string>   Time format (e.g., 'epoch' or 'iso') (default \"iso\")\n")
		_, _ = fmt.Fprintf(out, "  -v, --version           Print version\n")
		_, _ = fmt.Fprintf(out, "  -h, --help              Show this help message\n")
	}

	format := fs.String("format", "iso", "Time format (e.g., 'epoch' or 'iso')")
	fs.StringVar(format, "f", "iso", "Shorthand for format")

	version := fs.Bool("version", false, "Print version")
	fs.BoolVar(version, "v", false, "Shorthand for version")

	if err := fs.Parse(args[1:]); err != nil {
		// If the error is flag.ErrHelp (caused by -h or --help), exit successfully
		if errors.Is(err, flag.ErrHelp) {
			return 0
		}
		return 1
	}

	if *version {
		if _, err := fmt.Fprintf(out, "now-cli version: %s\n", Version); err != nil {
			panic(err)
		}
		return 0
	}

	if *format == "" {
		*format = "iso"
	}

	switch *format {
	case "epoch":
		if _, err := fmt.Fprintln(out, nowFunc().Unix()); err != nil {
			panic(err)
		}
	case "iso":
		if _, err := fmt.Fprintln(out, nowFunc().Format(time.RFC3339)); err != nil {
			panic(err)
		}
	default:
		if _, err := fmt.Fprintf(out, "unknown format: %s\n", *format); err != nil {
			panic(err)
		}
	}
	return 0
}
