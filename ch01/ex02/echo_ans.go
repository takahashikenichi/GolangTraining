package main

import (
	"fmt"
	"os"
)

func main() {
	for i, s : range os.Args {
		fmt.Printf("%v, %v\n", i, s)
	}
}
