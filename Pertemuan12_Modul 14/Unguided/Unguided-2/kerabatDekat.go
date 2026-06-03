package main

import "fmt"

const NMAX = 1000

type arrInt [NMAX]int

func selectionSort(A *arrInt, n int) {
	var i, j, idxMin, temp int

	for i = 0; i < n-1; i++ {
		idxMin = i

		for j = i + 1; j < n; j++ {
			if A[j] < A[idxMin] {
				idxMin = j
			}
		}

		temp = A[i]
		A[i] = A[idxMin]
		A[idxMin] = temp
	}
}

func main() {
	var daerah, m, x int
	var ganjil, genap arrInt
	var ng, ne int

	fmt.Scan(&daerah)

	for d := 0; d < daerah; d++ {

		ng = 0
		ne = 0

		fmt.Scan(&m)

		for i := 0; i < m; i++ {

			fmt.Scan(&x)

			if x%2 == 1 {
				ganjil[ng] = x
				ng++
			} else {
				genap[ne] = x
				ne++
			}
		}

		selectionSort(&ganjil, ng)
		selectionSort(&genap, ne)

		for i := 0; i < ng; i++ {
			fmt.Print(ganjil[i], " ")
		}

		for i := ne - 1; i >= 0; i-- {
			fmt.Print(genap[i], " ")
		}

		fmt.Println()
	}
}
