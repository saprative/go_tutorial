package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		str := input.Text()
		fmt.Println(basename(str))
	}

	if err := input.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error during read", err)
		os.Exit(1)
	}
}

func basename(s string) {

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
		}
		break
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}

	return s
}
