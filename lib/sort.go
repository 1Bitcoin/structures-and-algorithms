package lib

import (
	"math/rand"
	"time"
)

func GenerateSlice(size int) []int {

	slice := make([]int, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func QuickSort(ar []int) {
	if len(ar) < 2 {
		return
	}

	split := partition(ar)

	QuickSort(ar[:split])
	QuickSort(ar[split:])
}

func partition(ar []int) int {
	pivot := ar[len(ar) / 2]

	left := 0
	right := len(ar) - 1

	for left < right {
		for ar[left] < pivot {
			left++
		}

		for ar[right] > pivot {
			right--
		}

		ar[left], ar[right] = ar[right], ar[left]
	}

	return left
}

func BubbleSort(array []int) []int {
	for i, _ := range array {
		for j, _ := range array {
			if array[i] < array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}

	return array
}