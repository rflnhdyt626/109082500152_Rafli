package main

import "fmt"

var totalBiaya, biayaKg, biayaGr int

func biayaKirim(parsel int) (int, int, int) {

	var biayaGr int
	
	hargaPerKg := 10000
	kilo := parsel / 1000
	sisa := parsel % 1000
	biayaKg := kilo * hargaPerKg
	if sisa > 0 && sisa < 500 {
		biayaGr = sisa * 15
	} else if sisa >= 500 {
		biayaGr = sisa * 5
	}

	if kilo > 10 && sisa > 0 {
		totalBiaya = biayaKg
	} else {
		totalBiaya = biayaKg + biayaGr
	}

	return biayaKg, biayaGr, totalBiaya
}

func main() {

	var (
		beratParsel int
	)

	fmt.Print("Masukkan berat parsel (gram): ")
	fmt.Scanln(&beratParsel)
	
	biayaKg, biayaGr, totalBiaya = biayaKirim(beratParsel)

	fmt.Printf("Berat parsel (gram): %d\n", beratParsel)
	if beratParsel%1000 == 0 {
		fmt.Printf("Detail berat parsel: %d kg\n", beratParsel/1000)
		fmt.Printf("Detail biaya: Rp %d\n", biayaKg)
	} else {
		fmt.Printf("Detail berat parsel: %d kg + %d gr\n", beratParsel/1000, beratParsel%1000)
		fmt.Printf("Detail biaya: Rp %d + Rp %d\n", biayaKg, biayaGr)
	}
	fmt.Printf("Total biaya: Rp %d\n", totalBiaya)
}
