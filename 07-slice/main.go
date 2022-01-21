package main

import (
	"fmt"
)

func main() {
	// Slice adalah reference elemen array. Slice bisa dibuat, atau bisa juga dihasilkan dari manipulasi sebuah array ataupun slice lainnya. Karena merupakan data reference, menjadikan perubahan data di tiap elemen slice akan berdampak pada slice lain yang memiliki alamat memori yang sama dan sebaliknya.

	// var fruitsA = []string{"apple", "grape"}     // slice
	// var fruitsB = [2]string{"banana", "melon"}   // array
	// var fruitsC = [...]string{"papaya", "grape"} // array

	var fruits = [...]string{"Apple", "Banana", "Orange", "Manggo"}
	var newFruitss = fruits[:2] // slice all elements
	// var newFruits = fruits[0:2] // slice from index 0 - before index 2
	// fmt.Println(newFruits) // [Apple Banana]
	fmt.Println(newFruitss)      // [Apple Banana]
	fmt.Println(cap(newFruitss)) // mendapatkan kapasitas dari slice tersebut (4)
	fmt.Println(len(newFruitss)) // mendapatkan length dari slice (2)

	// append slice
	var fruit2 = append(newFruitss, "new")
	fmt.Println(fruit2)
	fmt.Println(newFruitss)
	fmt.Println(fruits)

}
