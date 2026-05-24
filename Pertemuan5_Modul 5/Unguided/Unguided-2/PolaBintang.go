package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan Nilai: ")
	fmt.Scan(&n)
	bintang(n, 1)
}

func bintang(n int, baris int) {
	if baris > n {
		return
	}

	for i := 1; i <= baris; i++ {
		fmt.Print("*")
	}
	fmt.Println()
	bintang(n, baris+1)
}
