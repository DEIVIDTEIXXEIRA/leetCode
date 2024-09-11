package main

import (
	"fmt"
	"strconv"
)

func main() {
	teste := 0
	result := IsPalindrome(teste)
	fmt.Printf("result: %v\n", result)
}

// minha solução demorou muito mais e utilizou alocação de memória
func IsPalindrome(x int) bool {
	s := strconv.Itoa(x)

	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	ss := string(runes)

	return ss == s
}

// solução fixada no desafio bem mais rapida e sem alocação de memória
func IsPalindrome2(num int) bool {
	if num < 0 {
		return false
	}
	x := num
	reversed := 0
	for x != 0 {
		reversed = 10*reversed + x%10
		x /= 10
	}
	return (reversed == num)
}
