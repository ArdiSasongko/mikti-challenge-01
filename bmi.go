package main

import "fmt"

//deklarasi interface
type actPerson interface {
	bmi() float64
	getInfo()
}

//deklarasi struct
type person struct {
	name          string
	berat, tinggi float64
}

//membuat func menghitung bmi
func (p person) bmi() float64 {
	return p.berat / p.tinggi * p.tinggi
}

func (p person) getInfo() {
	fmt.Printf("%s memiliki berat %.2f Kg dan tinggi %.2f M dengan bmi %.2f\n", p.name, p.berat, p.tinggi, p.bmi())
}

func MarkHigherBMI(p, q person) {
	//deklarasi markHigherBMI
	markHigherBMI := false

	p.getInfo()
	q.getInfo()
	if p.bmi() > q.bmi() {
		markHigherBMI = true
		fmt.Println("Mark BMI Higher than Jhon", markHigherBMI)
	} else {
		fmt.Println("Mark BMI Higher than Jhon", markHigherBMI)
	}
}

func main() {
	//Data 1: Berat Mark 78 kg dan tinggi 1.69 m. Berat John 92 kg dan tinggi 1.95 m.
	//Data 2: Berat Mark 95 kg dan tinggi 1.88 m. Berat John 85 kg dan tinggi 1.76 m.

	//inisiasi data pertama mark dan john
	mark1 := person{"Mark", 78.0, 1.69}
	john1 := person{"John", 92.0, 1.95}

	//inisiasi data kedua mark dan john
	mark2 := person{"Mark", 95.0, 1.88}
	john2 := person{"John", 85.0, 1.76}

	fmt.Println("Data Pertama : ")
	MarkHigherBMI(mark1, john1)

	fmt.Println("Data Kedua : ")
	MarkHigherBMI(mark2, john2)
}
