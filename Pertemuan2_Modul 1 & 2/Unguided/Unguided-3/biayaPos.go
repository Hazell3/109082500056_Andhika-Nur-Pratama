package main

import "fmt"

func main() {
	var gram int

	fmt.Print("Masukkan Berat (Kg): ")
	fmt.Scan(&gram)

	kg := gram / 1000
	sisa := gram % 1000

	berat := kg * 10000
	ongkir := 0

	if kg > 10 {
		ongkir = 0
	} else {
		if sisa >= 500 {
			ongkir = sisa * 5
		} else {
			ongkir = sisa * 15
		}
	}
	total := berat + ongkir

	fmt.Println("\n===== Detail Perhitungan =====")
	fmt.Println("Detail Berat: ", kg, "Kg +", sisa, "gram")
	fmt.Println("Detail Biaya: Rp.", berat, "+ Rp.", ongkir)
	fmt.Println("Total Biaya: Rp", total)
}
