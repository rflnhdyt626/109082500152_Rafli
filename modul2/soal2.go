package main

import "fmt"

func main() { 
	var (
		gelas1, gelas2, gelas3, gelas4 string
		percobaan                       int = 1
		berhasil                         bool = true
	)

	for percobaan < 5 {
		fmt.Print("Masukkan warna zat cair pada gelas 1 hingga 4 pada percobaan ke-", percobaan, ": ")
		fmt.Scan(&gelas1, &gelas2, &gelas3, &gelas4)

		if !(gelas1 == "merah" && gelas2 == "kuning" && gelas3 == "hijau" && gelas4 == "ungu") {
			berhasil = false
		}
		percobaan++
	}
	fmt.Println("Berhasil:", berhasil)
}
