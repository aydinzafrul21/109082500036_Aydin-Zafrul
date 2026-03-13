# <h1 align="center">Laporan Praktikum Modul 1 - Review Algoritma & Pemrograman 1</h1>
<p align="center">[Aydin Zafrul] - [109082500036]</p>

## Unguided 

### 1. [Soal 1]
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
![Screenshot Output Soal 1](https://github.com/aydinzafrul21/109082500036_Aydin-Zafrul/blob/main/WEEK2/output/output-soal1.png)
[Program di atas digunakan untuk menukar atau menggeser posisi tiga buah string yang dimasukkan oleh pengguna. Pertama, program meminta pengguna memasukkan tiga kata lalu menampilkannya sebagai output awal. Setelah itu, program memakai variabel sementara temp untuk membantu menukar posisi data, sehingga nilai satu pindah ke tiga, dua pindah ke satu, dan tiga pindah ke dua.]

### 2. [Soal 2]
#### soal2.go

```go
package main

import "fmt"

func main() {
	var i, j int
	var w1, w2, w3, w4 string
	var berhasil bool

	for i = 1; i <= 5; i++ {
		fmt.Print("Percobaan ", i, ": ")
		fmt.Scan(&w1, &w2, &w3, &w4)

		if w1 == "merah" && w2 == "kuning" && w3 == "hijau" && w4 == "ungu" {
			j++
		}

	}

	berhasil = j == 5
	fmt.Println("BERHASIL:", berhasil)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Soal 2](https://github.com/aydinzafrul21/109082500036_Aydin-Zafrul/blob/main/WEEK2/output/output-soal2.png)
[Program di atas digunakan untuk mengecek apakah urutan warna yg dimasukkan benar selama 5 percobaan. Program akan meminta pengguna memasukkan 4 kata (w1, w2, w3, w4) pada setiap percobaan. Jika urutannya tepat yaitu "merah kuning hijau ungu", maka variabel j akan bertambah sebagai tanda percobaan itu benar.]

### 3. [Soal 3]
#### soal3.go

```go
package main

import "fmt"

func main() {
	var bp, sb, d1, bk, bj, tb int

	fmt.Print("Berat Parsel (gram): ")
	fmt.Scan(&bp)

	d1 = bp / 1000
	sb = bp % 1000

	if sb >= 500 && bp <= 10000 {
		bk = sb * 5
	} else if sb < 500 && bp <= 10000 {
		bk = sb * 15
	} else {
		bk = sb * 0
	}

	bj = d1 * 10000
	tb = bj + bk

	fmt.Printf("Detail berat: %d kg + %d gr\n", d1, sb)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", bj, bk)
	fmt.Printf("Total biaya pengiriman: Rp. %d\n", tb)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Soal 3](https://github.com/aydinzafrul21/109082500036_Aydin-Zafrul/blob/main/WEEK2/output/output-soal3.png)
[Program di atas berfungsi untuk menghitung biaya pengiriman parsel berdasarkan beratnya. Pengguna memasukkan berat dalam gram, lalu program mengubahnya menjadi kilogram dan sisa gram. Biaya dihitung dengan tarif Rp10.000 per kilogram, sedangkan sisa gram dikenakan biaya tambahan tergantung jumlahnya (≥500 gram Rp5/gram, <500 gram Rp15/gram) selama berat tidak lebih dari 10 kg. Setelah itu, program menampilkan detail berat, rincian biaya, dan total biaya pengiriman.]