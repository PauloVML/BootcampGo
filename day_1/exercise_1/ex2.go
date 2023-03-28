package main

import "fmt"

func main() {

	// nombre := "Paulo"

	// fmt.Println(len(nombre))

	// transform(nombre)

	usingSwitch()

}

func transform(nombre string) {
	for _, letra := range nombre {
		fmt.Println(string(letra))
	}
}

func usingSwitch() {

	numero := 15

	switch {
	case numero == 15:
		fmt.Println(numero)
		fallthrough
	case numero == 16:
		fmt.Println("no evalúa")
		fallthrough
	case numero == 17:
		fmt.Println("17 no evalúa")
	}

}
