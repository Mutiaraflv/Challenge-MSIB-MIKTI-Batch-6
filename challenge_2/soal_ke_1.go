package main

import "fmt"

// function untuk menyimpan data skor suatu tim kedalam array
func InputData(nama_tim *[3]int, skor1 int, skor2 int, skor3 int) {
	nama_tim[0] = skor1
	nama_tim[1] = skor2
	nama_tim[2] = skor3
}

// function untuk menhitung rata-rata skor suatu tim
func RataRataSkor(nama_tim [3]int) float64 {
	total := 0
	for _, skor := range nama_tim {
		total += skor
	}
	rata_rata := float64(total) / float64(len(nama_tim))

	return rata_rata
}

// function untuk menentukan tim yang menang
func PenentuanPemenang(nama_data string, tim1 [3]int, tim2 [3]int){
	SkorTimLumbaLumba := RataRataSkor(tim1)
	SkorTimKoala := RataRataSkor(tim2)
	
	var Pemenang string

	if SkorTimLumbaLumba >= 100 || SkorTimKoala >= 100 {
		if SkorTimLumbaLumba < SkorTimKoala {
			Pemenang = "Tim Koala menang"
		} else if SkorTimLumbaLumba > SkorTimKoala {
			Pemenang = "Tim Lumba-Lumba menang"
		} else {
			Pemenang = "Skor kedua tim seri"
		}
	} else {
		Pemenang = "Tidak ada tim yang menang"
	}

	fmt.Println("Hasil", nama_data, ":", Pemenang)
}

func main() {
	// deklarasi variabel untuk data masing-masing tim  
	var Lumba_lumba [3]int
	var Koala [3]int
	

	InputData(&Lumba_lumba, 96, 108, 89)
	InputData(&Koala, 88, 91, 110)
	PenentuanPemenang("Data 1", Lumba_lumba, Koala)


	InputData(&Lumba_lumba, 97, 112, 101)
	InputData(&Koala, 109, 95, 123)
	PenentuanPemenang("Data Bonus 1", Lumba_lumba, Koala)


	InputData(&Lumba_lumba, 97, 112, 101)
	InputData(&Koala, 109, 95, 106)
	PenentuanPemenang("Data Bonus 2", Lumba_lumba, Koala)

}
