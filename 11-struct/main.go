package main

import (
	"fmt"
)

type student struct {
	name  string
	grade int
}

type newstudent struct {
	grade int
	person
}
type person struct {
	name string
	age  int
}

type alias_person = person

// type person_horizontal struct { name string; age int }

func main() {
	var s1 student
	s1.name = "Prio"

	var s2 = student{}
	s2.name = "Arief"

	// var s3 = student{"Gunawan"} return error
	var s3 = student{"Gunawan", 5}
	var s4 = student{name: "Prio Arief Gunawan"}

	fmt.Println(s1, "s1")
	fmt.Println(s2, "s2")
	fmt.Println(s3, "s3")
	fmt.Println(s4, "s4")

	var s5 = student{name: "John", grade: 10}
	var s6 *student = &s5

	s6.name = "Yoga"

	fmt.Println(s5, "s5")
	fmt.Println(s6, "s6")

	var s7 = newstudent{}
	s7.age = 21
	s7.name = "Mbappe"
	s7.grade = 10

	fmt.Println(s7)

	// anonymous struct = cocok untuk sekali pakai
	var s8 = struct {
		name string
		age  int
	}{
		name: "John",
		age:  20,
	}

	var s9 = struct {
		name string
		age  int
	}{}

	var s10 struct {
		name string
		age  int
	}

	s9.age = 20
	s9.name = "Messi"

	s10.age = 20
	s10.name = "Ronaldo"

	fmt.Println(s8)
	fmt.Println(s9)
	fmt.Println(s10)

	// slice with struct
	var students = []student{
		{name: "John", grade: 10},
		{name: "John", grade: 10},
		{name: "John", grade: 10},
	}

	for _, e := range students {
		fmt.Println("My name is", e.name, "I'm", e.grade)
	}

	var p1 = alias_person{name: "Saya", age: 22}
	fmt.Println(p1)
}
