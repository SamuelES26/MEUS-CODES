package main

import "fmt"

func main() {
    var n int
    fmt.Print("Digite o tamanho do array: ")
    fmt.Scan(&n)

    arr := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Printf("Digite o elemento %d: ", i+1)
        fmt.Scan(&arr[i])
    }

    // Inversão
    for i := 0; i < n/2; i++ {
        arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
    }

    fmt.Println("Array invertido:", arr)
}
