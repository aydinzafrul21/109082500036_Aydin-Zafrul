# <h1 align="center">Laporan Praktikum Modul 3 - Fungsi</h1>
<p align="center">[Aydin Zafrul] - [109082500036]</p>

## Unguided 

### 1. [Soal 1]
#### soal1.go

```go
package main

import "fmt"

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	hasil := 1
	for i := 2; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

func permutation(n, r int) int {
	return factorial(n) / factorial(n-r)
}

func combination(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {
	var a, b, c, d int
	var p1, c1, p2, c2 int

	fmt.Scan(&a, &b, &c, &d)

	p1 = permutation(a, c)
	c1 = combination(a, c)
	p2 = permutation(b, d)
	c2 = combination(b, d)

	fmt.Println(p1, c1)
	fmt.Println(p2, c2)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Soal 1](https://github.com/aydinzafrul21/109082500036_Aydin-Zafrul/blob/main/WEEK3/output/output-1.png)
[Kita perlu menginputkan bilangan a, b, c, dan d. Program akan menghitung nilai permutasi a dgn c dan dilanjut menghitung kombinasi a dgn c, begitu juga dengan bilangan b dgn d. Permutasi dihitung dgn menggunakan faktorial terlebih dahulu di dalam subprogram. Setelah itu, program akan menampilkan output hasil permutasi dan kombinasi.]

### 2. [Soal 2]
#### soal2.go

```go
package main

import "fmt"

func f(x int) int {
	return x * x
}

func g(x int) int {
	return x - 2
}

func h(x int) int {
	return x + 1
}

func main() {
	var a, b, c int

	fmt.Scan(&a, &b, &c)

	fmt.Println(f(g(h(a))))
	fmt.Println(g(h(f(b))))
	fmt.Println(h(f(g(c))))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Soal 2](https://github.com/aydinzafrul21/109082500036_Aydin-Zafrul/blob/main/WEEK3/output/output-2.png)
[Kita perlu menginputkan bilangan a, b, c. Program akan menghitung f(x), g(x), dan h(x). Lalu, program akan menghitung fungsi komposisi (fogoh)(a), (gohof)(b), dan (hofog)(c). Setelah itu, program akan menampilkan output hasilnyas.]

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
![Screenshot Output Soal 3](https://github.com/aydinzafrul21/109082500036_Aydin-Zafrul/blob/main/WEEK3/output/output-3.png)
[Kita perlu menginputkan bilangan pada baris pertama dan kedua terdiri dari titik koordinat pusat serta radius lingkaran (a, b, c, dan d) dan baris ketiga adalah titik koordinat yg akan diperiksa. Program akan menghitung jarak antara titik (a, b) dan (c, d), dan program akan menentukan posisi suatu titik tersebut berada di dalam atau di luar lingkaran. Setelah itu, program akan menampilkan ouput hasilnya.]