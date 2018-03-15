package main

import (
	"fmt"
	"os"
	"strconv"
//	"./src"
	"github.com/takahashikenichi/GolangTraining/ch02/ex02/tempconv"
	"github.com/takahashikenichi/GolangTraining/ch02/ex02/distconv"
	"github.com/takahashikenichi/GolangTraining/ch02/ex02/weightconv"
)

func main() {
	if len(os.Args[1:]) < 2 {
		fmt.Println("input: Error argument.");
		os.Exit(1)
	}
	args1 := os.Args[1]
	args2 := os.Args[2]

	value, err := strconv.ParseFloat(args2, 64)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Args[2] is not collect: %value\n")
		os.Exit(1)
	}

	switch args1 {
		case "t":
			c := tempconv.Celsius(value)
			f := tempconv.Fahrenheit(value)
			fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
		case "d":
			m := distconv.Meter(value)
			f := distconv.Feet(value)
			fmt.Printf("%s = %s, %s = %s\n", m, distconv.MToF(m), f, distconv.FToM(f))
		case "w":
			p := weightconv.Pound(value)
			k := weightconv.KiloGram(value)
			fmt.Printf("%s = %s, %s = %s\n", p, weightconv.PToK(p), k, weightconv.KToP(k))
	}
}
