package main

import "fmt"

func main() {
	//-------------------------------------------------------------------------
	const baseCuadrado = 10
	areaAlCuadrado := baseCuadrado * baseCuadrado

	var aSprint = fmt.Sprintf("Area al cuadrado de %v", areaAlCuadrado)

	fmt.Println(aSprint)

	fmt.Printf("Esto es un %T \n", aSprint)
	//-------------------------------------------------------------------------
	const radiusOfSphere float64 = 15

	sphereArea := radiusOfSphere / 2

	fmt.Println(sphereArea)
	//-------------------------------------------------------------------------

}
