package main

import "fmt"

// defer adalah fungsi yang tetap dijalankan meskipun terjadi error (di eksekusi terakhir)

// panic adalah fungsi yang dipanggil ketika terjadi error pada saat program berjalan
// saat panic dipanggil program akan terhento namun defer tetap akan dieksekusi

// recover adalah fungsi yang digunakan untuk menangkap data panic
// dengan recover proses panic akan berhenti

func logging() {
	// tangkep error
	message := recover()
	if message != nil {
		fmt.Println("Terjadi kesalahan", message)
	}

	fmt.Println("logged")
}

func runApp(error bool) {
	defer logging()

	if error {
		// set panic
		panic("ERROR")
	}

	fmt.Println("Aplikasi Berjalan")
}

func main() {
	runApp(true)
	// runApp(false)
}
