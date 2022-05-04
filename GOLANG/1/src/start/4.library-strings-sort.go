package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	//------------------------------------------------------------------------
	stringValue, anotherString := "This is a string value", "This is the value of the other string" // var stringValue String = "string"

	//------------------------------------------------------------------------
	fmt.Println(strings.Contains(stringValue, "value")) //true

	//------------------------------------------------------------------------

	fmt.Println(strings.ReplaceAll(stringValue, stringValue, "This is a string that changed but not afectted the ORIGINAL stringValue"))
	//? the changed

	fmt.Println(strings.ReplaceAll(stringValue, "value", "yet"))
	//? spicified change

	fmt.Println(stringValue, anotherString)
	//------------------------------------------------------------------------

	fmt.Println(strings.Index(stringValue, "string"))
	//------------------------------------------------------------------------

	fmt.Println(strings.Split(stringValue, " "))
	//? we convert the string to an Array

	//----------------------------------------------------------------
	//----------------------------------------------------------------

	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}

	sort.Ints(ages) //* this method has altered the original Array ages
	fmt.Println(ages)

	index := sort.SearchInts(ages, 90) //* 90 doesn't exist in ages array
	fmt.Println(index)                 //* Add it to the Array

	//----------------------------------------------------------------
	//----------------------------------------------------------------

	names := []string{"Juan", "Stuart", "Juniors", "Pepe"}

	sort.Strings(names) //* this method has altered the original
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "Pepe"))

	names = append(names, "extra", "otherextra")

	fmt.Println(names)

}
