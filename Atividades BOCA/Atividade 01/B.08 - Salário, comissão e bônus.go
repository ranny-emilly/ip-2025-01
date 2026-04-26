package main

import "fmt"

func main() {

	var vB, bonus float64

	fmt.Scan(&vB)

	comissao := vB * 0.09

	if vB > 15000.00 {
		bonus = 800.00
	} else {
		bonus = 0.00
	}

	salarioFinal := 500.00 + comissao + bonus

	fmt.Printf("%f", salarioFinal)

}
