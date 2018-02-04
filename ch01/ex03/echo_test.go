package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func BenchmarkAppendWithNaive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s, sep := "", ""
		for _, arg := range os.Args[1:] {
			s += sep + arg
			sep = " "
		}
		fmt.Println(s)
	}
}

func BenchmarkAppendWithStringsJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s, sep := "", " "
		s = strings.Join(os.Args[1:], sep) + sep
		fmt.Println(s)
	}
}

