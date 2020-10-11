package main

import (
	"fmt"
)

func main() {
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("Angka", i)
	// }

	// // JUST CONDITION
	// var i = 0

	// for i < 5 {
	// 	fmt.Println("Angka", i)
	// 	i++
	// }

	// // WITHOUT ARGUMENT
	// var i = 0

	// for {
	// 	fmt.Println("Angka", i)

	// 	i++
	// 	if i == 5 {
	// 		break
	// 	}
	// }

	// USING LABEL
outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}
}
