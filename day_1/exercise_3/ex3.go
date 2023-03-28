package main

import "fmt"

func main() {

	// Solución 1

	// 	Realizar una aplicación que reciba  una variable con el número del mes.
	// Según el número, imprimir el mes que corresponda en texto.
	// ¿Se te ocurre que se puede resolver de distintas maneras? ¿Cuál elegirías y por qué?
	// Ej: 7, Julio.
	// Nota: Validar que el número del mes, sea correcto.

	// var mes int

	// fmt.Println("Ingrese el número de mes")
	// fmt.Scanln(&mes)

	// if mes <= 12 {
	// 	switch mes {
	// 	case 1:
	// 		fmt.Println("Enero")
	// 	case 2:
	// 		fmt.Println("Febrero")
	// 	case 3:
	// 		fmt.Println("Marzo")
	// 	case 4:
	// 		fmt.Println("Abril")
	// 	case 5:
	// 		fmt.Println("Mayo")
	// 	case 6:
	// 		fmt.Println("Junio")
	// 	case 7:
	// 		fmt.Println("Julio")
	// 	case 8:
	// 		fmt.Println("Agosto")
	// 	case 9:
	// 		fmt.Println("Septiembre")
	// 	case 10:
	// 		fmt.Println("Octubre")
	// 	case 11:
	// 		fmt.Println("Noviembre")
	// 	case 12:
	// 		fmt.Println("Diciembre")
	// 	}
	// }

	// Solución 2

	// meses := []string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto",
	// 	"Septiembre", "Octubre", "Noviembre", "Diciembre"}
	// var mes uint

	// fmt.Println("Ingrese el número del mes")
	// fmt.Scanln(&mes)

	// if mes <= 12 {
	// 	fmt.Println(meses[mes-1])
	// }

	// Solución 3

	var meses = map[uint]string{
		1: "Enero",
		2: "Febrero",
		3: "Marzo",
		4: "Abril",
		5: "Mayo",
		6: "Junio",
		7: "Julio",
	}

	var mes uint

	fmt.Println("Ingrese el número del mes, recuerde que debe ser anterior a agosto")
	fmt.Scanln(&mes)

	if mesName, ok := meses[mes]; ok {
		fmt.Println(mesName)
	} else {
		fmt.Println("Número de mes ingresado inválido")
	}

}
