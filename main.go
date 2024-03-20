package main

import (
	"fmt"
	"strings"
)

// membuat fungsi untuk menghitung rata-rata
func average(data []int) float64 {
	result := 0.0
	for _, n := range data {
		//n diubah ke float
		result = result + float64(n)
	}
	//len(data) diubah ke float
	return result / float64(len(data))
}

// membuat fungsi compare
func compare(a, b float64) {
	if a >= 100 && a > b {
		fmt.Printf("Lumba-Lumba Menang, rata-rata : %.2f  \n", a)
	} else if b >= 100 && b > a {
		fmt.Printf("Koala Menang, rata-rata : %.2f  \n", b)
	} else {
		if a > b {
			fmt.Printf("Lumba-Lumba Menang, rata-rata : %.2f  \n", a)
		} else if b > a {
			fmt.Printf("Koala Menang, rata-rata : %.2f  \n", b)
		} else {
			fmt.Printf("Imbang, rata-rata dari keduanya %.2f \n", a)
		}
	}
}

func main() {
	for {
		//Data Pertama
		dataLumba_1 := [3]int{96, 108, 89}
		dataKoala_1 := [3]int{88, 91, 110}

		//Data Kedua
		dataLumba_2 := [3]int{97, 112, 101}
		dataKoala_2 := [3]int{109, 95, 123}

		//Data Ketiga
		dataLumba_3 := [3]int{97, 112, 101}
		dataKoala_3 := [3]int{109, 95, 106}

		//deklarasi variabel input
		var input int

		fmt.Println("Static")
		fmt.Println("Pilih data yang akan digunakan")
		fmt.Printf("Data 1 \n Lumba-Lumba : %v \n Koala : %v \n", dataLumba_1, dataKoala_1)
		fmt.Printf("Data 2 \n Lumba-Lumba : %v \n Koala : %v \n", dataLumba_2, dataKoala_2)
		fmt.Printf("Data 3 \n Lumba-Lumba : %v \n Koala : %v \n", dataLumba_3, dataKoala_3)
		fmt.Printf("Pilih Data : ")
		//memasukan pilihan
		fmt.Scanln(&input)

		switch input {
		case 1:
			a := average(dataLumba_1[:])
			b := average(dataKoala_1[:])
			compare(a, b)
		case 2:
			a := average(dataLumba_2[:])
			b := average(dataKoala_2[:])
			compare(a, b)
		case 3:
			a := average(dataLumba_3[:])
			b := average(dataKoala_3[:])
			compare(a, b)
		default:
			fmt.Println("Pilihan tidak ada")
		}

		var y string
		fmt.Print("Ulang? (y/n) : ")
		fmt.Scanln(&y)

		//melakukan perulangan sampai tidak sama y
		if strings.ToLower(y) != "y" {
			break
		}
	}

}
