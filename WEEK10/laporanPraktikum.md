# <h1 align="center">Laporan Praktikum Modul 9 - ARRAY</h1>
<p align="center">[Aydin Zafrul] - [109082500036]</p>

## Unguided 

### 1. [Soal 1]
#### soal1.go

```go
package main

import "fmt"

func main() {
	var kelinci [1000]float64
	var n int

	fmt.Print("Masukkan jumlah kelinci: ")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&kelinci[i])
	}

	max := kelinci[0]
	min := kelinci[0]

	for i := 1; i < n; i++ {
		if kelinci[i] > max {
			max = kelinci[i]
		}

		if kelinci[i] < min {
			min = kelinci[i]
		}
	}

	fmt.Print("Min ", min)
	fmt.Print()
	fmt.Print("Max ", max)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Soal 1](https://github.com/aydinzafrul21/109082500036_Aydin-Zafrul/blob/main/WEEK10/output/1.png)
[Program ini menggunakan array dengan kapasitas 1000 untuk menyimpan data berat anak kelinci, kemudian melakukan iterasi pada array tersebut untuk membandingkan setiap elemen guna menentukan nilai minimum dan maksimum dari seluruh berat yang telah dimasukkan.]

### 2. [Soal 2]
#### soal2.go

```go
package main

import "fmt"

func main() {
	var beratIkan [1000]float64
	var x, y int
	fmt.Print("Masukkan jumlah ikan (x) dan kapasitas wadah (y): ")
	fmt.Scan(&x, &y)

	for i := 0; i < x; i++ {
		fmt.Scan(&beratIkan[i])
	}

	var totalPerWadah float64
	var jumlahWadah int
	var totalSeluruhWadah float64

	for i := 0; i < x; i++ {
		totalPerWadah += beratIkan[i]
		if (i+1)%y == 0 || i == x-1 {
			fmt.Print(totalPerWadah, " ")

			totalSeluruhWadah += totalPerWadah
			jumlahWadah++
			totalPerWadah = 0
		}
	}
	fmt.Println()

	if jumlahWadah > 0 {
		rataRata := totalSeluruhWadah / float64(jumlahWadah)
		fmt.Println(rataRata)
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Soal 2](https://github.com/aydinzafrul21/109082500036_Aydin-Zafrul/blob/main/WEEK10/output/2.png)
[Program ini menyimpan data berat ikan ke dalam array berkapasitas 1000, lalu mengelompokkan berat tersebut ke dalam beberapa wadah berdasarkan kapasitas yang ditentukan menggunakan operasi modulo dan akumulasi nilai, serta menghitung rata-rata berat per wadah dengan membagi total seluruh berat dengan jumlah wadah yang terbentuk.]

### 3. [Soal 3]
#### soal3.go

```go
package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]
	for i := 1; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

func rerata(arrBerat arrBalita, n int) float64 {
	var total float64
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}
	return total / float64(n)
}

func main() {
	var data arrBalita
	var n int
	var min, max float64

	fmt.Print("Masukan banyak data berat balita: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&data[i])
	}
	hitungMinMax(data, n, &min, &max)

	fmt.Printf("Berat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rerata(data, n))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Soal 3](https://github.com/aydinzafrul21/109082500036_Aydin-Zafrul/blob/main/WEEK10/output/3.png)
[Program ini menggunakan tipe bentukan array khusus arrBalita untuk menampung data berat badan, lalu memanfaatkan subprogram berupa prosedur dengan parameter pointer untuk mencari nilai ekstrem (minimum dan maksimum) serta sebuah fungsi untuk mengalkulasi nilai rerata secara modular berdasarkan jumlah data yang diinputkan.]