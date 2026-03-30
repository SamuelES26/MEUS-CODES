package main

import "fmt"

func main() {
	var num int64
		fmt.Println("Insira o número:")
		fmt.Scan(&num)
			if num%2 == 0 {
				fmt.Println("O número é Par")
			} else {
				fmt.Println("O número e Ímpar")
			}
}
