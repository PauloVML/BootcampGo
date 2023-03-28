package main

import "fmt"

// Crear un programa que cumpla los siguiente puntos:
// Tener una estructura llamada Product con los campos ID, Name, Price, Description y Category.
// Tener un slice global de Product llamado Products instanciado con valores.
// 2 métodos asociados a la estructura Product: Save(), GetAll(). El método Save() deberá tomar el slice de Products
// y añadir el producto desde el cual se llama al método. El método GetAll() deberá imprimir
// todos los productos guardados en el slice Products.
// Una función getById() al cual se le deberá pasar un INT como parámetro y
// retorna el producto correspondiente al parámetro pasado.
// Ejecutar al menos una vez cada método y función definido desde main().

type Product struct {
	ID          int
	Name        string
	Price       int
	Description string
	Category    string
}

var Products = []Product{{
	ID:          1,
	Name:        "Taza",
	Price:       20,
	Description: "Taza",
	Category:    "Cocina",
}, {
	ID:          2,
	Name:        "Mac",
	Price:       1500,
	Description: "Computadora",
	Category:    "Tecnología",
}}

func (p *Product) save() {
	Products = append(Products, *p)
}

func (p *Product) getAll() {
	for _, prod := range Products {
		fmt.Println(prod)
	}
}

func getById(id int) Product {
	for _, producto := range Products {
		if producto.ID == id {
			return producto
		}
	}
	return Product{}
}

func main() {

	newProduct := Product{
		ID:          3,
		Name:        "Mouse",
		Price:       75,
		Description: "Raton",
		Category:    "Tecnología",
	}

	newProduct.save()
	newProduct.getAll()

	np := getById(3)
	fmt.Println(np)

}
