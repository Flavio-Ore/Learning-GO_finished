package main

import "fmt"

func main() {

	//* Maps allows us to store key values pairs where the keys can be deifferent types and the values the same

	//----------------------------------------------------------------------

	//creating a map
	//? keys can be of multiple different types (str, int, float). But all the keys inside in a single map must have the same type and all the values must have the same length

	menu := map[string]float64{ //? specify the type of keys [string] and outside what type of value is gonna be
		// string := float64
		"canned food": 9.99,
		"cookies":     2.99,
		"salad":       6.99,
	}

	fmt.Println(menu)          //the complete map
	fmt.Println(menu["salad"]) //get the value of any of the items

	//----------------------------------------------------------------------

	// looping maps

	for key, value := range menu { //key="strs" and value="float64"
		fmt.Println(key, "-", value)
	}

	for key, value := range menu { //key="int" and value="str"}
		if value > 5 {
			fmt.Printf("This %v is expensive for you\n", key)
		} else {
			fmt.Printf("You can buy this %v\n", key)
		}
	}

	//----------------------------------------------------------------------
	//changing the values of the maps
	classroom := map[int]string{
		201: "Maths",
		202: "Algebra",
		203: "Vectors",
		204: "Stadistics",
	}

	fmt.Println(classroom)
	fmt.Println(classroom[201])

	//? classroom[202] = 100 //this is an error, the value must be a string
	classroom[201] = "Matemáticas" //TODO Replacing the values of the keys
	classroom[202] = "Álgebra"
	classroom[203] = "Vectores"
	classroom[204] = "Estadísticas"

	fmt.Println(classroom)
	fmt.Println(classroom[201])

	//----------------------------------------------------------------------

	// Nested Maps
	// Nested Maps are those in which values in key:value pairs of the outer map are also maps. In other words maps inside a map, hence nested maps.

	var nestedMap = map[string]map[string]string{} //main map

	nestedMap["fruits"] = map[string]string{} //fruits key
	nestedMap["colors"] = map[string]string{} //colors key

	nestedMap["fruits"]["a"] = "apple" //fruits values
	nestedMap["fruits"]["b"] = "banana"

	nestedMap["colors"]["r"] = "red" //colors values
	nestedMap["colors"]["b"] = "blue"

	fmt.Println(nestedMap)

}
