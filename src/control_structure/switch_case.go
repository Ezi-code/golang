package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// names := []string{"foo", "bar", "buzz"}

	// for _, v := range names {
	// 	switch size := len(v); size {
	// 	case 1:
	// 		fmt.Println(v, "Word size is 1")
	// 	case 2:
	// 		fmt.Println(v, "word size is 2")
	// 	case 3:
	// 		fmt.Println(v, "word size is 3")
	// 	default:
	// 		fmt.Println(v, "is a long word")
	// 	}
	// }

	// switch case with a break statement
loop:
	for i := 0; i < 10; i++ {
		switch {
		case i%2 == 0:
			fmt.Println(i, "is even")
		case i%3 == 0:
			fmt.Println(i, "is divisible by 3 but not 2")
		case i%7 == 0:
			fmt.Println("exit the loop!")
			break loop
		default:
			fmt.Println(i, "is boring")
		}
	}

	// switch case with a goto statement
	a := rand.Intn(10)
	for a < 100 {
		if a%5 == 0 {
			goto done
		}
		a = a*2 + 1
	}

	fmt.Println("do something when the loop completes normally")
done:
	fmt.Println("do complicated stuff no matter why we left the loop")
	fmt.Println(a)

}
