package main

import "fmt"

// Digunakan untuk menghitung nilai faktorial dari n
func Faktorial(n int) int {
	var hasil int

	if n == 0 {
		return 1
	}

	hasil = 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}

	return hasil
}

// Digunakan untuk permutasi
func Permutasi(n, r int) int {
	var hasil int
	hasil = Faktorial(n) / Faktorial(n-r)

	return hasil
}

// Fungsi yang digunakan untuk mengombinasikan 
func Kombinasi(n, r int) int {
	var hasil int
	hasil = Faktorial(n) / (Faktorial(r) * Faktorial(n-r))
	return hasil
}

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	if a >= c && b >= d {
		fmt.Println(Permutasi(a, c), Kombinasi(a, c))
		fmt.Println(Permutasi(b, d), Kombinasi(b, d))
	}
}
