package main

import (
	"fmt"
	"math"
)

//creating functions
//? n string, this functions arguments is gonna be a string
func sayGreeting(n string) {
	fmt.Printf("Good afternoon %v \n", n)
}

func sayBye(n string) {
	fmt.Printf("Goodbye %v \n", n)
}

//? Calling a function for each element inside on an array
func cycleNames(name []string, function func(string)) {
	for _, value := range name { //each element of the slice is a value
		function(value)
	}
}

// var avar = func (string){}

func circleArea(radius float64) float64 { //! with float64 I'm specifed the type of numer that this function gonna return before {}
	//? return the value of the circle
	return math.Pi * radius * radius
}

func main() {
	//calling functions

	//? sayGreeting("Steven")
	//? sayGreeting("Gon")
	//? sayGreeting("Stuart")

	cycleNames([]string{"Stuart", "Gon", "Steven"}, sayGreeting)
	// sayGreeting and SayBye are reference of f function
	cycleNames([]string{"Stuart", "Gon", "Steven"}, sayBye)

	area1 := circleArea(13.5)
	area2 := circleArea(7)

	fmt.Println(area1, area2) //output: many decimals
	fmt.Printf("Circle 1 is %0.3f and circle 2 is %0.3f\n", area1, area2)
	//? %0.3f to three decimals places
}
