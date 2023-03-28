package main

import (
	"fmt"
	"strings"
)

type Product interface {
	GetPrice() float64
}

type SmallProduct struct {
	Price float64
}

type MediumProduct struct {
	Price float64
}

type LargeProduct struct {
	Price float64
}

func (s *SmallProduct) GetPrice() float64 {
	return s.Price
}

func (m *MediumProduct) GetPrice() float64 {
	return m.Price * 1.03
}

func (l *LargeProduct) GetPrice() float64 {
	return (l.Price * 1.06) + 2500
}

func createProduct(typeP string, price float64) Product {
	switch strings.ToLower(typeP) {
	case "s":
		return &SmallProduct{price}
	case "m":
		return &MediumProduct{price}
	case "l":
		return &LargeProduct{price}
	default:
		return nil
	}
}

func main() {

	var tipo string
	var precio float64

	fmt.Println("Ingrese el tipo de producto que desea (s: pequeño, m: mediano, l: grande)")
	fmt.Scanln(&tipo)

	fmt.Println("Ingrese el precio del producto: ")
	fmt.Scanln(&precio)

	p1 := createProduct(tipo, precio)
	fmt.Println(p1.GetPrice())
}
