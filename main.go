package main

import (
	"strings"
)

func main() {
	//var slice = []int{7, 9, 20, 50, 55, 80, 90, 777, 5555}
	//fmt.Println("\n--- Unsorted --- \n\n", slice)

	//sort.QuickSort(slice)
	//result := lib.BinarySearch(slice, 900)
	//fmt.Println("Result: ", result)
	//fmt.Println("\n--- Sorted ---\n\n", slice, "\n")

	/*t := lib.TreeNode{Val: 8, Left: nil, Right: nil}

	  t.Insert(1)
	  t.Insert(200)
	  t.Insert(-5)
	  t.Insert(12)
	  t.Insert(-63)
	  t.Insert(13)
	  t.Insert(0)

	  t.PrintInorder()*/

	/*var array []int
	array = []int{1, 2}

	array = append(array, 23)

	fmt.Println(array)*/

}

// Generates a slice of size, size filled with random numbers

func FindOdd(seq []int) int {
	t := make(map[int]int)
	var answer int

	for _, elem := range seq {
		t[elem]++
	}

	for _, elem := range seq {
		if t[elem]%2 != 0 {
			answer = elem
			break
		}
	}
	return answer
}

func Find(a []string, x string) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return len(a)
}

func Solution(str string) []string {
	var result []string
	for i := 0; i < len(str); i += 2 {
		result = append(result, str[i:i+2])
	}
	return result
}

func CamelCase(s string) string {
	resultString := ""
	words := strings.Fields(s)

	for _, word := range words {
		resultString += strings.Title(word)
	}
	return resultString
}

func FindShort(s string) int {
	var minLen int
	var currentLen int
	flag := true

	words := strings.Fields(s)
	for _, word := range words {
		currentLen = len(word)
		if flag {
			minLen = currentLen
			flag = false
		}
		if currentLen < minLen {
			minLen = currentLen
		}
	}
	return minLen
}

func DuplicateEncode(word string) string {
	t := make(map[rune]int)
	word = strings.ToLower(word)
	newString := ""

	for _, c := range word {
		t[c]++
	}

	for _, c := range word {
		if t[c] == 1 {
			newString += "("
		} else {
			newString += ")"
		}
	}

	return newString
}

/*func myFindMax(array []int) int {
	var maxElem = 0
	for i := 1; i < len(array); i++ {
		if array[i] > maxElem {
			maxElem = array[i]
		}
	}
	return maxElem
}

func myAppend (array []int, number int)  []int {
	return append(array, number)
}*/
