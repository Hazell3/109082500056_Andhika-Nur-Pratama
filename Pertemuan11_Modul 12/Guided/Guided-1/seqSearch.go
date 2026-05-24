package main

import "fmt"

func sequentialSearch(arrBuah [5]string, dataDicari string) int {
	var idx_found = -1
	for i := 0; i < len(arrBuah); i++ {
		if arrBuah[i] == dataDicari {
			idx_found = i
			break
		}
	}
	return idx_found
}

func main() {
	var arrBuah [5]string

	for i := 0; i < len(arrBuah); i++ {
		fmt.Printf("Masukkan data buah ke-%d : ", i)
		fmt.Scan(&arrBuah[i])
	}

	var dataDicari string
	fmt.Println("Masukkan data buah yang ingin dicari :  ")
	fmt.Scan(&dataDicari)

	var index_data = sequentialSearch(arrBuah, dataDicari)

	if index_data > -1 {
		fmt.Printf("Data %s ditemukan pada index ke-%d!", dataDicari, index_data)
	} else if index_data == -1 {
		fmt.Printf("data %s tidak ditemukan!", dataDicari)
	}
}
