# <h1 align="center">Laporan Praktikum Modul 3 - Fungsi </h1>
<p align="center">Rafli Nurhidayat - 109082500152</p>


## Unguided 

### 1. Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian membantu Jonas? (tidak tentunya ya :p). Masukan terdiri dari empat buah bilangan asli a, b, c, dan d yang dipisahkan oleh spasi, dengan syarat a ≥ c dan b ≥ d. Keluaran terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi a terhadap c, sedangkan baris kedua adalah hasil permutasi dan kombinasi b terhadap d.
#### soal1.go

```go
package main

import ("fmt"
)

func faktorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * faktorial(n-1)
}

func permutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}

func kombinasi(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}


func main() {
	var (
		a, b, c, d int
	)
	fmt.Scanln(&a, &b, &c, &d)
	fmt.Print(permutasi(a, c), " ")
	fmt.Println(kombinasi(a, c))

	fmt.Print(permutasi(b, d), " ")
	fmt.Println(kombinasi(b, d))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rflnhdyt626/109082500152_Rafli/blob/main/modul3/output/output-soal1.png)
Program ini berfungsi untuk menghitung permutasi dan kombinasi dari dua pasang bilangan. Program menerima empat bilangan bulat sebagai input, yaitu a, b, c, dan d, kemudian menghitung permutasi dan kombinasi dari a terhadap c, serta b terhadap d. Terdapat tiga fungsi utama dalam program ini. Pertama, fungsi faktorial(n int) yang bekerja secara rekursif/memanggil dirinya sendiri untuk menghitung nilai faktorial dari bilangan n. Fungsi ini akan terus memanggil dirinya sendiri dengan nilai n-1 hingga n == 0, yang mengembalikan nilai 1. Dengan cara ini, hasil akhirnya adalah perkalian beruntun dari n hingga 1. Kedua, fungsi permutasi(n, r int) yang menghitung banyaknya susunan r elemen dari n elemen yang tersedia menggunakan rumus P(n,r) = n! / (n-r)!. Ketiga, fungsi kombinasi(n, r int) yang menghitung banyaknya cara memilih r elemen dari n elemen tanpa memperhatikan urutan, menggunakan rumus C(n,r) = n! / (r! * (n-r)!). Kedua fungsi ini memanfaatkan fungsi faktorial untuk melakukan perhitungannya. Pada fungsi main, program membaca empat nilai input sekaligus, lalu mencetak hasil permutasi dan kombinasi untuk pasangan pertama (a, c) pada baris pertama, dan pasangan kedua (b, d) pada baris kedua.

### 2. Diberikan tiga buah fungsi matematika yaitu f(x) = x^2, g(x) = x − 2 dan h(x) = x + 1. Fungsi komposisi (fogoh)(x) artinya adalah f(g(h(x))). Tuliskan f(x), g(x) dan h(x) dalam bentuk function. Masukan terdiri dari sebuah bilangan bulat a, b dan c yang dipisahkan oleh spasi. Keluaran terdiri dari tiga baris. Baris pertama adalah (fogoh)(a), baris kedua (gohof)(b), dan baris ketiga adalah (hofog)(c)!
#### soal2.go

```go
package main

import ("fmt"
		"math")

func fx(n int) int {
	return int(math.Pow(float64(n), 2))
}

func gx(n int) int {
	return n - 2
}

func hx(n int) int {
	return n + 1
}

func main() {
	var (
		a, b, c int
	)
	fmt.Scanln(&a, &b, &c)
	fmt.Println(fx(gx(hx(a))))
	fmt.Println(gx(hx(fx(b))))
	fmt.Println(hx(fx(gx(c))))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rflnhdyt626/109082500152_Rafli/blob/main/modul3/output/output-soal2.png)
Program ini berfungsi untuk menghitung komposisi fungsi matematika dari tiga buah fungsi yaitu f(x), g(x), dan h(x). Program menerima tiga bilangan bulat sebagai input, yaitu a, b, dan c, kemudian menghitung hasil komposisi fungsi untuk masing-masing nilai tersebut. Terdapat tiga fungsi utama dalam program ini. Pertama, fungsi fx(n int) yang menghitung nilai f(x) = x² menggunakan math.Pow untuk memangkatkan nilai n dengan eksponen 2. Kedua, fungsi gx(n int)yang menghitung nilai g(x) = x - 2 dengan mengurangi nilai n sebesar 2. Ketiga, fungsi hx(n int) yang menghitung nilai h(x) = x + 1 dengan menambahkan nilai n sebesar 1. Pada fungsi main, program membaca tiga nilai input sekaligus, kemudian mencetak tiga baris output. Baris pertama adalah hasil komposisi (fogoh)(a) yaitu f(g(h(a))), baris kedua adalah hasil komposisi (gohof)(b) yaitu g(h(f(b))), dan baris ketiga adalah hasil komposisi (hofog)(c) yaitu h(f(g(c)))`. Setiap komposisi dihitung dengan cara meneruskan hasil satu fungsi sebagai input ke fungsi berikutnya secara bersarang.

### 3. [Lingkaran] Suatu lingkaran didefinisikan dengan koordinat titik pusat (cx, cy) dengan radius r. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (x, y) berdasarkan dua lingkaran tersebut. Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan bilangan bulat. Keluaran berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titik di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2".
#### soal3.go

```go
package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt((a-c)*(a-c) + (b-d)*(b-d))
}

func dalam(cx, cy, r, x, y float64) bool {
	return jarak(x, y, cx, cy) <= r
}

func main() {
	var (
		cx1, cy1, r1 float64
		cx2, cy2, r2 float64
		x, y         float64
	)

	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)

	in1 := dalam(cx1, cy1, r1, x, y)
	in2 := dalam(cx2, cy2, r2, x, y)

	if in1 && in2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if in1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if in2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rflnhdyt626/109082500152_Rafli/blob/main/modul3/output/output-soal3.png)
Program ini berfungsi untuk menentukan posisi sebuah titik sembarang (x, y) terhadap dua buah lingkaran. Program menerima tiga baris input, yaitu koordinat titik pusat dan radius lingkaran pertama (cx1, cy1, r1), koordinat titik pusat dan radius lingkaran kedua (cx2, cy2, r2), serta koordinat titik sembarang  (x, y). Terdapat dua fungsi utama dalam program ini. Pertama, fungsi jarak(a, b, c, d float64) yang menghitung jarak antara dua titik menggunakan rumus √((a-c)² + (b-d)²) dengan menggunakan math.Sqrt. Kedua, fungsi dalam(cx, cy, r, x, y float64) yang mengembalikan nilai true apabila jarak antara titik (x, y) dengan titik pusat lingkaran (cx, cy) kurang dari atau sama dengan radius r, yang berarti titik tersebut berada di dalam lingkaran. Pada fungsi main, program membaca koordinat kedua lingkaran dan titik sembarang, lalu menyimpan hasil pengecekan ke dalam variabel in1 dan in2. Kemudian program mencetak output berupa string sesuai kondisi: "Titik di dalam lingkaran 1 dan 2" jika titik berada di dalam kedua lingkaran, "Titik di dalam lingkaran 1" jika hanya di dalam lingkaran pertama, "Titik di dalam lingkaran 2" jika hanya di dalam lingkaran kedua, atau "Titik di luar lingkaran 1 dan 2" jika titik berada di luar kedua lingkaran.

