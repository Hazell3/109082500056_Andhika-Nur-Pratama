package main

import "fmt"

func pangkat(a, b int) int {
	if b == 0 {
		return 1
	}
	return a * pangkat(a, b-1)
}

func main() {
	var a, b int
	fmt.Print("Masukkan Nilai a dan b: ")
	fmt.Scan(&a, &b)

	hasil := pangkat(a, b)
	fmt.Println("Hasil", hasil)
}
