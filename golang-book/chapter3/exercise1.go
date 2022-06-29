package main

import "fmt"

func main() {
	var x string
	x = "first"
	fmt.Println(x)
	x = x + "second"
	fmt.Println(x)
	x += "third" // result is 'firstthird'

	// creating variables with a starter value
	name := "Golang" // no need to add the type and the var keyword

}
