package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiarray(t *tabel, n *int) { 
	var kata rune
	*n = 0
	fmt.Print("Masukkan kata: ")
	for {
		fmt.Scanf("%c", &kata)
		if kata == '.' || *n >= NMAX { 
			break
		}
		if kata != '\n' { 
			t[*n] = kata
			*n++
		}
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c ", t[i]) 
	}
	fmt.Println()
}

func balikArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

func palindrome(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	isiarray(&tab, &m) 

	fmt.Print("Teks : ")
	cetakArray(tab, m)

	isPalindrome := palindrome(tab, m)

	balikArray(&tab, m)

	fmt.Print("Reverse teks : ")
	cetakArray(tab, m)

	fmt.Printf("Palindrome ? %t\n", isPalindrome)
}