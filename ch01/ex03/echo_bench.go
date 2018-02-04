package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const BENCHMARK_COUNT=1000 

func main() {
	var naive_ave, stringsjoin_ave int64
	naive_ave, stringsjoin_ave = 0 ,0

	fmt.Printf("Benchmark count:  %v\n", BENCHMARK_COUNT)

	for i:=0; i < BENCHMARK_COUNT; i++ {
		naive_time_start := time.Now()
		AppendWithNaive()
		naive_time_end := time.Now()
		naive_ave += naive_time_end.Sub(naive_time_start).Nanoseconds();
	}
	fmt.Printf("Function Naive Ave.:  %v nanosec\n", naive_ave / 1000)

	for i:=0; i < BENCHMARK_COUNT; i++ {
		stringsjoin_time_start := time.Now()
		AppendWithStringsJoin()
		stringsjoin_time_end := time.Now()
		stringsjoin_ave += stringsjoin_time_end.Sub(stringsjoin_time_start).Nanoseconds();
	}
	fmt.Printf("Function StringsJoin Ave.:  %v nanosec\n", stringsjoin_ave / 1000)
}

func AppendWithNaive() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func AppendWithStringsJoin() {
	s, sep := "", " "
	s = strings.Join(os.Args[1:], sep) + sep
	fmt.Println(s)
}

