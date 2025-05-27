HAMMAM TIHAR ANANDA
103112400276
IF-12-05

Soal 1: Barisan Bilangan Ganjil

// Fungsi rekursif untuk mencetak bilangan ganjil hingga N
func printOdd(n, i int) {
    if i > n {
        return
    }
    fmt.Println(i)
    printOdd(n, i+2)
}

// Program utama
func main() {
    var n int
    fmt.Scan(&n)
    printOdd(n, 1)
}


Soal 2: Faktor Bilangan

// Fungsi rekursif untuk mencetak faktor bilangan N
func printFactors(n, i int) {
    if i > n {
        return
    }
    if n%i == 0 {
        fmt.Println(i)
    }
    printFactors(n, i+1)
}

// Program utama
func main() {
    var n int
    fmt.Scan(&n)
    printFactors(n, 1)
}

SSoal 3: X Pangkat Y

// Fungsi rekursif untuk menghitung x pangkat y
func power(x, y int) int {
    if y == 0 {
        return 1
    }
    return x * power(x, y-1)
}

// Program utama
func main() {
    var x, y int
    fmt.Scan(&x, &y)
    fmt.Println(power(x, y))
}


Soal 4: Fibonacci
// Fungsi rekursif untuk menghitung bilangan Fibonacci ke-n
func fibonacci(n int) int {
    if n == 0 {
        return 0
    } else if n == 1 {
        return 1
    }
    return fibonacci(n-1) + fibonacci(n-2)
}

// Program utama
func main() {
    var n int
    fmt.Scan(&n)
    fmt.Println(fibonacci(n))
}


Soal 5: Pola Bintang
// Fungsi rekursif untuk mencetak pola bintang
func printStars(n, i int) {
    if i > n {
        return
    }
    for j := 0; j < i; j++ {
        fmt.Print("*")
    }
    fmt.Println()
    printStars(n, i+1)
}

// Program utama
func main() {
    var n int
    fmt.Scan(&n)
    printStars(n, 1)
}
