package main

//! Capital letter = exported | func ExportedFunction
//! no Capital letter = imported | func importedFunction

import (
	"bill_app/bill"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

//!||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||

// //OBJECT STRUCTURE | can hold many types of values

// //* In GO we don't need to use classes, instead we can create custom types using the STRUCT, it will be a blueprint which describes a type of data for that object

// // OWN CUSTOM TYPE
// type Bill struct { // ↓ properties of sthis struc
// 	name  string             // name type string
// 	items map[string]float64 // a map type
// 	tip   float64            // tip type float64
// }

// // make new bills
// func NewBill(name string) Bill { // we assign the name value to the name properties of the bill struct

// 	var b = Bill{ // new object

// 		// initial build object
// 		name: name,

// 		items: map[string]float64{}, // gonna work with this items {"canned food": 9.99, "cookies": 2.99, "salad": 6.99}

// 		tip: 0,
// 	}

// 	//? This is a initial build object with its initial values
// 	return b // returning the b value
// }

// //* adding this parentheses and saying that we receive a bill into this function now we're limiting this function only to bill objects, so it can only be called from a build object now just doing this (b bill) ↓

// // format the bill
// func (b *Bill) Format() string { // format() is not associated with bill object

// 	fs := "Bill breakdown: \n" //the start of the formatting string

// 	//cycle through the items ↓
// 	var total float64 = 0 //add each item cost to the total, and output it

// 	// list items | list out the items inside the formatted string fs
// 	for k, v := range b.items { // b.items to call the items properties

// 		//we want to output the key value
// 		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)

// 		total += v //to adding the cost to the total
// 	}
// 	// add tip

// 	fs += fmt.Sprintf("%-25v ...$%v\n", "tip:", b.tip)

// 	// total
// 	fs += fmt.Sprintf("%-25v ...$%0.2f \n", "total:", total+b.tip)

// 	/*
// 		I add -25 to add 25 characters long and it's going to push the prices
// 		I made this: Sprintf("%-25v ...$%v \n", k+":", v)
// 		inside: "%-25v ...$%v \n", k+":", v)
// 		k + ":" to put the : after the keys and not before the values
// 	*/
// 	return fs
// }

// //! Pointers *
// //? Los punteros es una variable que nos permite acceder a la dirección de memoria, sirve para pasar una variable como referencia dentro de una función para modificar internamente esa variable al interior de esa función y que al exterior se conserve este valor modificado. EVITA CREAR COPIAS.
// /*
// * Operador de Dirreción => & 0xc0001
// * Operador de Desreferenciación => *
// e.g.
// var myVar string = "original message"
// var puntero *string = &myVar
// *puntero = "original message updated"
// he modificado el valor de myVar mediante el puntero
// */

// // update tip
// func (b *Bill) UpdateTip(tip float64) { //with the * pointer, I'm updating the values and keys inside that struct
// 	b.tip = tip //updating the tip value
// }

// // add an item to the bill
// func (b *Bill) AddItem(name string, price float64) {
// 	b.items[name] = price
// }

// // b.items is a shorthand for (*b).items

// // save bill
// func (b *Bill) Save() {
// 	data := []byte(b.Format())
// 	err := os.WriteFile("../saved_bills/"+b.name+".txt", data, 0644)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("bill was saved to file")
// }
//!||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
// func to don't call every time the createBill reader function. Is gonna be easeier in the future to call this function anytime we want to get some user input

var wg = sync.WaitGroup{}

func getInput(prompt string, r *bufio.Reader) (string, error) { //accepting as an argument a reader
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err

}

func createBill() bill.Bill {
	//using reader to read the information
	reader := bufio.NewReader(os.Stdin) //NewReader to read information
	//----------------------------------------
	// the func called getInput do same func as well as those lines ↓ ↓ ↓

	// fmt.Print("Create a new bill name: ")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)
	//----------------------------------------
	name, _ := getInput("Create a new bill name: ", reader)
	b := bill.NewBill(name)

	fmt.Println("Created the bill - ", name) //I wrote , name instead of b.name for a error

	return b
}

func promptOptions(b bill.Bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip) ", reader)

	// switch option to choose a option
	switch opt {
	// ask what option want to add
	case "a":

		fmt.Println("you chose add item")

		name, _ := getInput("Item name ", reader)
		price, _ := getInput("Item price ", reader)

		// parsing floats price string to float value
		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}
		b.AddItem(name, p)

		fmt.Println("item added - ", name, p)

		promptOptions(b)

	case "t":

		fmt.Println("you chose add tip")

		tip, _ := getInput("Enter tip amount ($)", reader)

		// parsing floats strconv
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}
		b.UpdateTip(t)

		fmt.Println("tip added - ", tip)

		promptOptions(b)

	case "s":
		b.Save()
		fmt.Println("bill was saved to file", b)

	default:
		fmt.Println("that was not a valid option...")

		promptOptions(b)
		// I called promptOptions(b) to recycle the function
	}

	// fmt.Println(opt) //to print it before the switch statement was created
}

func main() {
	//user input and switch statements-------------------------------------
	//! FINAL PRESENTATION BILL APPLICATION ↓
	mybill := createBill()

	promptOptions(mybill)

	// fmt.Println(mybill)

	//using reciever functions----------------------------------------------

	//mybill := bill.NewBill("Golang's bill") //this return the struct value

	//mybill.Format() associated with bill objects

	//mybill.Format() //to get the copy of the bill, like when we take in arguments into functions go creates a copy it does the same thing in receiver functions for he item we use it on for the object we use it on

	//using reciever functions with pointers--------------------------------

	//fmt.Println(mybill) // it print the Structured Class

	//fmt.Println(mybill.Format()) // Instead of printing the mybill we gonna print the formating string

	//using receiver functions with pointers and adding items-----------------

	// mybill.AddItem("canned food", 9.99)
	// mybill.AddItem("cookies", 2.99)
	// mybill.AddItem("salad", 6.99)

	// mybill.UpdateTip(10)

	// fmt.Println(mybill.Format())

}
