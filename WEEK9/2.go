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

	// f & g. Rata-rata dan Standar Deviasi [cite: 384, 385]
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
