package main

import "fmt"

func main() {
    var n int
    fmt.Print("Digite o tamanho do array: ")
    fmt.Scan(&n)

    arr := make([]float64, n)
    soma := 0.0

    for i := 0; i < n; i++ {
        fmt.Printf("Digite o valor %d: ", i+1)
        fmt.Scan(&arr[i])
        soma += arr[i]
    }

    fmt.Printf("Soma dos valores = %.2f\n", soma)
}
