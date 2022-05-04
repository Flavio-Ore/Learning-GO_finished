package bill

import (
	"fmt"
	"os"
	"time"
)

//OBJECT STRUCTURE | can hold many types of values

//* In GO we don't need to use classes, instead we can create custom types using the STRUCT, it will be a blueprint which describes a type of data for that object

// OWN CUSTOM TYPE
type Bill struct { // ↓ properties of sthis struc
	name  string             // name type string
	items map[string]float64 // a map type
	tip   float64            // tip type float64
}

// make new bills
func NewBill(name string) Bill { // we assign the name value to the name properties of the bill struct

	var b = Bill{ // new object

		// initial build object
		name: name,

		items: map[string]float64{}, // gonna work with this items {"canned food": 9.99, "cookies": 2.99, "salad": 6.99}

		tip: 0,
	}

	//? This is a initial build object with its initial values
	return b // returning the b value
}

//* adding this parentheses and saying that we receive a bill into this function now we're limiting this function only to bill objects, so it can only be called from a build object now just doing this (b bill) ↓

// format the bill
func (b *Bill) Format() string { // format() is not associated with bill object

	fs := "Bill breakdown: \n" //the start of the formatting string

	//cycle through the items ↓
	var total float64 = 0 //add each item cost to the total, and output it

	// list items | list out the items inside the formatted string fs
	for k, v := range b.items { // b.items to call the items properties

		//we want to output the key value
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)

		total += v //to adding the cost to the total
	}
	// add tip

	fs += fmt.Sprintf("%-25v ...$%v\n", "tip:", b.tip)

	// total
	fs += fmt.Sprintf("%-25v ...$%0.2f \n", "total:", total+b.tip)

	/*
		I add -25 to add 25 characters long and it's going to push the prices
		I made this: Sprintf("%-25v ...$%v \n", k+":", v)
		inside: "%-25v ...$%v \n", k+":", v)
		k + ":" to put the : after the keys and not before the values
	*/
	return fs
}

//! Pointers *
//? Los punteros es una variable que nos permite acceder a la dirección de memoria, sirve para pasar una variable como referencia dentro de una función para modificar internamente esa variable al interior de esa función y que al exterior se conserve este valor modificado. EVITA CREAR COPIAS.
/*
* Operador de Dirreción => & 0xc0001
* Operador de Desreferenciación => *
e.g.
var myVar string = "original message"
var puntero *string = &myVar
*puntero = "original message updated"
he modificado el valor de myVar mediante el puntero
*/

// update tip
func (b *Bill) UpdateTip(tip float64) { //with the * pointer, I'm updating the values and keys inside that struct
	b.tip = tip //updating the tip value
}

// add an item to the bill
func (b *Bill) AddItem(name string, price float64) {
	b.items[name] = price
}

// b.items is a shorthand for (*b).items

// save bill
func (b *Bill) Save() {

	time.Sleep(10 * time.Second)

	data := []byte(b.Format())

	err := os.WriteFile("saved_bills/"+b.name+".txt", data, 0644)
	if err != nil { //si hay un error
		panic(err)
	}

	fmt.Println("bill was saved to file")
}
