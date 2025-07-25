package main

import "fmt"

func main() {

	fmt.Println("Maps in go....")

	var myMap = map[string][]string{
		"cola": []string{"drink", "caffein", "Bad for me"},
		"malt": []string{"good for me", "also a drink", "Vhim"},
	}

	myMap["new drink"] = myMap["cola"]

	fmt.Println("Map data for key malt: ", myMap["malt"])
	fmt.Println("Map data for key cola: ", myMap["cola"])
	fmt.Println(myMap["new drink"])
}
