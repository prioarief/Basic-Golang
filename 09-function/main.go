package main

import (
	"fmt"
	"math"
)

func main() {
	// var names = []string{"John", "Wick"}
	// // printMessage("halo", names)
	// fmt.Println(time.Now().Unix())
	// fmt.Println(rand.Int())

	// fmt.Println(getValue(2, 10))
	// fmt.Println(getValue(2, 10))
	// fmt.Println(getValue(2, 10))

	// fmt.Println(calculate(5))
	fmt.Println(sumMin(5, 2))
}

// func printMessage(message string, arr []string) {
// 	var nameString = strings.Join(arr, " ") // concat array
// 	fmt.Println(message, nameString)
// }

// // FUNCTION WITH RETURN
// func getValue(min, max int) int {
// 	value := rand.Int()%(max-min+1) + min
// 	return value
// }

// FUNCTION MULTIPLE RETURN
func calculate(d float64) (float64, float64) {
	area := math.Pi * math.Pow(d/2, 2)
	circumference := math.Pi * d

	return area, circumference
}

// FUNCTION MULTIPLE RETURN WITH PREDEFINED VALUE

func sumMin(a, b int) (value1 int, value2 int) {
	value1 = a + b
	value2 = a - b

	return
}

// func nameOfFunc(paramA type, paramB type, paramC type) returnType
// func nameOfFunc(paramA, paramB, paramC type) returnType

// func randomWithRange(min int, max int) int
// func randomWithRange(min, max int) int
