package helper

import "fmt"

func SayHello() {
	fmt.Println("SayHello From Package helper has been called")
}

var Version = "2.9"

// untuk export harus menggunakan kapital di awal nama variabel
