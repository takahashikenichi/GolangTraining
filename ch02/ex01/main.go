package main

import (
	"fmt"
//	"./src"
	"github.com/takahashikenichi/GolangTraining/ch02/ex01/tempconv"
)

func main() {
	fmt.Println(tempconv.CToK(tempconv.FreezingC))
	fmt.Println(tempconv.CToF(tempconv.FreezingC))
	fmt.Println(tempconv.KToF(tempconv.AbsoluteZeroK))
}
