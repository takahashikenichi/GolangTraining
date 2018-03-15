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

func BenchmarkPopCountByShift64(b *testing.B) {
	results = []int{}
	for i := 0; i < b.N; i++ {
		results = append(results, PopCountByShift64(uint64(i)))
	}
}

func BenchmarkPopCountByNewFormula(b *testing.B) {
	results = []int{}
	for i := 0; i < b.N; i++ {
		results = append(results, PopCountByNewFormula(uint64(i)))
	}
}
