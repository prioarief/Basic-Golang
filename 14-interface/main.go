package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func (p Person) GetName() string {
	return p.Name
}

func (a Animal) GetName() string {
	return a.Name
}

func main() {
	var p1 = Person{Name: "Mbappe"}

	// p1 (Person) bisa di pass ke fungsi SayHello jika Person implement kontrak interface hasName yang mempunyai fungsi GetName
	// dan return valuenya string
	SayHello(p1)

	// cat (Animal) bisa di pass ke fungsi SayHello jika Animal implement kontrak interface hasName yang mempunyai fungsi GetName
	// dan return valuenya string
	var cat = Animal{Name: "Cat"}
	SayHello(cat)
}
