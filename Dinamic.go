package main

import (
	"fmt"
	"strings"
)

func inputData(n int) []int {
	// inisiasil data sebagai slice 0 dengan cap n
	data := make([]int, 0, n)
	for i := 0; i < n; i++ {
		var input int
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		fmt.Scanln(&input)
		data = append(data, input)
	}
	return data
}

// membuat fungsi untuk menghitung rata-rata
func averageData(data []int) float64 {
	sum := 0
	for _, n := range data {
		sum += n
	}
	return float64(sum) / float64(len(data))
}

// membuat fungsi compare
func compareData(a, b float64) {
	if a > b {
		fmt.Printf("Lumba-Lumba Menang, rata-rata: %.2f\n", a)
	} else if b > a {
		fmt.Printf("Koala Menang, rata-rata: %.2f\n", b)
	} else {
		fmt.Printf("Imbang, rata-rata dari keduanya: %.2f\n", a)
	}
}

func main() {
	for {
		var input int

		fmt.Println("Dinamik")
		fmt.Print("Banyak data yang ingin diinputkan: ")
		fmt.Scanln(&input)

		fmt.Println("Data Lumba-Lumba")
		dataLumba := inputData(input)

		fmt.Println("Data Koala")
		dataKoala := inputData(input)

		fmt.Printf("Data Lumba-Lumba: %v, average : %.2f\nData Koala: %v, average : %.2f\n", dataLumba, averageData(dataLumba), dataKoala, averageData(dataKoala))

		averageLumba := averageData(dataLumba)
		averageKoala := averageData(dataKoala)

		compareData(averageLumba, averageKoala)
		var y string
		fmt.Print("Ulang? (y/n) : ")
		fmt.Scanln(&y)

		//melakukan perulangan sampai tidak sama y
		if strings.ToLower(y) != "y" {
			break
		}
	}

}
