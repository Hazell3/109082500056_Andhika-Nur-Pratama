package main

import "fmt"

func tanggunganHari(jumlahHari int, tujuan string) int {
	if tujuan == "domestik" {
		if jumlahHari > 3 {
			return 3
		}
		return jumlahHari
	} else {
		if jumlahHari > 8 {
			return 8
		}
		return jumlahHari
	}
}

func biayaPerHari(jumlahMhs int) int {
	makan := 2 * 35000
	inap := 250000
	saku := 300000
	return jumlahMhs * (makan + inap + saku)
}

func perhitunganBiaya(jumlahMhs, lamaPerjalanan int, tujuan string, totalBiaya *float64) {
	hari := tanggunganHari(lamaPerjalanan, tujuan)
	biayaDomestik := biayaPerHari(jumlahMhs) * hari
	if tujuan == "domestik" {
		*totalBiaya = float64(biayaDomestik)
	} else {
		*totalBiaya = float64(biayaDomestik) * 1.5
	}
}

func main() {
	var jumlah, lama int
	var tujuan string
	var biaya float64

	fmt.Print("masukkan jumlah mahasiswa : ")
	fmt.Scan(&jumlah)
	fmt.Print("Masukkan lama hari study tour : ")
	fmt.Scan(&lama)
	fmt.Print("Masukkan tujuan study tour (domestik/mancanegara) : ")
	fmt.Scan(&tujuan)

	perhitunganBiaya(jumlah, lama, tujuan, &biaya)

	fmt.Printf("Biaya perjalanan yang harus dikeluarkan Tel-U : Rp. %.0f\n", biaya)
}
