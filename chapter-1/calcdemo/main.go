package main

import (
	"fmt"

	"github.com/shijuvar/go-web/chapter-1/calc"
)

func main() {
	//var x, y int= 10, 5
	x, y := 10, 5
	fmt.Println(calc.Add(x, y))
	fmt.Println(calc.Subtract(x, y))
}
