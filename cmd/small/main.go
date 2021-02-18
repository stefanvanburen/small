package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/peterbourgon/ff/v3/ffcli"
	"github.com/svanburen/small/pkg/small"
)

func main() {
	if err := run(os.Args, os.Stdout, os.Stdin); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}

const name = "small"

func run(args []string, stdout io.Writer, stdin *os.File) error {
	fs := flag.NewFlagSet(name, flag.ExitOnError)
	transformName := fs.String("t", "", "specify transform type")
	list := fs.Bool("l", false, "list transform types")

	root := &ffcli.Command{
		Name:    name,
		FlagSet: fs,
		UsageFunc: func(c *ffcli.Command) string {
			return `USAGE
  small [FLAGS] TEXT
  command | small [FLAGS]

FLAGS
  -h         print this help message
  -l         list transform types
  -t=TYPE    specify transform type
`
		},
		Exec: func(_ context.Context, args []string) error {
			if list != nil && *list {
				fmt.Fprintln(stdout, "Supported transformations:")

				for _, supportedTransformation := range small.SupportedTransformations() {
					fmt.Fprintf(stdout, "  %s\n", supportedTransformation)
				}

				return nil
			}

			info, err := stdin.Stat()
			if err != nil {
				return fmt.Errorf("can't access stdin: %s", err)
			}

			var name string
			if transformName != nil {
				name = *transformName
			}

			transform, err := small.GetTransform(name)
			if err != nil {
				return err
			}

			if (info.Mode() & os.ModeCharDevice) == os.ModeCharDevice {
				// Normal
				if len(args) == 0 {
					return flag.ErrHelp
				}

				if len(args) > 1 {
					return flag.ErrHelp
				}

				buf := bytes.NewReader([]byte(args[0]))
				if err := small.PerformTransform(transform, buf, stdout); err != nil {
					return fmt.Errorf("could not write to stdout: %s", err)
				}

				// To play nicely with the terminal, write out a newline.
				stdout.Write([]byte("\n"))
			} else if info.Size() > 0 {
				// We're being piped to
				// command | small ...
				if err := small.PerformTransform(transform, stdin, stdout); err != nil {
					return fmt.Errorf("could not write to stdout: %s", err)
				}
			}

			return nil
		},
	}

	return root.ParseAndRun(context.Background(), args[1:])
}
