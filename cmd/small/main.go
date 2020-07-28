package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/peterbourgon/ff"
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
	fs := flag.NewFlagSet("small", flag.ExitOnError)
	transformName := fs.String("transform", "", "specify transform type")

	if err := ff.Parse(fs, args[1:]); err != nil {
		return err
	}

	root := &ffcli.Command{
		FlagSet: fs,
		Exec: func(_ context.Context, args []string) error {
			info, err := stdin.Stat()
			if err != nil {
				return err
			}

			if (info.Mode() & os.ModeCharDevice) == os.ModeCharDevice {
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
