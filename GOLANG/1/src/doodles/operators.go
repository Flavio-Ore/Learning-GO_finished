package main

import (
	"fmt"
)

func main() {
	var a, b int
	//input
	fmt.Print("Number 01: \n")
	fmt.Scanln(&a)
	fmt.Print("Number 02: \n")
	fmt.Scanln(&b)

	//processing
	addition, subtraction, multiplication, division, modulus := a+b, b-a, a*b, a/b, a%b

	//output
	fmt.Print("Results a+b: \n", addition, "Results a-b: \n", subtraction, "Results a*b: \n", division, "Results a/b: \n", multiplication, "Results a%b: \n", modulus)
}
