package main

import "fmt"

func main() {
	var mahasiswa [3] string

	mahasiswa[0] = "Mia"
	mahasiswa[1] = "Nomi"

	var i int
	for i = 0; i < 3; i++ {
		fmt.Printf("Masukkan data Index ke-%d : ")
		fmt.Scan(&i)
	}

	fmt.Println("Index ke-0 : ", mahasiswa[0])
	fmt.Println("Index ke-1 : ", mahasiswa[1])

	var j int
	for j = 0; j < 3; j++ {
		fmt.Println("Data ke-", j, " : ", mahasiswa[i])
	}

	fmt.Println("Index [:3] : ", mahasiswa[:3])
	fmt.Println("Index [1:] : ", mahasiswa[1:])		


}