package main

import (
	"fmt"
)

func main() {

	// iterating over dictionaries (maps)
	// uniqueNames := map[string]bool{"Dan": true, "bingi": true, "karim": true}

	// for _, v := range uniqueNames {
	// 	fmt.Println(v)
	// }

	// m := map[string]int{
	// 	"a": 1,
	// 	"c": 3,
	// 	"b": 2,
	// }
	// for t := 0; t < 3; t++ {
	// 	fmt.Println(t)
	// 	for k, v := range m {
	// 		fmt.Println(k, v)
	// 	}
	// }

	// iterating over strings

	sample := []string{"foo", "bar", "baz", "line", "live"}
	// for i, v := range sample {
	// 	// i=Index v=Value
	// 	fmt.Println(i, v)
	// }
outer:
	for _, v := range sample {
		// this ignores the index and returns the value.
		for i, r := range sample {
			fmt.Println(i, r, string(r))
			if r == "l" {
				continue outer
			}
		}
		fmt.Println(v)
	}
	// for i := range sample {
	// 	// this ignores the value and returns the index.
	// 	fmt.Println(i)
	// }

}
