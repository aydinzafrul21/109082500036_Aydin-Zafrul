# <h1 align="center">Laporan Praktikum Modul 9 - ARRAY</h1>
<p align="center">[Aydin Zafrul] - [109082500036]</p>

## Unguided 

### 1. [Soal 1]
#### soal1.go

```go
package main

import (
	"fmt"
	"math"
)

type Titik struct {
	x, y float64
}

type Lingkaran struct {
	pusat  Titik
	radius float64
}

func jarak(p, q Titik) float64 {
	return math.Sqrt(math.Pow(p.x-q.x, 2) + math.Pow(p.y-q.y, 2))
}

func didalam(c Lingkaran, p Titik) bool {
	return jarak(c.pusat, p) < c.radius
}

func main() {
	var l1, l2 Lingkaran
	var p Titik

	fmt.Scan(&l1.pusat.x, &l1.pusat.y, &l1.radius)
	fmt.Scan(&l2.pusat.x, &l2.pusat.y, &l2.radius)
	fmt.Scan(&p.x, &p.y)

	inL1 := didalam(l1, p)
	inL2 := didalam(l2, p)

	if inL1 && inL2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if inL1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if inL2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Soal 1](https://github.com/aydinzafrul21/109082500036_Aydin-Zafrul/blob/main/WEEK9/output/1.png)
[Program pertama menggunakan tipe bentukan (struct) Titik dan Lingkaran untuk menyimpan koordinat pusat serta radius , lalu menghitung jarak antara titik sembarang dengan pusat lingkaran menggunakan rumus akar kuadrat selisih koordinat guna menentukan apakah titik tersebut berada di dalam atau di luar radius lingkaran yang diberikan.]

### 2. [Soal 2]
#### soal2.go

```go
package main

import (
	"fmt"
	"math"
)

const NMAX int = 2000

type IntArray [NMAX]int

func main() {
	var data IntArray
	var n, x, delIdx int

	fmt.Print("Masukkan jumlah elemen (N): ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}

	fmt.Print("Isi array: ")
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", data[i])
	}
	fmt.Println()

	fmt.Print("Indeks ganjil: ")
	for i := 1; i < n; i += 2 {
		fmt.Printf("%d ", data[i])
	}
	fmt.Print("\nIndeks genap: ")
	for i := 0; i < n; i += 2 {
		fmt.Printf("%d ", data[i])
	}

	fmt.Print("\nMasukkan nilai x untuk kelipatan indeks: ")
	fmt.Scan(&x)
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Printf("%d ", data[i])
		}
	}

	fmt.Print("\nHapus indeks ke: ")
	fmt.Scan(&delIdx)
	for i := delIdx; i < n-1; i++ {
		data[i] = data[i+1]
	}
	n--
	fmt.Print("Isi setelah dihapus: ")
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", data[i])
	}

	sum := 0.0
	for i := 0; i < n; i++ {
		sum += float64(data[i])
	}
	mean := sum / float64(n)

	var devSum float64
	for i := 0; i < n; i++ {
		devSum += math.Pow(float64(data[i])-mean, 2)
	}
	sd := math.Sqrt(devSum / float64(n))

	fmt.Printf("\nRata-rata: %.2f\nStandar Deviasi: %.2f\n", mean, sd)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Soal 2](https://github.com/aydinzafrul21/109082500036_Aydin-Zafrul/blob/main/WEEK9/output/2.png)
[Program ini mengelola sekumpulan data dalam array statis dengan kapasitas tetap, di mana pengguna dapat menampilkan elemen berdasarkan kriteria tertentu seperti indeks ganjil, genap, atau kelipatan $x$. Selain itu, program ini mampu melakukan penghapusan data dengan metode pergeseran elemen (shifting) untuk menjaga kontinuitas array, serta menghitung analisis statistik yang mencakup nilai rata-rata, standar deviasi, dan frekuensi kemunculan angka tertentu.]

### 3. [Soal 3]
#### soal3.go

```go
package main

import "fmt"

func main() {
	var klub1, klub2 string
	var pemenang []string
	var skor1, skor2 int
	pertandingan := 1

	fmt.Print("Klub A: ")
	fmt.Scan(&klub1)
	fmt.Print("Klub B: ")
	fmt.Scan(&klub2)

	for {
		fmt.Printf("Pertandingan %d: ", pertandingan)
		fmt.Scan(&skor1, &skor2)

		if skor1 < 0 || skor2 < 0 {
			break
		}

		if skor1 > skor2 {
			pemenang = append(pemenang, klub1)
		} else if skor2 > skor1 {
			pemenang = append(pemenang, klub2)
		} else {
			pemenang = append(pemenang, "Draw")
		}
		pertandingan++
	}

	for i, hasil := range pemenang {
		fmt.Printf("Hasil %d: %s\n", i+1, hasil)
	}
	fmt.Println("Pertandingan selesai")
}
```
### Output Unguided :

##### Output 
![Screenshot Output Soal 3](https://github.com/aydinzafrul21/109082500036_Aydin-Zafrul/blob/main/WEEK9/output/3.png)
[Program ini berfungsi untuk merekap hasil pertandingan sepak bola dengan menyimpan nama klub pemenang ke dalam sebuah array. Proses penginputan skor dilakukan secara berulang dan hanya akan berhenti apabila pengguna memasukkan nilai negatif pada salah satu skor klub. Selama proses berlangsung, program membandingkan perolehan gol untuk menentukan pemenang atau hasil seri, kemudian menyimpan serta menampilkan daftar hasil seluruh pertandingan secara berurutan setelah sesi input berakhir.]

### 4. [Soal 4]
#### soal4.go

```go
package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var char rune
	*n = 0
	for *n < NMAX {
		fmt.Scanf("%c", &char)
		if char == '.' {
			break
		}
		if char != '\n' && char != ' ' {
			t[*n] = char
			(*n)++
		}
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

func isPalindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	isiArray(&tab, &m)

	pal := isPalindrom(tab, m)

	fmt.Print("Reverse teks: ")
	balikanArray(&tab, m)
	cetakArray(tab, m)

	fmt.Printf("Palindrom? %t\n", pal)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Soal 4](https://github.com/aydinzafrul21/109082500036_Aydin-Zafrul/blob/main/WEEK9/output/4.png)
[Program ini dirancang untuk memproses sekumpulan karakter dalam array dengan batas maksimal NMAX 127. Subprogram isiArray mengisi elemen hingga ditemukan karakter titik, sementara balikanArray mengubah urutan isi array menggunakan teknik pertukaran elemen (swapping) dari kedua ujung ke tengah. Selain membalikkan teks, program ini juga menyertakan fungsi palindrom untuk memvalidasi apakah susunan karakter terbaca sama dari depan maupun belakang, seperti pada kata "KATAK".]