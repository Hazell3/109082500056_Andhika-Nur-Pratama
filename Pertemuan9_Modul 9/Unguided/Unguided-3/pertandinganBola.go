package main

import "fmt"

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var hasil [100]string
	var n int = 0

	fmt.Print("Klub A : ")
	fmt.Scan(&klubA)

	fmt.Print("Klub B : ")
	fmt.Scan(&klubB)

	i := 1

	for {
		fmt.Printf("Pertandingan %d : ", i)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			hasil[n] = klubA
		} else if skorB > skorA {
			hasil[n] = klubB
		} else {
			hasil[n] = "Draw"
		}

		n++
		i++
	}

	for i := 0; i < n; i++ {
		fmt.Printf("Hasil %d : %s\n", i+1, hasil[i])
	}

	fmt.Println("Pertandingan selesai")
}
