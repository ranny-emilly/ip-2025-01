package main

import "fmt"

func main() {

	var ano int64

	fmt.Scan(&ano)

	if ano%4 != 0 {
		fmt.Println("Nao bissexto")

	} else if ano%4 == 0 && ano%100 != 0 {
		fmt.Println("Bissexto")

	} else if ano%100 == 0 && ano%400 != 0 {
		fmt.Println("Nao bissexto")

	} else if ano%400 == 0 {
		fmt.Println("Bissexto")

	}

}
