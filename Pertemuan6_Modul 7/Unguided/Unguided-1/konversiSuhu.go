package main

import "fmt"

type suhu float64

func CelciusToReamur(celcius suhu) suhu {
	return (4.0 / 5.0) * celcius
}

func CelciusToFahrenheit(celcius suhu) suhu {
	return (9.0/5.0)*celcius + 32
}

func CelciusToKelvin(celcius suhu) suhu {
	return celcius + 273.15
}

func main() {
	var celcius suhu

	fmt.Println("=== KONVERTER CELCIUS ===")
	fmt.Print("Masukan Suhu (celcius) : ")
	fmt.Scan(&celcius)

	reamur := CelciusToReamur(celcius)
	fahrenheit := CelciusToFahrenheit(celcius)
	kelvin := CelciusToKelvin(celcius)

	fmt.Println(celcius, "celcius = ", reamur, "reamur")
	fmt.Println(celcius, "celcius = ", fahrenheit, "fahrenheit")
	fmt.Println(celcius, "celcius = ", kelvin, "kelvin")
}
