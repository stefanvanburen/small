package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"text/tabwriter"

	"github.com/peterbourgon/ff/v3/ffcli"

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
	fs := flag.NewFlagSet(name, flag.ExitOnError)
	var (
		transformName = fs.String("t", "", "specify transform type")
		list          = fs.Bool("l", false, "list transform types")
	)

	root := &ffcli.Command{
		Name:       name,
		FlagSet:    fs,
		ShortUsage: "small [FLAGS] <text>",
		Exec: func(_ context.Context, args []string) error {
			if *list {
				supportedTransforms := small.SupportedTransformations()
				supportedTransformNames := make([]string, 0, len(supportedTransforms))
				for supportedTransformName := range supportedTransforms {
					supportedTransformNames = append(supportedTransformNames, supportedTransformName)
				}
				sort.Strings(supportedTransformNames)
				w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
				for _, name := range supportedTransformNames {
					out := &bytes.Buffer{}
					small.PerformTransform(supportedTransforms[name], bytes.NewBufferString(name), out)
					fmt.Fprintf(
						w,
						"  %s\t%s\n",
						name,
						out,
					)
				}
				fmt.Fprintln(stdout, "Supported transformations:")
				return w.Flush()
			}

			info, err := stdin.Stat()
			if err != nil {
				return fmt.Errorf("can't access stdin: %s", err)
			}

			transform, err := small.GetTransform(*transformName)
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
