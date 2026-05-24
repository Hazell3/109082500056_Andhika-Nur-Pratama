package main

import "fmt"

func main() {
	var nilai map[string] int = make(map[string]int)

	nilai["Miku"] = 90
	
	fmt.Println(nilai["Miku"])

	nilai["Miku"] = 50
	fmt.Println(["Miku"])

	delete()
}