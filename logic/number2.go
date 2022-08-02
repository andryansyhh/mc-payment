package main

import (
	"strings"
)

// func main() {
// 	fmt.Print("masukan angka: ")
// 	var input int
// 	fmt.Scanf("%d", &input)
// 	mcPayment := McPayment(input)
// 	fmt.Println(mcPayment)
// }

func McPayment(i int) (res string) {
	mc := "MC"
	pay := "Pay"
	ment := "Ment"

	if i > 100 {
		return "invalid input"
	}

	if i%25 == 0 {
		return mc
	} else if i%40 == 0 {
		return pay
	} else if i%60 == 0 {
		return ment
	} else if i%99 == 0 {
		return strings.ToUpper(mc + pay + ment)
	}
	return
}
