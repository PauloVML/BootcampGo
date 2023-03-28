package main

import (
	"errors"
	"fmt"
)

func main() {

	// 	Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de calificaciones
	// 	de los/as estudiantes de un curso. Requieren calcular los valores mínimo, máximo y promedio de sus calificaciones.
	// 	Para eso, se solicita generar una función que indique qué tipo de cálculo se quiere realizar (mínimo, máximo o promedio)
	// 	y que devuelva otra función y un mensaje (en caso que el cálculo no esté definido)
	//  que se le pueda pasar una cantidad N de enteros y devuelva el cálculo que se indicó en la función anterior.

	const (
		minimo   = "minimo"
		maximo   = "maximo"
		promedio = "promedio"
	)

	var notas = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	maxFunc, _ := elegirOperacion(maximo)
	minFunc, _ := elegirOperacion(minimo)
	promedioFunc, _ := elegirOperacion(promedio)
	_, err4 := elegirOperacion("asdasdasd")

	fmt.Println(maxFunc(notas...))
	fmt.Println(minFunc(notas...))
	fmt.Println(promedioFunc(notas...))
	fmt.Println(err4)

}

func elegirOperacion(mensaje string) (func(...int) float64, error) {
	var funcion func(...int) float64 = nil
	switch mensaje {
	case "minimo":
		funcion = minimo
	case "maximo":
		funcion = maximo
	case "promedio":
		funcion = promedio
	default:
		return nil, errors.New("función no válida")
	}
	return funcion, nil
}

func minimo(notas ...int) float64 {
	var min = notas[0]
	for _, nota := range notas {
		if nota < min {
			min = nota
		}
	}
	return float64(min)
}

func maximo(notas ...int) float64 {
	var max = notas[0]
	for _, nota := range notas {
		if nota > max {
			max = nota
		}
	}
	return float64(max)
}

func promedio(notas ...int) float64 {
	var total int
	for _, nota := range notas {
		total += nota
	}
	return float64(total / len(notas))
}
