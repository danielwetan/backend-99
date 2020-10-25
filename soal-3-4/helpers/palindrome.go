package helpers

import "strconv"

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
