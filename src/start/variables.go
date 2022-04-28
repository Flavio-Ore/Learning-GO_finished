package main

import "fmt"

func main() {

	//* Strings------------------------------------------------------
	var nameOne string = "MarioBros" //only strings with "" | Explicit
	fmt.Println(nameOne)

	var nameTwo = "Luigi" //without string | Implicit
	fmt.Println(nameTwo)

	var nameThree string //Empty string
	fmt.Println(nameThree)

	//! nameOne = 123 // I can't change their value

	nameOne = "Peach" // variable reassignment
	nameTwo = "Pachac"
	fmt.Println(nameOne, nameTwo)

	nameFour := "" //shorthan , we can't use := outside of a func
	nameFour = "Joshua"
	fmt.Println(nameFour)

	//* Ints------------------------------------------------------

	var ageOne int = 10
	var ageTwo = 30
	ageThree := 40
	//? Podemos usar variaciones del int para específicar el tamaño de los bits a integrar
	fmt.Println(ageOne, ageTwo, ageThree)

	//* Bits and Memory--------------------------------------------

	//? int8,16,32,64 is the set of all signed 8,16,32,64-bit integers.
	/*
		*  Ranges:
		?  8-bis : -128 through 127.
		? 16-bis : -32768 through 32767.
		? 32-bis : -2147483648 through 2147483647
		? 64-bis :  -9223372036854775808 through 9223372036854775807
		? El autocompletado se encarga de los rangos :)
	*/

	//todo integra 32 bits sin específicar, pero es un tipo distinto
	var num int = 32
	//* La mayoria de las veces no es necesario especificar los bits
	var numOne int8 = 25
	var numTwo int16 = 3276
	var onlyPositiveNumber uint = 999 //uint starts from 0
	var numNine int
	fmt.Println(num, numOne, numTwo, onlyPositiveNumber, numNine)

	//* floats : wich has a decimal point, if it has decimals, is a float
	var scoreOne float32 = 25.95
	var scoreTwo float64 = 56165181813124894.89
	fmt.Println(scoreOne, scoreTwo)

	scoreThree := 123.123 //! float-64
	scoreFour := 123      //! int
	fmt.Println(scoreThree, scoreFour)
}
