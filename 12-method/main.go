package main

import (
	"fmt"
	"strings"
)

type student struct {
	name  string
	grade int
}

func (s student) sayHello() {
	fmt.Println("Hello", s.name)
}

func (s *student) changeNameWithPointer(name string) {
	s.name = name
}

func (s student) changeNameWithoutPointer(name string) {
	s.name = name
}

func (s student) getFirstName() {
	fmt.Println("My first name is", strings.Split(s.name, " ")[0])
}

func main() {
	var s1 = student{name: "Kylian Mbappe", grade: 12}
	var s2 = &student{name: "Kylian Mbappe", grade: 12}
	s1.sayHello()
	s1.getFirstName()

	s2.sayHello()

	s1.changeNameWithoutPointer("John")
	fmt.Println("changeNameWithoutPointer", s1.name)

	s1.changeNameWithPointer("Pessi")
	fmt.Println("changeNameWithPointer", s1.name)

	// ketika kita melakukan manipulasi nilai pada property lain yang masih satu struct,
	// nilai pada property tersebut akan di rubah pada reference nya

}
