package main

import (
	"fmt"
)

func fatorial(n uint) uint {
	if n == 0 {
		return 1
	}
	return n * fatorial(n-1)
}

func main() {
	resultado := fatorial(5)
	fmt.Println(resultado)

	resultado2 := fatorial(0)
	fmt.Println(resultado2)
}
