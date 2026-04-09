# <h1 align="center">Laporan Praktikum Modul 4 - Prosedur</h1>
<p align="center">[Aydin Zafrul] - [109082500036]</p>

## Unguided 

### 1. [Soal 1]
#### soal1.go

```go
package main

import "fmt"

func factorial(n int, hasil *int) {
	var i int

	*hasil = 1

	for i = 2; i <= n; i++ {
		*hasil *= i
	}
}

func permutation(n, r int, hasil *int) {
	var fn, fnr int

	factorial(n, &fn)
	factorial(n-r, &fnr)

	*hasil = fn / fnr
}

func combination(n, r int, hasil *int) {
	var fn, fr, fnr int

	factorial(n, &fn)
	factorial(r, &fr)
	factorial(n-r, &fnr)

	*hasil = fn / (fr * fnr)
}

func main() {
	var (
		a, b, c, d, p1, p2, c1, c2 int
	)

	fmt.Print("Input a, b, c, & d: ")
	fmt.Scan(&a, &b, &c, &d)

	permutation(a, c, &p1)
	combination(a, c, &c1)
	permutation(b, d, &p2)
	combination(b, d, &c2)

	fmt.Println(p1, c1)
	fmt.Println(p2, c2)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Soal 1](https://github.com/aydinzafrul21/109082500036_Aydin-Zafrul/blob/main/WEEK4/output/output-1.png)
[Kita perlu menginputkan bilangan a, b, c, dan d. Program akan menghitung nilai permutasi a dgn c dan dilanjut menghitung kombinasi a dgn c, begitu juga dengan bilangan b dgn d. Permutasi dihitung dgn menggunakan faktorial terlebih dahulu di dalam subprogram. Setelah itu, program akan menampilkan output hasil permutasi dan kombinasi.]

### 3. [Soal 3]
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

func didalam(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) < r
}

func main() {
	var cx1, cy1, r1 float64
	var cx2, cy2, r2 float64
	var x, y float64
	var d1, d2 bool

	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)

	d1 = didalam(cx1, cy1, r1, x, y)
	d2 = didalam(cx2, cy2, r2, x, y)

	if d1 && d2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if d1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if d2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Soal 3](https://github.com/aydinzafrul21/109082500036_Aydin-Zafrul/blob/main/WEEK4/output/output-2.png)
[Kita perlu menginputkan nama peserta yg kemudian diikuti dgn durasi peserta menyelesaikan soal. Program akan menjumlahkan setiap waktu yg diinputkan sampai 8 soal. Program akan terus berjalan hingga kita menginputkan nama peserta dengan kata "Selesai". Setelah itu, program akan menampilkan hasil output pemenangnya.]