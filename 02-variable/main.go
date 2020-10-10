package main

import "fmt"

func main() {
	// // use var
	// var name string = "Prio"
	// fmt.Printf("Halo %s \n", name)
	// fmt.Println("Halo" + name)

	// // use inference
	// lastName := "Arief"
	// lastName = "Gunawan"
	// fmt.Println(lastName)

	// // multi declatation
	// var name, age, isHandsome = "Prio", 18, true
	// var firstName, lastName, nickName string = "Prio Arief", "Gunawan", "Arief"
	// firstName, lastName, nickName := "Prio Arief", "Gunawan", "Arief"
	// fmt.Println(firstName, lastName, nickName)

	// // use new
	// name := new(string)
	// fmt.Println(name)  // result 0xc0000881e0
	// fmt.Println(*name) // result ""

	// use const for constanta
	const name string = "prio"
	fmt.Println(name)

}
