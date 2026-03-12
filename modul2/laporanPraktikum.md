# <h1 align="center">Laporan Praktikum Modul 2 - REVIEW ALGORITMA DAN PEMROGRAMAN 1 </h1>
<p align="center">Rafli Nurhidayat - 109082500152</p>

## Unguided 

### 1. Soal Latihan Modul 2A
#### soal1.go

```go
package main

import "fmt"

func main() {
	var (
		satu, dua, tiga string
		temp            string
	)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rflnhdyt626/109082500152_Rafli/blob/main/modul2/output/output-soal1.png)
[Program tersebut meminta user untuk memasukan string tiga kali, lalu program tersebut mencetak/print output awalnya yaitu tiga string yang diinput user sesuai urutan input. Setelah itu, program menukar posisi ketiga string tersebut (string pertama menjadi kedua, kedua menjadi ketiga, dan ketiga menjadi pertama), lalu mencetak/print output akhir dengan urutan yang sudah berubah.]

### 2. Soal Latihan Modul 2B
#### soal2.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rflnhdyt626/109082500152_Rafli/blob/main/modul2/output/output-soal2.png)
[Program meminta user untuk memasukkan susunan warna zat cair pada gelas 1 hingga 4 secara berurutan sebanyak 5 kali percobaan menggunakan perulangan. Pada setiap percobaan, program memeriksa apakah keempat warna yang dimasukkan sesuai dengan urutan yang ditentukan, yaitu 'merah', 'kuning', 'hijau', dan 'ungu'. Jika pada salah satu percobaan terdapat ketidaksesuaian warna, program akan mengubah nilai variabel berhasil menjadi false dan nilai tersebut tidak akan berubah kembali menjadi true meskipun percobaan berikutnya benar, karena nilai false bersifat permanen hingga program selesai. Setelah semua 5 percobaan selesai, program mencetak hasil akhir berupa true jika seluruh percobaan berhasil, atau false jika minimal satu percobaan gagal.]

### 3. Soal Latihan Modul 2C
#### soal3.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rflnhdyt626/109082500152_Rafli/blob/main/modul2/output/output-soal3.png)
[Program ini terdiri dari dua fungsi, yaitu 'main sebagai fungsi utama dan 'biayaKirim' sebagai fungsi untuk menghitung ongkos kirim. Pada fungsi 'main', user diminta memasukkan berat parsel dalam satuan gram. Nilai berat tersebut dikirim ke fungsi 'biayaKirim' untuk diproses menjadi tiga komponen/variabel biaya yaitu, biaya per kilogram ('biayaKg'), biaya sisa gram ('biayaGr'), dan total biaya ('totalBiaya').

Di dalam fungsi 'biayaKirim', berat parsel dibagi menjadi bagian kilogram ('kilo') dan sisa gram ('sisa'). Biaya kilogram dihitung dengan harga Rp10.000 per kg. Untuk sisa gram, tarifnya bertingkat: jika sisa 1–499 gram maka dikenakan Rp15 per gram, sedangkan jika sisa 500 gram atau lebih dikenakan Rp5 per gram. Namun, ada aturan khusus: jika berat sudah lebih dari 10 kg dan masih ada sisa gram, maka sisa gram tidak dihitung, sehingga total biaya hanya berdasarkan biaya kilogram.

Setelah fungsi 'biayaKirim' mengembalikan ketiga nilai tersebut, fungsi 'main` menampilkan detail berat parsel, rincian biaya, dan total biaya akhir. Jika berat tepat kelipatan 1000 gram, output hanya menampilkan kilogram; jika tidak, output menampilkan format kilogram + gram beserta rincian biaya masing-masing.]
```

