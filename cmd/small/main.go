package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"

	"github.com/peterbourgon/ff/v4"

	"github.com/stefanvanburen/small/internal/small"
)

func main() {
	if err := run(os.Args, os.Stdout, os.Stdin); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}

const name = "small"

func run(args []string, stdout io.Writer, stdin *os.File) error {
	fs := ff.NewFlagSet(name)
	var transformFlag = fs.String('t', "--transform", "", "specify transform")

	root := &ff.Command{
		Name:  name,
		Flags: fs,
		Exec: func(_ context.Context, args []string) error {
			info, err := stdin.Stat()
			if err != nil {
				return fmt.Errorf("can't access stdin: %s", err)
			}

			transform, err := small.GetTransform(*transformFlag)
			if err != nil {
				return err
			}

			if (info.Mode() & os.ModeCharDevice) == os.ModeCharDevice {
				if len(args) != 1 {
					return fmt.Errorf("only a single argument is supported")
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
