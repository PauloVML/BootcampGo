package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	ID          int
	Nombre      string
	DateOfBirth int
}

type Employee struct {
	ID       int
	Position string
	Person   Person
}

func (e *Employee) printEmployee() {
	JSON, _ := json.Marshal(*e)
	fmt.Println(string(JSON))
}

func main() {

	p := Person{
		ID:          1,
		Nombre:      "Paulo",
		DateOfBirth: 471995,
	}

	e := Employee{
		ID:       1,
		Position: "Gerente",
		Person:   p,
	}

	e.printEmployee()

}
