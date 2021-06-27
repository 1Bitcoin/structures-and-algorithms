package lib

import "fmt"

func BinarySearch(array []int, target int) bool {
	if len(array) > 0 {
		left := 0
		right := len(array) - 1
		middle := len(array) / 2

		if array[middle] == target {
			return true

		}

		if array[middle] > target {

			right = middle - 1
		}

		if array[middle] < target {
			left = middle + 1
		}

		return BinarySearch(array[left:right+1], target)

	} else {
		fmt.Println("Элемент не найден!")
		return false
	}
}
