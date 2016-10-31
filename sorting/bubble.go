package main

import "fmt"

func BubbleSort(slice *[]int) {
	for i := range *slice {
		for j := len(*slice) - 1; j > i; j-- {
			if (*slice)[j] > (*slice)[j-1] {
				(*slice)[j], (*slice)[j-1] = (*slice)[j-1], (*slice)[j]
			}
		}
	}
}

func main() {
	slice := []int{4, 2, 3, 5, 6, 1, 9, 8, 7}
	BubbleSort(&slice)
	fmt.Println(slice)
}
