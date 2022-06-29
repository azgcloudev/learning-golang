package main

import "fmt"

/*
* scopes are defined by the brackets or scopes
* variables inside a bracket block scope can be access within the same block or nested blocks
* vartiables in a block cannot be access by other blocks
* insideVariable cannot be access by function outside, will provide a compile error
*/

outsideVariable := "I am outside"

func main() {
	insideVariable := "I am inside a bracket scope"
	fmt.Println(outsideVariable)
}

func outside() {
	fmt.Println(insideVariable) 	// compile error: chapter3\scopes.go:12:1: syntax error: non-declaration statement outside function body
}

