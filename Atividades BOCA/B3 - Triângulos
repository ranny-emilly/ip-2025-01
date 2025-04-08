package main

import "fmt"

func main() {
	var X, Y, Z float64

	fmt.Scan(&X, &Y, &Z)

	if X+Y <= Z || X+Z <= Y || Y+Z <= X {
		fmt.Println("Nao forma triangulo")
		return
	}
	if X == Y && Y == Z {

		fmt.Println("Equilatero")

	} else if X == Y || Y == Z || Z == X {

		fmt.Println("Isosceles")

	} else if X != Y && Y != Z && Z != X {

		fmt.Println("Escaleno")

	}

}
