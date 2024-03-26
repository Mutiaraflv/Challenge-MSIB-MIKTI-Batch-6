package main

import "fmt"

// function untuk menyimpan data skor suatu tim kedalam array
func InputData(nama_tim *[3]int, skor1 int, skor2 int, skor3 int) {
	nama_tim[0] = skor1
	nama_tim[1] = skor2
	nama_tim[2] = skor3
	fmt.Println(&nama_tim[0])
}

type Data struct{
	name string
	berat, tinggi float64
}

func (data *Data) getBerat(){ 
	fmt.Println(data.berat)
}



func main() {
	// // deklarasi variabel untuk data masing-masing tim  
	// var Lumba_lumba [3]int
	// var Koala [3]int
	

	// InputData(&Koala, 96, 108, 89)
	// fmt.Println(&Koala[0])
	// InputData(&Koala, 88, 91, 110)
	// // PenentuanPemenang("Data 1", Lumba_lumba, Koala)
	// fmt.Println(&Koala[0])
	
	// Data uji 1
	John := Data{"John", 92, 1.95}
	var am *float64 = &John.berat
	fmt.Println(am)

	// Data uji 2
	John = Data{"John", 85, 1.76}
	var am2 *float64 = &John.berat
	fmt.Println(am2)

	John.getBerat()
	

}
