package main

import (
	"fmt"
)

func main() {
	age, age2 := 64, 30
	fmt.Println(age <= 64, age >= 64, age == 64, age != 64)

	if age < 30 {
		fmt.Println("age is less than 30")
	} else if age < 50 {
		fmt.Println("age is less than 50")
	} else {
		fmt.Println("age is greater than 50")
	}

	fmt.Println(age2)

	//----------------------------------------------------------------

	names := []string{"Juan", "Stuart", "Juniors", "Pepe", "Gon", "Berth"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("continuing at pos", index) // index = 1
			continue                                //to continue the loop
		}
		if index > 2 {
			fmt.Println("breaking at pos", index) // index > 2
			break                                 //break and doesn't print if index > 2
		}

		fmt.Printf("the value at pos %v is %v \n", index, value)
	}

}
