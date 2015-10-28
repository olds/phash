package main

import (
	"fmt"
	"os"

	"github.com/dahernan/phash"
)

func main() {

	h1, err := phash.VideoHash(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("h1: ", h1)
}
