package main

import "fmt"

func main() {
	result := add(1, 2)
	fmt.Println(result)

	result2 := add(1.5, 2.5)
	fmt.Println(result2)

	result3 := add("1.4", "2.2")
	fmt.Println(result3)
}

func add[T int | float64 | string](a, b T) T {

	return a + b
}
