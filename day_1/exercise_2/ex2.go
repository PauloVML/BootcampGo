package main

import "fmt"

func main() {

	// Un banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos.
	//Para ello tiene ciertas reglas para saber a qué cliente se le puede otorgar.
	//Solo le otorga préstamos a clientes cuya edad sea mayor a 22 años, se encuentren empleados y tengan más de un año de antigüedad en su trabajo. Dentro de los préstamos que otorga no les cobrará interés a los que posean un sueldo superior a $100.000.
	// Es necesario realizar una aplicación que reciba  estas variables y que imprima un mensaje de acuerdo a cada caso.
	// Tip: tu código tiene que poder imprimir al menos 3 mensajes diferentes.

	var edad int
	var empleado bool
	var antiguedad int
	var sueldo float64

	fmt.Println("Ingrese su edad:")
	fmt.Scanln(&edad)

	fmt.Println("¿Se encuentra empleado? (true/false):")
	fmt.Scanln(&empleado)

	fmt.Println("Ingrese la cantidad de años de antigüedad en su trabajo:")
	fmt.Scanln(&antiguedad)

	fmt.Println("Ingrese su sueldo:")
	fmt.Scanln(&sueldo)

	if edad > 22 && empleado && antiguedad > 1 {
		if sueldo > 100000 {
			fmt.Println("¡Felicidades! Usted cumple con los requisitos para acceder a un préstamo sin interés.")
		} else {
			fmt.Println("Usted cumple con los requisitos para acceder a un préstamo con interés.")
		}
	} else {
		fmt.Println("Lo sentimos, usted no cumple con los requisitos para acceder a un préstamo.")
	}

}
