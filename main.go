package main

import (
	"fmt"
	"os"
)

func main() {
	// Validate arguments.
	if len(os.Args) != 3 || os.Args[1] == "" || os.Args[2] == "" {
		fmt.Fprintln(os.Stderr, "usage: atomicswap <first-path> <second-path>")
		os.Exit(1)
	}

	// Perform the swap.
	if err := swap(os.Args[1], os.Args[2]); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err.Error())
		os.Exit(1)
	}
}
