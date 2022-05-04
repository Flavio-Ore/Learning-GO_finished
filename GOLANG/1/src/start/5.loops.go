package main

import (
	"fmt"
)

func main() {
	x := 0

	for x < 5 { // while loop, but is called always for
		fmt.Printf("value of x is: %v \n", x)
		x++ //to avoid infinite loop
	}

	for i := 0; i < 5; i++ {
		fmt.Printf("i++: %v \n", i)
		x++ //to avoid infinite loop
	}

	//* A loop for slice of integers or strings

	//todo----------------------------------------------------------------
	names := []string{"Juan", "Stuart", "Juniors", "Pepe"}
	//todo----------------------------------------------------------------

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	//----------------------------------------------------------------

	for index, value := range names {
		//* To know the position and value of names Array
		fmt.Printf("The value at index %v is: %v\n", index, value)
	}

	for _, value := range names {
		//* The underscore if we don't want to know the index of the value

		fmt.Printf("The value is %v\n", value)
		value = "new string" //value is a local copy
		//* this doesn't update the value of the slice
	}

}
