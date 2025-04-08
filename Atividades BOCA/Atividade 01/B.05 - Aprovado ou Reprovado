package main

import "fmt"

func main() {

	var nota1, nota2, nota3 float64
	var faltas int
	aulas := 64

	fmt.Scanf("%f %f %f %d", &nota1, &nota2, &nota3, &faltas)

	if float64(faltas) > 0.25*float64(aulas) {

		fmt.Println(" Reprovado por falta")
		return
	}

	media := (nota1 + nota2 + nota3) / 3

	if media >= 7 {
		fmt.Println("Aprovado")
	} else if 5 <= media && media < 7 {
		fmt.Println("Prova final")
	} else {
		fmt.Println("Reprovado")
	}

}
