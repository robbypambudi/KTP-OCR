package cmd

import (
	"fmt"
	"os"
)

// function to test ktp reader using cmdline
func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Printf("usage: %s <nama_file>", args[0])
	}
}
