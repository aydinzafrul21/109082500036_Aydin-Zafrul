# <h1 align="center">Laporan Praktikum Modul 5 - Rekursif</h1>
<p align="center">[Aydin Zafrul] - [109082500036]</p>

## Unguided 

### 1. [Soal 1]
#### soal1.go

```go
package main

import "fmt"

func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah suku: ")
	fmt.Scan(&n)

	fmt.Println("Deret Fibonacci:")
	for i := 0; i <= n; i++ {
		fmt.Printf("%d ", fibonacci(i))
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Soal 1](https://github.com/aydinzafrul21/109082500036_Aydin-Zafrul/blob/main/WEEK5/output/output-1.png)
[Program meminta input jumlah suku Fibonacci, lalu menghitung setiap suku menggunakan fungsi rekursif. Fungsi akan memanggil dirinya sendiri dengan rumus Sn = Sn-1 + Sn-2 hingga mencapai kondisi dasar (0 dan 1). Setelah itu, program mencetak deret dari suku ke-0 sampai ke-n.]

### 2. [Soal 2]
#### soal2.go

```go
package main

import "fmt"

func polaBintang(n int, i int) {
	if i > n {
		return
	}

	for j := 0; j < i; j++ {
		fmt.Print("*")
	}
	fmt.Println()

	polaBintang(n, i+1)
}

func main() {
	var n int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&n)

	polaBintang(n, 1)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Soal 2](https://github.com/aydinzafrul21/109082500036_Aydin-Zafrul/blob/main/WEEK5/output/output-2.png)
[Program meminta input N lalu mencetak pola bintang dari 1 hingga N menggunakan fungsi rekursif. Setiap pemanggilan fungsi mencetak satu baris dengan jumlah bintang sesuai nilai i. Fungsi akan terus memanggil dirinya hingga i > N, lalu berhenti.]

### 3. [Soal 3]
#### soal3.go

```go
package main

import "fmt"

func polaBintang(n int, i int) {
	if i > n {
		return
	}

	for j := 0; j < i; j++ {
		fmt.Print("*")
	}
	fmt.Println()

	polaBintang(n, i+1)
}

func main() {
	var n int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&n)

	polaBintang(n, 1)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Soal 3](https://github.com/aydinzafrul21/109082500036_Aydin-Zafrul/blob/main/WEEK5/output/output-2.png)
[Program meminta input bilangan N lalu mengecek setiap angka dari 1 hingga N menggunakan rekursi. Jika suatu angka habis membagi N, maka angka tersebut dicetak sebagai faktor. Fungsi akan terus memanggil dirinya dengan nilai i bertambah hingga melebihi N lalu berhenti.]