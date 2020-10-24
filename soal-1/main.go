package main

import (
	"fmt"
	"strings"
)

func main() {
	result := Solution("3A13 5X19 9Y20 2C18 1N20 3N20 1M21 1F14 9A21 3N21 0E13 5G14 8A23 9E22 3N14")
	fmt.Println(result)
}

func Solution(input string) string {
	arr := stringToArray(input)
	base := addPointToData(arr)
	sort := selectionSort(base)
	final := checkIfBookDoubled(sort)
	return arrayToString(final)
}

func stringToArray(str string) []string {
	result := strings.Fields(str)
	return result
}

func arrayToString(arr []string) string {
	var result, temp string
	for i := 0; i < len(arr); i++ {
		if arr[i] == arr[len(arr)-1] {
			temp = arr[i]
		} else {
			temp = arr[i] + " "
		}
		result += temp
	}
	return result
}

func addPointToData(arr []string) []string {
	for i := 0; i < len(arr); i++ {
		// book height min 13 & max 24
		switch arr[i][2:4] {
		case "24":
			arr[i] = "24-" + arr[i]
		case "23":
			arr[i] = "23-" + arr[i]
		case "22":
			arr[i] = "22-" + arr[i]
		case "21":
			arr[i] = "21-" + arr[i]
		case "20":
			arr[i] = "20-" + arr[i]
		case "19":
			arr[i] = "19-" + arr[i]
		case "18":
			arr[i] = "18-" + arr[i]
		case "17":
			arr[i] = "17-" + arr[i]
		case "16":
			arr[i] = "16-" + arr[i]
		case "15":
			arr[i] = "15-" + arr[i]
		case "14":
			arr[i] = "14-" + arr[i]
		case "13":
			arr[i] = "13-" + arr[i]
		default:
			break
		}

		switch arr[i][3:4] {
		case "6":
			arr[i] = "10" + arr[i]
		case "7":
			arr[i] = "09" + arr[i]
		case "0":
			arr[i] = "08" + arr[i]
		case "9":
			arr[i] = "07" + arr[i]
		case "4":
			arr[i] = "06" + arr[i]
		case "8":
			arr[i] = "05" + arr[i]
		case "1":
			arr[i] = "04" + arr[i]
		case "2":
			arr[i] = "03" + arr[i]
		case "5":
			arr[i] = "02" + arr[i]
		case "3":
			arr[i] = "01" + arr[i]
		default:
			break
		}
	}

	return arr
}

func selectionSort(arr []string) []string {
	lowest := 0
	for i := 0; i < len(arr); i++ {
		lowest = i
		for j := i + 1; j < len(arr); j++ {
			if arr[j][0:4] > arr[lowest][0:4] {
				lowest = j
			}
		}

		if i != lowest {
			temp := arr[i]
			arr[i] = arr[lowest]
			arr[lowest] = temp
		}
	}

	// remove helper point value
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i][5:]
	}

	return arr
}

func checkIfBookDoubled(arr []string) []string {
	temp := 0
	for i := 0; i < len(arr); i++ {
		for j := 1; j < len(arr); j++ {
			if arr[i][0:2] == arr[j][0:2] {
				if temp < 2 {
					temp++
				} else {
					removeArrayItem(arr, j)
				}
			}
		}
		temp = 0
	}

	if len(arr) > 0 {
		arr = arr[:len(arr)-1]
	}

	return arr
}

func removeArrayItem(slice []string, i int) []string {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
