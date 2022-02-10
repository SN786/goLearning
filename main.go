package main

import (
	"fmt"
	check "github.com/SN786/goLearning/stringUtil"
)

func main() {
	var s string
	fmt.Scanln(&s)
	if check.IsPalindrome(s) {
		fmt.Println(s," is palindrome")
	} else {
		fmt.Println(s,"is not palindrome")
	}
}