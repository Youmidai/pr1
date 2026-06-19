package main

import (
	"fmt"
)

func MergeSort(arr []int) []int {
	mid := len(arr) / 2

	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))

	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

func main() {
	arr := []int{542, -565, 531, -294, -56, 14, 270, -51, -914, 605, -117, -768, 331, 708, -603, 84, -548, 579, 434, 751, 592, -349, 408, -602, 721, 909, 170, -432, -970, -171, -972, 316, 405, -676, -929, -795, -682, -646, 46, -609, -84, 180, -158, -662, -384, 854, -721, 39, 180, -197, -818, -946, -529, -555, -36, -853, -322, 540, -936, -919, 473, 978, 782, 586, 869, 333, -977, -548, -789, 988, -393, 807, -609, 997, 824, -480, -205, -576, 856, 494, 131, 40, -601, 467, 221, -640, 34, -220, 482, 948, 523, -27, -771, -914, 438, 957, 205, -411, -749, -723}

	sorted := MergeSort(arr)
	fmt.Println(sorted)
}
