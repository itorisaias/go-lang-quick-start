package main

import (
	"fmt"
)

func main() {
	fmt.Print("Mesma")
	fmt.Print(" Linha.")

	fmt.Println("Nova")
	fmt.Println("linha.")

	x := 3.1415

	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é " + xs)
	fmt.Println("O valor de x é", x, "!!!")

	fmt.Printf("O valor do x é %.2f", x)

	a := 1
	b := 1.9999
	c := false
	d := "opa"
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %.1v %t %v", a, b, b, c, d)

}
