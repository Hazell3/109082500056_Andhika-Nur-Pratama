package main

import "fmt"

const NMAX = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku [NMAX]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n *int) {
	var i int

	fmt.Scan(n)

	for i = 0; i < *n; i++ {
		fmt.Scan(
			&pustaka[i].id,
			&pustaka[i].judul,
			&pustaka[i].penulis,
			&pustaka[i].penerbit,
			&pustaka[i].eksemplar,
			&pustaka[i].tahun,
			&pustaka[i].rating,
		)
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	var idxMax, i int

	idxMax = 0

	for i = 1; i < n; i++ {
		if pustaka[i].rating > pustaka[idxMax].rating {
			idxMax = i
		}
	}

	fmt.Println("=== Buku Terfavorit ===")
	fmt.Println("Judul    :", pustaka[idxMax].judul)
	fmt.Println("Penulis  :", pustaka[idxMax].penulis)
	fmt.Println("Penerbit :", pustaka[idxMax].penerbit)
	fmt.Println("Tahun    :", pustaka[idxMax].tahun)
	fmt.Println()
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	var temp Buku
	var i, j int

	for i = 1; i < n; i++ {

		temp = pustaka[i]
		j = i - 1

		for j >= 0 && temp.rating > pustaka[j].rating {
			pustaka[j+1] = pustaka[j]
			j--
		}

		pustaka[j+1] = temp
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	var batas int

	batas = 5
	if n < 5 {
		batas = n
	}

	fmt.Println("=== 5 Buku Rating Tertinggi ===")

	for i := 0; i < batas; i++ {
		fmt.Printf("%d. %s (Rating %d)\n",
			i+1,
			pustaka[i].judul,
			pustaka[i].rating)
	}

	fmt.Println()
}

func CariBuku(pustaka DaftarBuku, n, r int) {
	var kiri, kanan, tengah int
	var ketemu bool

	kiri = 0
	kanan = n - 1
	ketemu = false

	for kiri <= kanan && !ketemu {

		tengah = (kiri + kanan) / 2

		if pustaka[tengah].rating == r {

			fmt.Println("=== Buku Ditemukan ===")
			fmt.Println("ID        :", pustaka[tengah].id)
			fmt.Println("Judul     :", pustaka[tengah].judul)
			fmt.Println("Penulis   :", pustaka[tengah].penulis)
			fmt.Println("Penerbit  :", pustaka[tengah].penerbit)
			fmt.Println("Eksemplar :", pustaka[tengah].eksemplar)
			fmt.Println("Tahun     :", pustaka[tengah].tahun)
			fmt.Println("Rating    :", pustaka[tengah].rating)

			ketemu = true

		} else if r > pustaka[tengah].rating {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}

	if !ketemu {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}

func main() {
	var pustaka DaftarBuku
	var n, ratingCari int

	DaftarkanBuku(&pustaka, &n)

	CetakTerfavorit(pustaka, n)

	UrutBuku(&pustaka, n)

	Cetak5Terbaru(pustaka, n)

	fmt.Scan(&ratingCari)

	CariBuku(pustaka, n, ratingCari)
}
