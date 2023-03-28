package main

import "fmt"

func main() {

	// Un empleado de una empresa quiere saber el nombre y edad de uno
	// de sus empleados. Según el siguiente mapa, ayudá a imprimir la edad de Benjamin.

	// var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	// Por otro lado también es necesario:
	// Saber cuántos de sus empleados son mayores de 21 años.
	// Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
	// Eliminar a Pedro del mapa.

	// Solución 1

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	var mayores int
	var nuevoEmpleado string
	var edadNuevoEmpleado int

	fmt.Println(employees["Benjamin"])

	for key, value := range employees {
		if value > 21 {
			mayores++
			fmt.Println(key + " es mayor de 21 años")
		}
	}

	fmt.Println("En total ", mayores, " empleados son mayores de 21 años")

	fmt.Println("Ingrese el nombre del nuevo empleado")
	fmt.Scanln(&nuevoEmpleado)

	fmt.Println("Ingrese la edad del nuevo empleado")
	fmt.Scanln(&edadNuevoEmpleado)

	employees[nuevoEmpleado] = edadNuevoEmpleado
	delete(employees, "Pedro")

	fmt.Println("De esta manera la nómina de empleados quedaría de la siguiente manera: ", employees)
	// for k, v := range employees {
	// 	fmt.Println(v, " de ", k, " años de edad")
	// }

}
