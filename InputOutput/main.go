package main

import "fmt"


func incrementNumber(nb *int) {
	*nb = *nb + 1
}

func main() {
	nb := 1
	fmt.Printf("Meu número é: %v\n", nb)
	incrementNumber(&nb)
	fmt.Printf("Meu número é: %v\n", nb)
}