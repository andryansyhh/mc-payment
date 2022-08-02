package main

import "fmt"

func main() {
	fmt.Print("masukan kata: ")
	var str string
	fmt.Scanf("%s", &str)
	result := IsPalindrome(str)
	fmt.Println("result: ", result)
}

func IsPalindrome(input string) string {
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return "not palindrome"
		}
	}
	return "palindrome"
}
