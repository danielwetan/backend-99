package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(Palindrome(1, 10))    // 9
	fmt.Println(Palindrome(100, 150)) // 5
	fmt.Println(Palindrome(21, 31))   // 1
}

func Palindrome(start int, end int) int {
	result := 0

	for i := start; i < end; i++ {
		if reverseValue(i) {
			result++
		}
	}

	return result
}

func reverseValue(x int) bool {
	temp := strconv.Itoa(x)
	reversed := ""

	if len(temp) < 1 {
		return false
	}

	for i := len(temp) - 1; i >= 0; i-- {
		reversed += string(temp[i])
	}

	if reversed == temp {
		return true
	}

	return false
}
