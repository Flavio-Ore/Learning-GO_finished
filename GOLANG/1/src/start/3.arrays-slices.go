package main

import "fmt"

func main() {
	//------------------------------------------------------------------
	//* Arrays
	//! no podemos cambiar el largo de un array
	var ages [3]int = [3]int{17, 26, 38}

	fmt.Println(ages)
	fmt.Println(ages, len(ages))
	//------------------------------------------------------------------
	names := [5]string{"hola", "esto", "es", "un", "arreglo"}
	names[2] = "sigue siendo"
	fmt.Println(names)
	fmt.Println(names, len(names))

	//------------------------------------------------------------------
	//* Slices (use arrays under the hood)
	// Creando un slice, no un array (se parecen)
	//! podemos cambiar los valores, len
	scores := []int{99, 55, 100}

	fmt.Println(scores)

	scores[2] = 9999 //change

	scores = append(scores, 159357) //add values

	fmt.Println(scores, len(scores)) //len lenght

	//------------------------------------------------------------------
	//* slice ranges
	rangeOne := ages[1:3] //the position between
	fmt.Println(rangeOne, len(rangeOne))

	rangeTwo := names[2:] //starts from array 2 until the end
	fmt.Println(rangeTwo, len(rangeTwo))

	rangeThree := names[:3] //ends from array 3 until the end
	fmt.Println(rangeThree, len(rangeThree))

	//* slices
	rangeTwo = append(rangeTwo, "ARRAYS") //added
	fmt.Println(rangeTwo)
}
