package main

import "fmt"

func main() {
	// GO DOES NOT SUPPORT TERNARY OPERATION

	// // IF ELSE
	// value := 0

	// if value > 90 {
	// 	fmt.Println("Excelent")
	// } else if value > 85 && value <= 90 {
	// 	fmt.Println("Great")
	// } else {
	// 	fmt.Println("Try Again")
	// }

	// // SWITCH CASE 1 CONDITION
	// point := 70
	// switch point {
	// case 100:
	// 	fmt.Println("Excellent")

	// case 90:
	// 	fmt.Println("Great")

	// case 80:
	// 	fmt.Println("Not Bad")

	// default:
	// 	fmt.Println("Try Again, Never give up")
	// }

	// // SWITCH CASE MULTIPLE CONDITION
	// point := 30
	// switch point {
	// case 100:
	// 	fmt.Println("Excellent")

	// case 90:
	// 	fmt.Println("Great")

	// case 80, 70, 60:
	// 	fmt.Println("Not Bad")

	// default:
	// 	{
	// 		// COMBINATION WITH IF ELSE ON DEFAULT STATEMENT
	// 		if point == 50 {
	// 			fmt.Println("Try Again, Never give up")
	// 		} else {
	// 			fmt.Println("You Will be better")
	// 		}
	// 	}
	// }

	// SWITCH CASE WITH STYLE IF ELSE
	point := 90

	switch {
	case point == 90:
		{
			fmt.Println("Great")
		}
		fallthrough // force check next condition, althought next condition false

	case (point < 90 && point > 85):
		{
			fmt.Println("Awesome")
		}
	}
}
