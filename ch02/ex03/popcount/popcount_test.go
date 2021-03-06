package popcount

import (
	"testing"
)

var results []int

func BenchmarkPopCount(b *testing.B) {
	results = []int{}
	for i := 0; i < b.N; i++ {
		results = append(results, PopCount(uint64(i)))
	}
}

func BenchmarkPopCountByLoop(b *testing.B) {
	results = []int{}
	for i := 0; i < b.N; i++ {
		results = append(results, PopCountByLoop(uint64(i)))
	}
}
