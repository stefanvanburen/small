package main

import (
	"bufio"
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
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run(args []string, stdout io.Writer, stdin *os.File) error {
	fs := flag.NewFlagSet(args[0], flag.ExitOnError)
	transformName := fs.String("t", "", "specify transform type")
	list := fs.Bool("l", false, "list transform types")

	root := &ffcli.Command{
		FlagSet: fs,
		UsageFunc: func(c *ffcli.Command) string {
			return `USAGE
  small [FLAGS] [TEXT...]
  command | small [FLAGS]

FLAGS
  -h         print this help message
  -l         list transform types
  -t <type>  specify transform type
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
				return err
			}

			if (info.Mode() & os.ModeCharDevice) == os.ModeCharDevice {
				// Normal
				if len(args) == 0 {
					return flag.ErrHelp
				}

				name := ""
				if transformName != nil {
					name = *transformName
				}

				transform, err := small.GetTransform(name)
				if err != nil {
					return err
				}

				if _, err := fmt.Fprintln(stdout, small.PerformTransform(transform, args...)); err != nil {
					return err
				}
			} else if info.Size() > 0 {
				// We're being piped to
				// command | sm ...
				name := ""
				if transformName != nil {
					name = *transformName
				}

				transform, err := small.GetTransform(name)
				if err != nil {
					return err
				}

				scanner := bufio.NewScanner(stdin)
				for scanner.Scan() {
					str := small.PerformTransform(transform, scanner.Text())
					_, err = fmt.Fprintln(stdout, str)
					if err != nil {
						return err
					}
				}
			}

			return nil
		},
	}

	return root.ParseAndRun(context.Background(), args[1:])
}
