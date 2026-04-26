package main

import (
	"fmt"
	"math"
)

type IniArr [2026]int

func isiArray(a *IniArr, n *int) {
	fmt.Print("Masukkan jumlah N: ")
	fmt.Scan(n)

	fmt.Printf("Masukkan %d isi Array: \n", *n)
	for i := 0; i < *n; i++ {
		fmt.Scan(&a[i])
	}
}

func seluruh(a IniArr, n int) {
	fmt.Print("a. isi keseluruhan Array: ")
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", a[i]) 
	}
	fmt.Println()
}

func ganjil(a IniArr, n int) {
	fmt.Print("b. Indeks Ganjil: ")
	for i := 0; i < n; i++ {
		if i%2 != 0 {
			fmt.Printf("%d ", a[i]) 
		}
	}
	fmt.Println()
}

func genap(a IniArr, n int) {
	fmt.Print("c. Indeks Genap: ")
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Printf("%d ", a[i]) 
		}
	}
	fmt.Println()
}

func kelipatan(a IniArr, n, x int) {
	fmt.Print("Masukkan x indeks kelipatan: ")
	fmt.Scan(&x)
	fmt.Printf("d. Elemen indeks kelipatan %d: ", x)
	for i := 0; i < n; i++ {
		if x != 0 && i%x == 0 {
			fmt.Printf("%d ", a[i])
		}
	}
	fmt.Println()
}

func hapusindeks(a *IniArr, n *int, y int) {
	fmt.Print("Indeks yang dihapus: ")
	fmt.Scan(&y)
	for i := y; i < *n-1; i++ {
		a[i] = a[i+1]
	}
	*n--
	fmt.Print("e. Setelah dihapus: ")
	for i := 0; i < *n; i++ {
		fmt.Printf("%d ", a[i]) 
	}
	fmt.Println()
}

func ratarata(a IniArr, n int) {
	hasil := 0
	for i := 0; i < n; i++ {
		hasil += a[i]
	}
	fmt.Printf("f. Rata-rata: %.2f\n", float64(hasil)/float64(n))
}

func simpanganbaku(a IniArr, n int) { 
	hasil := 0
	for i := 0; i < n; i++ {
		hasil += a[i]
	}
	rata := float64(hasil) / float64(n)

	var kuadrat float64
	for i := 0; i < n; i++ {
		kuadrat += math.Pow(float64(a[i])-rata, 2)
	}

	simpangan := math.Sqrt(kuadrat / float64(n))
	fmt.Printf("g. Simpangan Baku: %.2f\n", simpangan)
}

func frekuensi(a IniArr, n int) { 
	var t int
	fmt.Print("Masukkan bilangan yang dicari: ")
	fmt.Scan(&t)

	hasil := 0
	for i := 0; i < n; i++ {
		if a[i] == t {
			hasil++
		}
	}
	fmt.Printf("g. Frekuensi bilangan %d: %d\n", t, hasil)
}

func main() {
	var arr IniArr
	var jumlah int
	var x, y int

	isiArray(&arr, &jumlah)

	fmt.Println("\n===============")

	seluruh(arr, jumlah)
	ganjil(arr, jumlah)
	genap(arr, jumlah)
	kelipatan(arr, jumlah, x)
	hapusindeks(&arr, &jumlah, y)
	ratarata(arr, jumlah)
	simpanganbaku(arr, jumlah) 
	frekuensi(arr, jumlah)     
}