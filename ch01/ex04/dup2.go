package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	countsByFile := make(map[string]map[string]int)

	files := os.Args[1:]
	if len(files) == 0 {
		countsPerFile := make(map[string]int)
		countLines(os.Stdin, counts, countsPerFile)
		countsByFile["stdin"] = countsPerFile
	} else {
		for _, arg := range files {
			countsPerFile := make(map[string]int)
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, countsPerFile)
			countsByFile[arg] = countsPerFile
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			for filename, counts := range countsByFile {
				if counts[line] > 0 {
					fmt.Printf("%s\t", filename)
				}
			}
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

/*
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// ignoring input.Err() 
}
*/

// 別のMapを引数にcountLineを2回呼ぶんでも2回設定してくれない
func countLines(f *os.File, counts map[string]int, countsPerFile map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		countsPerFile[input.Text()]++
	}
	// ignoring input.Err() 
}
