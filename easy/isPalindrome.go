package easy

import "strconv"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	value := strconv.Itoa(x)
	return compareString(value)
}

func compareString(input string) bool {
	array := make([]string, len(input))

	for i := range input {
		array[i] = string(input[i])
	}

	return rec(0, len(array)-1, array)
}

func rec(start int, end int, array []string) bool {
	if start == end {
		return true
	}
	var head string
	var tail string

	head = array[start]
	tail = array[end]

	result := head == tail

	if result == true && start+1 < len(array)/2 {
		return rec(start+1, end-1, array)
	}

	return result
}
