package main

import (
	"fmt"
	"math"
)
type  T struct{
	x, y float64
}
type ling struct {
	cx, cy, r float64
	
}

func jarak(a, b, c, d float64) float64 {
	hasil := math.Sqrt((a - c)*(a - c) + (b - d)*(b - d))
	return hasil
}

func didalam(cx, cy, r, x, y float64) bool {
	if jarak(cx, cy, x, y) < r {
		return true
	}
	return false
}

func main() {
	var ling1, ling2 ling
	var titik T

	fmt.Print("Nilai Lingkaran 1 cx1, cy1, r1: ")
	fmt.Scanln(&ling1.cx, &ling1.cy, &ling1.r)
	fmt.Print("Nilai Lingkaran 2 cx2, cy2, r2: ")
	fmt.Scanln(&ling2.cx, &ling2.cy, &ling2.r)
	fmt.Print("Titik x, y: ")
	fmt.Scanln(&titik.x, &titik.y)

	dalam1 := didalam(ling1.cx, ling1.cy, ling1.r, titik.x, titik.y)
	dalam2 := didalam(ling2.cx, ling2.cy, ling2.r, titik.x, titik.y)
	if dalam1 && dalam2 {
		fmt.Println("Titik didalam lingkaran 1 dan 2")
	} else if dalam1 {
		fmt.Println("Titik didalam lingkaran 1")
	} else if dalam2 {
		fmt.Println("Titik didalam lingkaran 2")
	} else {
		fmt.Println("Titik diluar lingkaran 1 dan 2")
	}
}