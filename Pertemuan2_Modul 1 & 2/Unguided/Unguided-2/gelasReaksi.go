package main
import "fmt"

func main() {
	var w1, w2, w3, w4 string
	berhasil := true

	for i := 1; i <= 5; i++ {
		fmt.Printf("Percobaan %d: ", i)
		fmt.Scan(&w1, &w2, &w3, &w4)

		if !(w1 == "Merah" && w2 == "Kuning" && w3 == "Hijau" && w4 == "Ungu") {
			berhasil = false
		}
	}
	fmt.Println("Berhasil", berhasil)
}