package main

import "fmt"

// Scope
// scope defines who has access to a variable, scopes are defined by the curly {} brackets, also applies to nested brackets

// Constant variables --> they cannot be reassigned with a value (are inmutable)
const pi float64 = 3.1415

// ----- practice ----
//function that converts farenheit into celcious

func FtoC() {
	fmt.Print("Enter the Farenheit: ")
	var fValue float64
	fmt.Scanf("%f", &fValue)

	var cValue = (fValue - 32) * 5 / 9
	fmt.Println("The value in celcius is", cValue)
}

func main() {

	// variables 'x' with type string
	var x string = "Hello World"
	fmt.Println(x)

	// variable can bi initialize without a value
	var y string
	y = "Welcome to go!"
	fmt.Println(y)

	// concatenation using variables
	// x = x + y
	x += y
	fmt.Println(x)

	// Short version to create a variable with a value from the beginning
	name := "Aldair" // Go compiler automatically will assign the string type to the variable. Cannot be used without a variable value
	fmt.Println(name)

	// get input from console with fmt.Scanf()
	fmt.Print("Enter any number: ")
	var userNumber float64
	fmt.Scanf("%f\n", &userNumber)

	FtoC() // call the function to convert
}
