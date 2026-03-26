package main

import "fmt"

func main() {

	var productNames = [4]string{"Laptop", "Smartphone", "Tablet", "Headphones"}

	prices := [4]float64{10.5, 20.0, 15.75, 30.0}
	fmt.Println(prices)

	productNames[2] = "Smartwatch"

	fmt.Println(productNames[2])

	featuredPrices := prices[1:]
	fmt.Println(featuredPrices)
}

type Product struct {
	title string
	id    string
	price float64
}

type TempratureDeata struct {
	day1 float64
	day2 float64
	day3 float64
	day4 float64
}
