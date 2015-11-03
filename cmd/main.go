package main

import (
	"fmt"
	"os"

	"github.com/dahernan/phash"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "USAGE command imagepath")
		os.Exit(1)
		return
	}

	filename := os.Args[1]
	_, err := os.Stat(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File not found: %v", err)
		os.Exit(1)
		return
	}

	h, err := phash.ImageHashDCT(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error calculating phash: %v", err)
		os.Exit(1)
		return
	}

	fmt.Fprintf(os.Stdout, "%v", h)

}
