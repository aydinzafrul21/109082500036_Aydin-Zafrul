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
