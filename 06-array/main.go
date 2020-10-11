package main

import (
	"fmt"
)

func main() {
	// var fruits [4]string // with 4 elements
	// fruits[0] = "Apple"
	// fruits[1] = "Grape"
	// fruits[2] = "Manggo"
	// fruits[3] = "Orange"

	// var fruits = [2]string{"Apple", "Orange"} // vertical
	// var fruits = [2]string{
	// 	"Apple", "Orange",
	// } // horizontal

	// WITHOUT COUNT ELEMENT
	var fruits = [...]string{"Apple", "Banana"}

	fmt.Println(fruits)
}
