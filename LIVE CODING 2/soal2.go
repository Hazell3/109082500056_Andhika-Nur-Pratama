package main

import "fmt"

const NMAX int = 51

type Mahasiswa struct {
	NIM   string
	Nama  string
	nilai int
}

type arrayMahasiswa [NMAX]Mahasiswa

func inputData(A *arrayMahasiswa, n int) {
	var i int

	for i = 0; i < n; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		fmt.Scan(&A[i].NIM, &A[i].Nama, &A[i].nilai)
	}
}

func cariNilai1(A arrayMahasiswa, n int, NIM string) int {
	var i int

	for i = 0; i < n; i++ {
		if A[i].NIM == NIM {
			return A[i].nilai
		}
	}
	return -1
}

func cariNilaiTerbesar(A arrayMahasiswa, n int, NIM string) int {
	var i, max int
	var found bool

	found = false
	max = -1

	for i = 0; i < n; i++ {
		if A[i].NIM == NIM {
			if !found || A[i].nilai > max {
				max = A[i].nilai
				found = true
			}
		}
	}
	return max
}

func main() {
	var A arrayMahasiswa
	var n int
	var nimCari string
	var nilai1, nilaiTerbesar int

	fmt.Print("Masukkan jumlah data: ")
	fmt.Scan(&n)

	inputData(&A, n)

	fmt.Print("Masukkan NIM yang ingin dicari: ")
	fmt.Scan(&nimCari)

	nilai1 = cariNilai1(A, n, nimCari)
	nilaiTerbesar = cariNilaiTerbesar(A, n, nimCari)

	if nilai1 != -1 {
		fmt.Println("Nilai pertama dari NIM", nimCari, "adalah: ", nilai1)
		fmt.Println("Nilai terbesar dari NIM", nimCari, "adalah: ", nilaiTerbesar)
	} else {
		fmt.Println("NIM", nimCari, "tidak ditemukan")
	}
}
