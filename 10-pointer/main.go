package main

import (
	"fmt"
)

func main() {
	fmt.Println("Halo from pointer")

	// DEKLARASI POINTER
	// var number *int
	// var name *string

	var numberA int = 4
	var numberB *int = &numberA
	var numberC int = *numberB
	var numberD = &numberA

	fmt.Println(numberA)
	fmt.Println(*numberD)
	fmt.Println(numberB)
	fmt.Println(numberC)

	// variabel biasa bisa diambil nilai pointernya, caranya dengan menambahkan tanda ampersand (&) tepat sebelum nama variabel. Metode ini disebut dengan referencing.

	// Dan sebaliknya, nilai asli variabel pointer juga bisa diambil, dengan cara menambahkan tanda asterisk (*) tepat sebelum nama variabel. Metode ini disebut dengan dereferencing.

	// =============================================================
	// = mengambil nilai pointer dari variable asli menggunakan &  =
	// = mengambil nilai asli dari pointer menggunakan *		   =
	// =============================================================

	// fmt.Println("numberA (value)   :", numberA)  // 4
	// fmt.Println("numberA (address) :", &numberA) // 0xc20800a220

	// fmt.Println("numberB (value)   :", *numberB) // 4
	// fmt.Println("numberB (address) :", numberB)  // 0xc20800a220

	// Parameter Pointer
	// number := 4
	// fmt.Println("before", number)
	// change(&number, 10)
	// fmt.Println("after", number)
}

func change(original *int, value int) {
	*original = value
}
