package main

import (
	"fmt"
)

func hitungSkor(soal *int, skor *int) {
	*soal = 0
	*skor = 0

	for i := 0; i < 8; i++ {
		var waktu int
		fmt.Scan(&waktu)

		if waktu <= 300 {
			*soal++
			*skor += waktu
		}
	}
}

func main() {
	var (
		nama, pemenang     string
		maxSoal, minSkor, totalSoal, totalSkor  int
		pertamaInput = true
	)

	for {
		fmt.Scan(&nama)
		if nama == "Selesai" {
			break
		}

		hitungSkor(&totalSoal, &totalSkor)

		if pertamaInput || totalSoal > maxSoal || (totalSoal == maxSoal && totalSkor < minSkor) {
			pemenang = nama
			maxSoal = totalSoal
			minSkor = totalSkor
			pertamaInput = false
		}
	}

	fmt.Println(pemenang, maxSoal, minSkor)
}