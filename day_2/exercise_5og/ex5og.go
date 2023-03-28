package main

import (
	"errors"
	"fmt"
)

// ------------------------------------------ Ejercicio 1 ------------------------------------------

func salaryTaxes(salary float32) float32 {
	var tax float32
	if salary > 50000 {
		tax = salary * 0.17
	} else if salary > 150000 {
		tax = salary * 0.27
	} else {
		return 0
	}
	return tax
}

// ------------------------------------------ Ejercicio 2 ------------------------------------------

func calculateGradesAvg(grades ...float64) float64 {
	sum := 0.0
	count := 0
	for _, grade := range grades {
		if grade < 0 {
			continue
		}
		sum += grade
		count++
	}
	if count == 0 {
		return 0
	}
	return sum / float64(count)
}

// ------------------------------------------ Ejercicio 3 ------------------------------------------

func calculateSalary(minutes int32, category string) float32 {
	hours := float32(minutes) / 60
	var salary float32

	switch category {
	case "A":
		salary = 3000 * hours
		salary += salary * 0.5
	case "B":
		salary = 1500 * hours
		salary += salary * 0.2
	case "C":
		salary = 1000 * hours
	default:
		salary = 0
	}

	return salary
}

// ------------------------------------------ Ejercicio 4 ------------------------------------------

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func getMinimum(numbers ...int32) float32 {
	min := numbers[0]
	for _, num := range numbers {
		if num < min {
			min = num
		}
	}
	return float32(min)
}

func getMaximum(numbers ...int32) float32 {
	max := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	return float32(max)
}

func getAverage(numbers ...int32) float32 {
	var sum float32
	for _, num := range numbers {
		sum += float32(num)
	}
	return sum / float32(len(numbers))
}

func getStatistics(option string) (func(numbers ...int32) float32, error) {
	switch option {
	case minimum:
		return getMinimum, nil
	case average:
		return getAverage, nil
	case maximum:
		return getMaximum, nil
	default:
		return nil, errors.New("La opción no existe.")
	}

}

// ------------------------------------------ Ejercicio 5 ------------------------------------------

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarantula"
)

const (
	dogFoodKg       = 10
	catFoodKg       = 5
	hamsterFoodKg   = 0.25
	tarantulaFoodKg = 0.15
)

func dogFunc(number uint16) float32 {
	return float32(dogFoodKg * number)
}

func catFunc(number uint16) float32 {
	return float32(catFoodKg * number)
}

func hamsterFunc(number uint16) float32 {
	return (hamsterFoodKg * float32(number))
}

func tarantulaFunc(number uint16) float32 {
	return (tarantulaFoodKg * float32(number))
}

func animal(animalName string) (func(number uint16) float32, string) {
	switch animalName {
	case dog:
		return dogFunc, ""
	case cat:
		return catFunc, ""
	case hamster:
		return hamsterFunc, ""
	case tarantula:
		return tarantulaFunc, ""
	default:
		return nil, "No existe el animal."
	}
}

// ------------------------------------------ Main ------------------------------------------

func main() {
	// Ejercicio 1
	fmt.Println("\n Ejercicio 1 \n", salaryTaxes(55000))

	//Ejercicio 2
	fmt.Println("\n Ejercicio 2 \n", calculateGradesAvg(4.5, 5.0, -2.0, 3.0))

	//Ejercicio 3
	fmt.Println("\n Ejercicio 3 \n", calculateSalary(60, "C"))

	//Ejercicio 4
	minFunc, _ := getStatistics(minimum)
	maxFunc, _ := getStatistics(maximum)
	avgFunc, _ := getStatistics(average)
	_, funcError := getStatistics("test")
	fmt.Println("\n Ejercicio 4 \n", minFunc(1, 2, 3, 4, 5), maxFunc(1, 2, 3, 4, 5), avgFunc(1, 2, 3, 4, 5), funcError)

	//Ejercicio 5
	getDogFunc, _ := animal(dog)
	getCatFunc, _ := animal(cat)
	getHamsterFunc, _ := animal(hamster)
	getTarantulaFunc, _ := animal(tarantula)
	_, funcAnimalError := animal("test")
	fmt.Println("\n Ejercicio 5 \n", getDogFunc(2), getCatFunc(3), getHamsterFunc(1), getTarantulaFunc(5), funcAnimalError)
}
