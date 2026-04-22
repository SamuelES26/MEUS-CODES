package main

import "fmt"

func main() {
    var x, n int
    fmt.Print("Digite x: ")
    fmt.Scan(&x)
    fmt.Print("Digite o expoente: ")
    fmt.Scan(&n)

    resultado := 1
    for i := 0; i < n; i++ {
        resultado *= x
    }

    fmt.Printf("%d^%d = %d\n", x, n, resultado)
}
