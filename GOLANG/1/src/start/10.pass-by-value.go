//?------------------------------------------------------------------------ When a parameter is passed to a function by value, it means the parameter is copied into another location of your memory. When accessing or modifying the variable within your function, only the copy is accessed or modified — the original value is never modified. All primitive/basic types (int and its variants, float and its variants, boolean, string, array, and struct) in Go are passed by value. Passing by value is usually how your values are passed on to functions.

//! Go is a Pass-by-value Language

//* GO MAKES "COPIES" OF VALUES WHEN PASSED INTO FUNCTIONS

//--------------------------------------------------------------------------

package main

import "fmt"

//? When we pass variables around as arguments in functions go makes a copy of those values for the functions to use

//todo Group A -> Strings - Ints - Floats - Booleans - Arrays - Structs
//todo Group B -> Slices - Maps - Functions

//--------------------------------------------------------------------------
//a func that update the copy variable
func updateName(x string) { //x === string
	x = "GO"
	//! here we are UPDATING the COPY variable that we defined as name := "golang"
}

//--------------------------------------------------------------------------
//a func that update and post the copy variable
func updateNameFixed(x string) string { //specify the type of return x===string
	x = "GO"
	return x
	//! here we are UPDATING AND RETURNING the COPY variable that we defined as name := "golang"
}

//--------------------------------------------------------------------------
//a function that update the characters key str
func updateMenu(c map[string]float64) { // a map of exactly the same charactersMenu
	c["extraCharacter"] = 10.99
}

//--------------------------------------------------------------------------
func main() {

	//! Non-Pointer Values -> strings, ints, bools, floats, arrays, structs

	name, nameTransformer := "golang", "this golang gonna change"
	updateName(name) // this is a COPY of the name variable, not the original

	fmt.Println(name) // Printing the original one
	// output: golang

	//--------------------------------------------------------------------------

	nameTransformer = updateNameFixed(nameTransformer) //name value is gona be updated to equal whatever this function returns to us

	fmt.Println(nameTransformer) //Printing the original one but is return of the function
	//ouput: GO

	//--------------------------------------------------------------------------

	//* Every time we pass a value or a variable into a function, GO create a copy of the value

	/*
	   * Computer Memory
	   ? values =   age	    name	  n
	     stores =|0x0001 | 0x0002 | 0x0003 | 0x0004 | => Unique memory address
	   ?  	  	    35	  "golang" "golang"
	*/

	//--------------------------------------------------------------------------

	//! Pointer Wrapper Values -> Slices - Maps - Functions
	charactersMenu := map[string]float64{
		"character1": 389.99,
		"character2": 159.66,
		"character3": 16.80,
	}

	fmt.Println(charactersMenu) //output: original map charactersMenu

	updateMenu(charactersMenu) //output: updating the original map value

	fmt.Println(charactersMenu) //output: updated map charactersMenu

}

/*
todo| Computer Memory
todo| What's happening under the hood, why the map is updated?
*   | the pointer point into another memory location
            ____________________________
? values = |  menuCharacter             ↓
stores =   |--------0x0001  |          0x0002         |    0x0003   |=>UMA
?  	  	          [pointer]  {"character1": 389.99...}

First GO stores the uderlying data in memory 0x0002 in its own bloack
The value slipt into multiple memory locations

* | Later when we pass this value as an argument into a function, GO still make a copy of the variable that bit doesn't change we always make a copybut it's copying the value stored inside the 0x0001 memory block

            ___________________________    ____________________________
? values = |  menuCharacter            ↓  ↓                    c      |
stores =   |         0x0001  |         0x0002         |      0x0003   |=>UMA
?  	  	   |-------[pointer]  {"character1": 389.99...}   [pointer]---|

! The variable c (inside of the function updateMenu(m map[string] float64)) is stored into another memory location but it points to the same location as original variable map when we call it updateMenu(menuCharacter)

* The first pointer called menuCharacter is the original variable
* The second function pointer called c is the COPY of the original one
* Both pointer have the same values in the memory block
*/
