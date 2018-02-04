#!/bin/bash

# benchmark without go test 
go run echo_bench.go aaa bbb ccc ddd eee fff

# go test
go test -bench=. 
# go test でベンチマークを使うときの引数の渡し方が不明 
