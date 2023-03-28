package main

import "fmt"

type Alumnos struct {
	Nombre   string
	Apellido string
	Dni      int
	Fecha    int
}

func (a *Alumnos) detalle() {
	fmt.Printf("Nombre: [%v]\nApellido: [%v] \nDNI: [%v]\nFecha: [%v]\n", a.Nombre, a.Apellido, a.Dni, a.Fecha)
}

func main() {

	// var a int = 19
	// var b *int
	// b = &a
	// fmt.Printf("La variable b hace referencia a la DIRECCIÓN DE PUNTERO %v \n ", b)
	// fmt.Printf("La variable b hace referencia al VALOR %d \n", *b)

	a := Alumnos{
		Nombre:   "Paulo",
		Apellido: "Vergara",
		Dni:      38813337,
		Fecha:    4071995,
	}

	a.detalle()

}
