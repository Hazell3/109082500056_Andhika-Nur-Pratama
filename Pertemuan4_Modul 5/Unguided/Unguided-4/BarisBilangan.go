package main

import "fmt"

func main() {
	var n int
	fmt.Print("Maskkan Nilai: ")
	fmt.Scan(&n)
	pola(n)
}

func pola(n int) {
	if n == 1 {
		fmt.Print("1 ")
		return
	}
	fmt.Print(n, " ")
	pola(n - 1)
	fmt.Print(n, " ")
}
