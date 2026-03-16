package main

import "fmt"

func hitungBiaya(jenis string, masuk, keluar int) int {

	jam := keluar - masuk
	tarif := 0

	if jenis == "motor" {
		if masuk < 17 {
			tarif = 4000
		} else {
			tarif = 5000
		}
	} else if jenis == "mobil" {
		if masuk < 17 {
			tarif = 6000
		} else {
			tarif = 7000
		}
	}

	return jam * tarif
}

func main() {

	var jenis string
	var masuk, keluar int
	total := 0
	kendaraan := 1

	for {
		fmt.Printf("\n*Kendaraan %d\n", kendaraan)
		fmt.Print("Kendaraan apa? (motor/mobil/- untuk selesai): ")
		fmt.Scan(&jenis)

		if jenis == "-" {
			break
		}

		fmt.Print("Masukkan Jam Masuk Kendaraan (0-24): ")
		fmt.Scan(&masuk)

		fmt.Print("Masukkan Jam Keluar Kendaraan (0-24): ")
		fmt.Scan(&keluar)

		biaya := hitungBiaya(jenis, masuk, keluar)

		fmt.Println("Biaya parkir", jenis, kendaraan, ":", biaya)

		total += biaya
		kendaraan++
	}

	fmt.Println("\nTotal Pendapatan Parkir Hari Ini:", total)
}
