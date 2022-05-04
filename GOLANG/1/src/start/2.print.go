package main

import "fmt"

func main() {
	//variables
	age := 17
	name := "Jhon"

	//-------------------------------------------------------------------
	//-------------------------------------------------------------------

	//* Print-------------------------------------------------
	fmt.Print("Heyo, ")
	//? output: this prints doesn't add a new line to the end of the string
	fmt.Print("boy? \n")

	//-------------------------------------------------------------------

	fmt.Print("hello, ")
	//? \n to create new lines in Strings
	fmt.Print("world! \n")
	fmt.Print("new line :o \n")

	//-------------------------------------------------------------------

	//* Print + ln ----------------------------------------------

	fmt.Println("hello boy!")
	//? Print + ln do the same as well as \n
	fmt.Println("cya boys")

	fmt.Println("My age is", age, "and my name is", name) //UNFORMATTING

	//-------------------------------------------------------------------

	//! Formatting Strings, a way to create variables embedded inside
	//* Print + f (formatted strings)

	//? the '_' will be replaced by extra arguments, variables and the orders matters
	//? fmt.Printf("My new age is _ and my name is _", age, name)

	//* Printf : %_ = format speciefier
	//-------------------------------------------------------------------

	//* To be more specifed, is used %v
	//todo__ como encontró el %v entonces buscará la primera variable que encuentre, el segundo %v buscará la segunda variable
	fmt.Printf("My new age is %v and my name is %v \n", age, name)
	//* %v = _ : un espacio para recibir información
	//* We can use any different letters

	//?fmt.Printf("My other age is %q and my last name is %q \n", age, name)
	/*
		! 	output:
		?	Esto ocurre porque age es un integer y solo funciona en Strings
		*	 "My other age is '\x11' and my last name is "Jhon" \n"
	*/

	//-------------------------------------------------------------------

	//* %T the type of data
	fmt.Printf("age is a type of %T \n", age)
	fmt.Printf("age is a type of %T \n", name)

	//-------------------------------------------------------------------

	//* %f to embebb flows numbers
	fmt.Printf("you scored %f points \n", 123.123)
	//? The limit of number of decimal points
	//? Aproxima el numero que especifiquemos
	fmt.Printf("you scored %0.9f points \n", 123.123)

	//-------------------------------------------------------------------
	//-------------------------------------------------------------------

	//* S + Print + f (save formatted strings)
	//? Sprintf is not going  to directly print it out to the console, instead it returns a formatted string to us and we can store it in a variable
	//jalamos el mismo ejemplo
	var str = fmt.Sprintf("My new age is %v and my name is %v \n", age, name)
	fmt.Println("My data is:", str)
}
