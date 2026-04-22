package main

import "fmt"

func main() {
    var n int
    fmt.Print("Digite um número decimal: ")
    fmt.Scan(&n)

    bin := ""
    for n > 0 {
        resto := n % 2
        bin = fmt.Sprintf("%d%s", resto, bin)
        n /= 2
    }

    fmt.Println("Binário:", bin)
}
