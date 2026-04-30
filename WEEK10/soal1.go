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
