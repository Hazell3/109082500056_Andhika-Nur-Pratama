package main

import "fmt"

func main() {
	var gram int

	fmt.Print("Masukkan Berat (Kg): ")
	fmt.Scan(&gram)

	kg := gram / 1000
	sisa := gram % 1000

	biayaKg := kg * 10000
	biayaSisa := 0

	if kg > 10 {
		biayaSisa = 0
	} else {
		if sisa >= 500 {
			biayaSisa = sisa * 5
		} else {
			biayaSisa = sisa * 15
		}
	}
	total := biayaKg + biayaSisa

	fmt.Println("\n===== Detail Perhitungan =====")
	fmt.Println("Detail Berat: ", kg, "Kg +", sisa, "gram")
	fmt.Println("Detail Biaya: Rp.", biayaKg, "+ Rp.", biayaSisa)
	fmt.Println("Total Biaya: Rp", total)
}
