package main

import (
	"fmt"
)

func main() {
	value := 0

	if value > 90 {
		fmt.Println("Excelent")
	} else if value > 85 && value <= 90 {
		fmt.Println("Great")
	} else {
		fmt.Println("Stupid")
	}
}
