package main

import "fmt"

func main() {
	//Integers
	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("1 + 1 =", 1.0+1.0)

	//Strings
	fmt.Println("\n_______Strings_________\n")

	fmt.Println(len("Hello, world"))
	fmt.Println("Hello, World"[1]) // will be display in bytes '101' and not the letter 'e'
	fmt.Println("Hello, " + "World")

	// Booleans
	fmt.Println("\n ----- Booleans -----")
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(false || true)
	fmt.Println(!false)
	fmt.Println(!true)
}
