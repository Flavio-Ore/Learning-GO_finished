package main

import "fmt"

func main() {
	var alvin string = "alvin"

	alvin2 := "teodoro"

	var alvinFinal = fmt.Sprintf("%v , %v\n ", alvin2, alvin)

	fmt.Printf("alvin es un: %T", alvinFinal)

}
