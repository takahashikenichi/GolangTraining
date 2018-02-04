package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for i := 0; i < len(os.Args); i++ {
		s += sep + "Args[" + fmt.Sprint(i) + "] : " + os.Args[i]
		sep = "\n"
	}
	fmt.Println(s)
}
