package main

import "fmt"

// convert Fahrenheit into Celsius
// (C = (F − 32) * 5/9)

func main() {
	var fah int
	celsius := ((fah - 32) * 5 / 9)

	// input data
	fmt.Scanf("%d", &fah)
	fmt.Println("The result is", celsius)
}
