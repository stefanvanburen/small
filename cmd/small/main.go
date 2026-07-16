package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"maps"
	"os"
	"slices"
	"strings"
	"text/tabwriter"

	"go.vanburen.xyz/small/internal/small"
)

const name = "small"

func main() {
	if err := run(os.Args, os.Stdout, os.Stdin); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}

func run(args []string, stdout io.Writer, stdin *os.File) error {
	flagSet := flag.NewFlagSet(name, flag.ExitOnError)
	var transformName = flagSet.String("t", "", "specify transform type")
	flagSet.Usage = func() {
		var usage strings.Builder
		usage.WriteString(`USAGE
  small [-t] <text>

FLAGS
  -t ... specify transform type

SUPPORTED TRANSFORMS
`)

		// sort the list of supported transforms, for stable output
		supportedTransforms := small.SupportedTransformations()
		supportedTransformNames := slices.Collect(maps.Keys(supportedTransforms))
		slices.Sort(supportedTransformNames)

		tabWriter := tabwriter.NewWriter(&usage, 0, 0, 1, ' ', 0)
		for _, name := range supportedTransformNames {
			out := &bytes.Buffer{}
			small.PerformTransform(supportedTransforms[name], bytes.NewBufferString(name), out)
			fmt.Fprintf(
				tabWriter,
				"  %s\t%s\n",
				name,
				out,
			)
		}
		tabWriter.Flush()

		fmt.Fprint(os.Stderr, usage.String())
	}

	if err := flagSet.Parse(args[1:]); err != nil {
		return err
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
		if len(flagSet.Args()) != 1 {
			return fmt.Errorf("only a single argument is supported")
		}

		buf := bytes.NewReader([]byte(flagSet.Args()[0]))
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
}
