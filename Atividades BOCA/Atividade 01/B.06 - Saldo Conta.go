package main

import "fmt"

func main() {

	var cc int64
	var limite, saldoInicial, depositos, retiradas, saldo float64

	fmt.Scanf("%d %f %f %f %f", &cc, &limite, &saldoInicial, &depositos, &retiradas)

	saldo = saldoInicial + depositos - retiradas

	if limite-saldo >= 0 {
		fmt.Printf("Saldo: %f", saldo)

	} else {
		fmt.Printf("Credito excedido: %f", saldo)

	}

}
