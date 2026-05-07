package main

import "fmt"

const nProv int = 10

type NamaProv [nProv]string
type PopProv [nProv]int
type TumbuhProv [nProv]float64

func InputData(prov *NamaProv, pop *PopProv, tumbuh *TumbuhProv) {
	var i int

	fmt.Println("=== Masukkan Nama Provinsi, Populasi Provinsi, Angka Pertumbuhan Provinsi ===")

	for i = 0; i < nProv; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		fmt.Scan(&prov[i], &pop[i], &tumbuh[i])
	}
}

func ProvinsiTercepat(tumbuh TumbuhProv) int {
	var i, idx int

	idx = 0

	for i = 1; i < nProv; i++ {
		if tumbuh[i] > tumbuh[idx] {
			idx = i
		}
	}

	return idx
}

func Prediksi(prov NamaProv, pop PopProv, tumbuh TumbuhProv) {
	var i int
	var prediksi float64

	fmt.Println()
	fmt.Println("=== Prediksi Jumlah Penduduk Tahun Depan Pada Provinsi Dengan Pertumbuhan Diatas 2% ===")

	for i = 0; i < nProv; i++ {
		if tumbuh[i] > 0.02 {
			prediksi = (tumbuh[i] + 1) * float64(pop[i])
			fmt.Println(prov[i], int(prediksi))
		}
	}
}

func IndeksProvinsi(prov NamaProv, nama string) int {
	var i int

	for i = 0; i < nProv; i++ {
		if prov[i] == nama {
			return i
		}
	}

	return -1
}

func main() {
	var prov NamaProv
	var pop PopProv
	var tumbuh TumbuhProv

	var cari string
	var idx int

	InputData(&prov, &pop, &tumbuh)

	fmt.Scan(&cari)

	idx = ProvinsiTercepat(tumbuh)

	fmt.Println()
	fmt.Println("Provinsi dengan angka pertumbuhan tercepat :", prov[idx])

	idx = IndeksProvinsi(prov, cari)

	fmt.Println()
	if idx != -1 {
		fmt.Println("Data provinsi yang dicari :", cari)
	} else {
		fmt.Println("Provinsi tidak ditemukan")
	}

	Prediksi(prov, pop, tumbuh)
}
