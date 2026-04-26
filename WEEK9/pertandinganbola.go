package main

import "fmt"

func gol(menang [100]int, goalA, goalB int, klubA, klubB string) {
	var n int
	
	for i := 0; i < 100; i++ {
		fmt.Printf("Pertandingan %d: ", i+1)
		fmt.Scan(&goalA, &goalB)
		
		if goalA < 0 || goalB < 0 {
			break
		} else if goalA > goalB {
			menang[i] = 1
		} else if goalB > goalA {
			menang[i] = 2
		} else if goalA == goalB {
			menang[i] = 0
		}
		n++
	}
	for i := 0; i < n; i++ {
		if menang[i] == 1 {
			fmt.Printf("Hasil %d: %s\n", i+1, klubA)
		} else if menang[i] == 2 {
			fmt.Printf("Hasil %d: %s\n", i+1, klubB)
		} else if menang[i] == 0 {
			fmt.Printf("Hasil %d: Draw\n", i+1)
			}
		}
		fmt.Println("Pertandingan selesai")
	}
	

func main() {
	var klubA, klubB string
	var menang [100]int
	var goalA, goalB int
	
	fmt.Print("Klub A: ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B: ")
	fmt.Scan(&klubB)
	
	fmt.Println()
	
	gol(menang, goalA, goalB, klubA, klubB)
}