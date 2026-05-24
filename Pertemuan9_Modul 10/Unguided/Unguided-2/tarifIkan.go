package main

import "fmt"

const NMAX = 1000

type arrIkan [NMAX]float64

func hitungWadah(tab arrIkan, x, y int) {
	jumlahWadah := x / y
	if x%y != 0 {
		jumlahWadah++
	}

	totalPerWadah := make([]float64, jumlahWadah)

	for i := 0; i < x; i++ {
		wadah := i / y
		totalPerWadah[wadah] += tab[i]
	}

	for i := 0; i < jumlahWadah; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Printf("%.2f", totalPerWadah[i])
	}
	fmt.Println()

	totalSemua := 0.0
	for i := 0; i < jumlahWadah; i++ {
		totalSemua += totalPerWadah[i]
	}
	rataRata := totalSemua / float64(jumlahWadah)
	fmt.Printf("%.2f\n", rataRata)
}

func main() {
	var arr arrIkan
	var x, y int

	fmt.Print("Banyak ikan yang dijual: ")
	fmt.Scan(&x)
	fmt.Print("Banyak ikan per wadah: ")
	fmt.Scan(&y)

	for i := 0; i < x; i++ {
		fmt.Printf("Berat ikan ke-%d: ", i+1)
		fmt.Scan(&arr[i])
	}

	hitungWadah(arr, x, y)
}
