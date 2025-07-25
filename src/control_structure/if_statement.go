package main

import (
	"fmt"
	"math/rand"
)

// BLOCKS
// Each place where a declaration occurs is called a block

// SHADOWING VARIABLES
func main() {

	// A shadowing variable is a variable that has the same name as a variable in a contain‐
	// ing block. For as long as the shadowing variable exists, you cannot access a shadowed
	// variable.

	var x int = 10

	if x > 5 {
		var x int = 5
		fmt.Println("value if x in the if block: ", x)
	}
	fmt.Println("value of x outside the if block: ", x)

	// random number in an if block
	n := rand.Intn(10)

	if n == 10 {
		fmt.Println("the value of the random number is: ", n)
	} else if n > 5 {
		fmt.Println("the value of the random number is greater than 5: ", n)
	} else {
		fmt.Println("the random value is not 10 and not greater than 5: ", n)
	}

	// scope condition variable
	if x := rand.Intn(20); x == 1 {

		fmt.Println("the value of x is ", x)
	} else {
		fmt.Println("the variable x is scoped to the if condition")
	}
}
