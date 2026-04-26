package main

import "fmt"

func main() {

	var num1, num2, num3 int64

	fmt.Scanf("%d %d %d", &num1, &num2, &num3)

	if num1 > num2 && num2 > num3 {
		fmt.Printf("%d %d %d", num3, num2, num1)

	} else if num1 > num3 && num3 > num2 {
		fmt.Printf("%d %d %d", num2, num3, num1)

	} else if num2 > num1 && num1 > num3 {
		fmt.Printf("%d %d %d", num3, num1, num2)

	} else if num2 > num3 && num3 > num1 {
		fmt.Printf("%d %d %d", num1, num3, num2)

	} else if num3 > num2 && num2 > num1 {
		fmt.Printf("%d %d %d", num1, num2, num3)

	} else {
		fmt.Println(num2, num1, num3)
	}

}
