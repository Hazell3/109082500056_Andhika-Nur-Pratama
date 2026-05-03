package main

import "fmt"

const NMAX = 1000

type arrKelinci [NMAX]float64

func MinMax(data arrKelinci, n int, max, min *float64) {
	*min = data[0]
	*max = data[0]

	for i := 1; i < n; i++ {
		if data[i] < *min {
			*min = data[i]
		}
		if data[i] > *max {
			*max = data[i]
		}
	}
}

func main() {
	var n int
	var data arrKelinci
	var max, min float64

	fmt.Print("Banyak anak kelinci: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Berat anak kelinci ke-%d: ", i+1)
		fmt.Scan(&data[i])
	}

	MinMax(data, n, &max, &min)
	fmt.Printf("Berat terkecil: %.2f\n", min)
	fmt.Printf("Berat terbesar: %.2f\n", max)
}
