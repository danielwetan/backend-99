package main

import (
	"fmt"
	"strconv"
)

func main() {
	test1 := "12346789" // 5
	// test2 := "101102103104105106107108109111112113" // 110
	result := Solution(test1)
	fmt.Println(result)
}

func Solution(str string) int {
	digits := checkDigits(str)
	arr := stringToArray(str, digits)
	idealSum := idealSum(arr, digits)
	missingNumber := findMissingNumber(arr, idealSum)
	return missingNumber
}

func stringToArray(str string, digits int) []string {
	temp := []string{}

	start := 0
	end := digits
	for {
		temp = append(temp, str[start:end])
		start += digits
		end += digits

		if end == len(str)+digits {
			break
		}
	}

	return temp
}

func checkDigits(str string) int {
	// is 1 digit
	a, _ := strconv.Atoi(str[0:1])
	a1, _ := strconv.Atoi(str[1:2])
	a2, _ := strconv.Atoi(str[2:3])
	if a1 == a+1 && a2 == a+2 {
		return 1
	}

	// is 10^1 digits
	b, _ := strconv.Atoi(str[0:2])
	b1, _ := strconv.Atoi(str[2:4])
	b2, _ := strconv.Atoi(str[4:6])
	if b1 == b+1 && b2 == b+2 {
		return 2
	}

	// is 10^2 digits
	c, _ := strconv.Atoi(str[0:3])
	c1, _ := strconv.Atoi(str[3:6])
	c2, _ := strconv.Atoi(str[6:9])
	if c1 == c+1 && c2 == c+2 {
		return 3
	}

	// is 10^3 digits
	d, _ := strconv.Atoi(str[0:4])
	d1, _ := strconv.Atoi(str[4:8])
	d2, _ := strconv.Atoi(str[8:12])
	if d1 == d+1 && d2 == d+2 {
		return 4
	}

	// is 10^4 digits
	e, _ := strconv.Atoi(str[0:5])
	e1, _ := strconv.Atoi(str[5:10])
	e2, _ := strconv.Atoi(str[10:15])
	if e1 == e+1 && e2 == e+2 {
		return 5
	}

	// is 10^5 digits
	f, _ := strconv.Atoi(str[0:6])
	f1, _ := strconv.Atoi(str[6:12])
	f2, _ := strconv.Atoi(str[12:18])
	if f1 == f+1 && f2 == f+2 {
		return 6
	}

	// is 10^6 digits
	g, _ := strconv.Atoi(str[0:7])
	g1, _ := strconv.Atoi(str[7:14])
	g2, _ := strconv.Atoi(str[14:21])
	if g1 == g+1 && g2 == g+2 {
		return 7
	}

	return -1
}

func idealSum(arr []string, digit int) int {
	temp := 0
	incr, _ := strconv.Atoi(arr[0])

	i := 0
	for {
		if i < len(arr) {
			temp += incr
			incr++
			i++
		} else {
			temp += incr
			break
		}
	}
	return temp
}

func findMissingNumber(arr []string, idealSum int) int {
	temp := []int{}

	for i := 0; i < len(arr); i++ {
		convert, _ := strconv.Atoi(arr[i])
		temp = append(temp, convert)
	}

	sum := calculateSum(temp)

	missingNumber := idealSum - sum
	return missingNumber
}

func calculateSum(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}
