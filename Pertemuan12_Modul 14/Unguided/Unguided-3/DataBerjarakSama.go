package main

import "fmt"

const NMAX = 1000

type arrInt [NMAX]int

func insertionSort(A *arrInt, n int) {
	var i, j, temp int

	for i = 1; i < n; i++ {
		temp = A[i]
		j = i - 1

		for j >= 0 && A[j] > temp {
			A[j+1] = A[j]
			j--
		}

		A[j+1] = temp
	}
}

func main() {
	var A arrInt
	var x, n int

	n = 0

	for {
		fmt.Scan(&x)

		if x < 0 {
			break
		}

		A[n] = x
		n++
	}

	insertionSort(&A, n)

	for i := 0; i < n; i++ {
		fmt.Print(A[i], " ")
	}
	fmt.Println()

	jarakTetap := true
	selisih := A[1] - A[0]

	for i := 2; i < n; i++ {
		if A[i]-A[i-1] != selisih {
			jarakTetap = false
		}
	}

	if jarakTetap {
		fmt.Println("Data berjarak", selisih)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}
