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
