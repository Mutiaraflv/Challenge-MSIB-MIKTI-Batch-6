package main

import "fmt"

// struct untuk menyimpan data nama, berat badan, dan tinggi badan seseorang
type Data struct{
	name string
	berat, tinggi float64
}

// method untuk menghitung BMI berdasarkan data struct
func (data Data) PerhitunganBMI() float64{ 
	// rumus BMI = massa / tinggi ** 2 = massa / (tinggi * tinggi)
	bmi := data.berat/(data.tinggi*data.tinggi)
	return bmi
}

// function untuk membandingkan BMI dari dua data struct
func PerbandinganDataBMI(data_1 Data, data_2 Data) bool{
	dataHigherBMI := data_1.PerhitunganBMI() > data_2.PerhitunganBMI()
	return dataHigherBMI
}

func main(){
	// membuat var markHigherBMI untuk membandingkan data BMI Mark dan John
	markHigherBMI := PerbandinganDataBMI

	// Data uji 1
	Mark := Data{"Mark", 78, 1.69}
	John := Data{"John", 92, 1.95}

	fmt.Println("Nilai variabel markHigherBMI data uji ke-1 :", markHigherBMI(Mark, John))


	// Data uji 2
	Mark = Data{"Mark", 95, 1.88}
	John = Data{"John", 85, 1.76}

	fmt.Println("Nilai variabel markHigherBMI data uji ke-2 :", markHigherBMI(Mark, John))
}