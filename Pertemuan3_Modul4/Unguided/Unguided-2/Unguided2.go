package main

import (
	"fmt"
	"math"
)

func persegi(sisi int) {
	luas := sisi * sisi
	keliling := 4 * sisi

	fmt.Println("Luas Persegi: ", luas)
	fmt.Println("Keliling Persegi: ", keliling)

}

func persegiPanjang(p, l int) {
	luas := p * l
	keliling := 2 * (p + l)

	fmt.Println("Luas Persegi Panjang: ", luas)
	fmt.Println("Keliling Persegi Panjang: ", keliling)

}

func lingkaran(r float64) {
	luas := math.Pi * r * r
	keliling := 2 * math.Pi * r

	fmt.Println("Luas Lingkaran: ", luas)
	fmt.Println("Keliling Lingkaran: ", keliling)

}

func main() {

	var pilih int

	fmt.Println("PROGRAM BANGUN DATAR")
	fmt.Println("1. Hitung Luas dan Keliling Persegi")
	fmt.Println("2. Hitung Luas dan Keliling Persegi Panjang")
	fmt.Println("3. Hitung Luas dan Keliling Lingkaran")
	fmt.Print("Pilih: ")

	fmt.Scan(&pilih)

	switch pilih {

	case 1:
		var sisi int
		fmt.Print("Masukkan Sisi: ")
		fmt.Scan(&sisi)
		persegi(sisi)

	case 2:
		var p, l int
		fmt.Print("Masukkan Panjang: ")
		fmt.Scan(&p)
		fmt.Print("Masukkan Lebar: ")
		fmt.Scan(&l)
		persegiPanjang(p, l)

	case 3:
		var r float64
		fmt.Print("Jari-Jari Lingkaran: ")
		fmt.Scan(&r)
		lingkaran(r)
	}
}
