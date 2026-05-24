package main

import "fmt"

func main() {
	var x int
	var suara [21]int
	var total, sah int

	for {
		fmt.Scan(&x)

		total++

		if x == 0 {
			break
		}

		if x >= 1 && x <= 20 {
			suara[x]++
			sah++
		}
	}

	ketua := 1
	wakil := 1

	for i := 2; i <= 20; i++ {
		if suara[i] > suara[ketua] {
			ketua = i
		}
	}

	for i := 1; i <= 20; i++ {
		if i != ketua {
			wakil = i
			break
		}
	}

	for i := 1; i <= 20; i++ {
		if i != ketua {
			if suara[i] > suara[wakil] {
				wakil = i
			}
		}
	}

	fmt.Println("Suara masuk:", total)
	fmt.Println("Suara sah:", sah)
	fmt.Println("Ketua RT:", ketua)
	fmt.Println("Wakil ketua:", wakil)
}
