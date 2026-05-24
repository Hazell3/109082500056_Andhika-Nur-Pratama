package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	var arr [100]int

	fmt.Print("Jumlah elemen: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Elemen ke-%d: ", i)
		fmt.Scan(&arr[i])
	}

	fmt.Println("\nIsi array:")
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}

	fmt.Println("\n\nIndeks ganjil:")
	for i := 1; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}

	fmt.Println("\n\nIndeks genap:")
	for i := 0; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}

	var x int
	fmt.Print("\n\nMasukkan kelipatan indeks: ")
	fmt.Scan(&x)

	fmt.Println("Elemen indeks kelipatan", x)
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}

	var hapus int
	fmt.Print("\n\nIndeks yang dihapus: ")
	fmt.Scan(&hapus)

	for i := hapus; i < n-1; i++ {
		arr[i] = arr[i+1]
	}
	n--

	fmt.Println("Array setelah dihapus:")
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}

	var total int
	for i := 0; i < n; i++ {
		total += arr[i]
	}

	rata := float64(total) / float64(n)
	fmt.Println("\n\nRata-rata =", rata)

	var jumlah float64

	for i := 0; i < n; i++ {
		jumlah += math.Pow(float64(arr[i])-rata, 2)
	}

	std := math.Sqrt(jumlah / float64(n))
	fmt.Println("Standar deviasi =", std)

	var cari, frek int
	fmt.Print("Bilangan yang dicari: ")
	fmt.Scan(&cari)

	for i := 0; i < n; i++ {
		if arr[i] == cari {
			frek++
		}
	}

	fmt.Println("Frekuensi =", frek)
}
