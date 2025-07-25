package main

import "fmt"

// types struct
type pet struct {
	name string
	age  int
}

func main() {

	fmt.Println("Structs in go")

	// using typed struct
	var Dog = pet{}
	Dog.name = "peace"
	Dog.age = 12

	// anonymous struct
	var animal = struct {
		name string
		age  int
	}{
		"cow",
		23,
	}

	fmt.Println(animal)

}
