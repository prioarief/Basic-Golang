package main

import "fmt"

type student struct {
	name  string
	grade int
}

type class struct {
	className string
	student
}

func main() {
	// fmt.Println("Hi from struct")
	// deklarasi struct

	// var s1 student // inisialisasi 1
	// var s1 = student{} // inisialisasi 2
	// s1.name = "Prio"
	// s1.grade = 18

	// var s1 = student{"Prio", 18} // inisialisasi 3
	// fmt.Println(s1)

	// Object Pointer
	// var s1 = student{name: "Prio", grade: 18}
	// var s2 *student = &s1

	// fmt.Println("student 1, name :", s1.name)
	// fmt.Println("student 4, name :", s2.name)

	// s2.name = "Arief"
	// fmt.Println("student 1, name :", s1.name)
	// fmt.Println("student 4, name :", s2.name)

	// embedded struct
	var s1 = class{}
	s1.className = "RPL"
	s1.name = "Prio"
	s1.grade = 18
	fmt.Println(s1.student.name)
	fmt.Println(s1.name)
	fmt.Println(s1.className)
	fmt.Println(s1.grade)

}
