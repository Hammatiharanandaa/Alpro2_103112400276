package main

import "fmt"

// Fungsi untuk menyimpan rumus f
func f(x int) int {
	var rumus int
	rumus = x * x
	return rumus
}

// Fungsi untuk menyimpan rumus g
func g(x int) int {
	var rumus int
	rumus = x - 2
	return rumus
}

// Fungsi untuk menyimpan rumus h
func h(x int) int {
	var rumus int
	rumus = x + 1
	return rumus
}

func main() {

	var x1, x2, x3, fogoh, gohof, hofog int
	fmt.Scan(&x1, &x2, &x3)

	fogoh = f(g(h(x1)))
	gohof = g(h(f(x2)))
	hofog = h(f(g(x3)))

	fmt.Println(fogoh, gohof, hofog)

}