package main

import (
	"fmt"
	"strings"
)

/*
? 1- we cant store a multi return functions return in one value
? 2- how append works in action
? 3- handling different situations in return of a function
? 4- return of Split is unexpected and we cant set an explicit array for storing return of it
*/

func getInitials(n string) (string, string) { //first value string and second string
	s := strings.ToUpper(n)

	names := strings.Split(s, " ")

	var initials []string

	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}

func main() {

	firstName1, secondName1 := getInitials("Gon Stuart")

	fmt.Println(firstName1, secondName1)

	fn2, sn2 := getInitials("no uppercases")

	fmt.Println(fn2, sn2)

	fn3, sn3 := getInitials("SoloGon")

	fmt.Println(fn3, sn3)

	fn4, sn4 := getInitials("here")

	fmt.Println(fn4, sn4)
}
