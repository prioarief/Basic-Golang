package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {
	// name, hobby := greeting("Prio Arief Gunawan", "Traveling")
	// fmt.Printf("my name is %v and my hobby is %v\n", name, hobby)
	// var names = []string{"John", "Wick"}
	// // printMessage("halo", names)
	// fmt.Println(time.Now().Unix())
	// fmt.Println(rand.Int())

	// fmt.Println(getValue(2, 10))
	// fmt.Println(getValue(2, 10))
	// fmt.Println(getValue(2, 10))

	// fmt.Println(calculate(5))
	// fmt.Println(sumMin(5, 2))
	// fmt.Println(variadicExample(2, 3, 4, 5, 6, 7, 8, 9, 10))

	// FUNC CLOSURE
	// var getMinMax = func(n []int) (int, int) {
	// 	var min, max int
	// 	for i, e := range n {
	// 		switch {
	// 		case i == 0:
	// 			max, min = e, e
	// 		case e > max:
	// 			max = e
	// 		case e < min:
	// 			min = e
	// 		}
	// 	}
	// 	return min, max
	// }

	// var numbers = []int{2, 3, 4, 3, 4, 2, 3}
	// var min, max = getMinMax(numbers)
	// fmt.Printf("data : %v\nmin  : %v\nmax  : %v\n", numbers, min, max)

	// Immediately-Invoked Function Expression (IIFE) == langsung di eksekusi
	// var numbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}

	// var newNumbers = func(min int) []int {
	// 	var r []int
	// 	for _, e := range numbers {
	// 		if e < min {
	// 			continue
	// 		}
	// 		r = append(r, e)
	// 	}
	// 	return r
	// }(3)

	// fmt.Println("original number :", numbers)
	// fmt.Println("filtered number :", newNumbers)
	// variadic := variadicExample(2, 2, 2, 2)
	// fmt.Println(variadic)

	// numbers := []int{2, 3, 4, 3, 2, 5}
	// max := 3
	// var dataLength, getNumbers = findMax(numbers, max)
	// theNumbers := getNumbers()
	// fmt.Println(dataLength)
	// fmt.Println(theNumbers)
	// min, max := getMinMax(numbers)
	// fmt.Println(min, max)

	data := []string{"wick", "jason", "ethan", "arief", "prio"}
	dataContainsO := filterData(data, func(each string) bool {
		return strings.Contains(each, "o")
	})

	dataWithLength5 := filterData(data, func(each string) bool {
		return len(each) == 5
	})

	fmt.Println(reflect.TypeOf(data))
	fmt.Println(dataContainsO)
	fmt.Println(dataWithLength5)

	name := "Prio Arief Gunawan"
	fmt.Println(reflect.TypeOf(name).String() == "string")
	// fmt.Println(strings.ToUpper(name))
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

func greeting(name, hobby string) (myname string, myhobby string) {
	myname = name
	myhobby = hobby
	return myname, myhobby
}

// func nameOfFunc(paramA type, paramB type, paramC type) returnType
// func nameOfFunc(paramA, paramB, paramC type) returnType

// func randomWithRange(min int, max int) int
// func randomWithRange(min, max int) int

// func variadic = params no limits
func variadicExample(value ...int) float64 {
	var total = 0
	for _, number := range value {
		total += number
	}

	avg := float64(total) / float64(len(value))
	return avg
}

// closure = anonymous function
var getMinMax = func(n []int) (int, int) {
	var min, max int
	for i, e := range n {
		switch {
		case i == 0:
			max, min = e, e

		case e > max:
			max = e

		case e < min:
			min = e
		}
	}

	return min, max
}

// closure as return function
func findMax(numbers []int, max int) (int, func() []int) {
	var res []int
	for _, e := range numbers {
		if e <= max {
			res = append(res, e)
		}
	}
	return len(res), func() []int {
		return res
	}
}

// function as parameter
func filterData(data []string, callback func(string) bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}

	return result
}
