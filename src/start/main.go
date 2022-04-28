// siempre un package main
package main

//?Package fmt implements formatted I/O with functions analogous to C's printf and scanf.

import "fmt"

//* si se usan más packages, se usan {}
//* Formatting Strings and printing messages to the console

//! el nombre de una func es el entry point de una aplicación
//! Cuando se acaba el código del programa principal se acaba el programa
func main() {
	//* Go no puede declarar variables sin usarlas
	fmt.Println("This is a print")
}

//?  fmt. : nombre del paquete
//? Println : identificador, llamado print line
