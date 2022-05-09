package main

import "fmt"

func main() {
	fmt.Println(1)
	fmt.Println(2)
	fmt.Println(3)
	fmt.Println(4)
	fmt.Println(5)
	fmt.Println(6)
	fmt.Println(7)
	fmt.Println(8)
	fmt.Println(9)
	fmt.Println(10)

	// The for statement
	fmt.Println("\nFor loop")
	i := 1
	for i <= 10 {
		println(i)
		i += 1
	}

	//another option to use 'for'
	fmt.Println("\nSecond for loop statement")
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// If statement
	fmt.Println("\nIf statements")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("\t", i, "even")
		} else if i%2 != 0 {
			fmt.Println("\t", i, "odd")
		} else {
			fmt.Println("Default value")
		}
	}

	//The switch statement
	fmt.Println("\nThe switch statement using a loop from 1 to 10")
	for i := 1; i <= 10; i++ {
		switch i {
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")
		case 5:
			fmt.Println("Five")
		case 9:
			fmt.Println("Nine")
		case 4:
			fmt.Println("Four")
		case 7:
			fmt.Println("Seven")
		case 6:
			fmt.Println("Six")
		case 8:
			fmt.Println("Eight")
		case 10:
			fmt.Println("Ten")
		default:
			fmt.Println("Unknown number")
		}
	}
}
