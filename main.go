package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	info, _ := os.Stdin.Stat()
	if (info.Mode() & os.ModeCharDevice) == os.ModeCharDevice {
		// TODO: what to do?
		if len(os.Args) < 2 {
			fmt.Println("Uhh here's some help")
		}
		s := transform(smallcaps, os.Args[1:]...)
		fmt.Println(s)
	} else if info.Size() > 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			str := transform(smallcaps, scanner.Text())
			fmt.Println(str)
		}
	}
}

func transform(mapping map[string]string, input ...string) string {
	sb := strings.Builder{}
	for _, strs := range input {
		for _, char := range strs {
			if val, ok := mapping[strings.ToUpper(string(char))]; ok {
				sb.WriteString(val)
			} else {
				sb.WriteString(string(char))
			}
		}
		// Write space if we're taking in multiple inputs
		sb.WriteString(" ")
	}

	return sb.String()
}
