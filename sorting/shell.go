package main

import "fmt"

func ShellSort(slice *[]int) {
	num := len(*slice)
	for i := num / 2; i > 0; i = i / 2 {
		for j := i; j < num; j++ {
			for k := j - i; k >= 0; k = k - i {
				if (*slice)[k+i] >= (*slice)[k] {
					break
				} else {
					(*slice)[k], (*slice)[k+i] = (*slice)[k+i], (*slice)[k]
				}
			}
		}
	}
}

func main() {
	slice := []int{4, 2, 3, 5, 6, 1, 9, 8, 7}
	ShellSort(&slice)
	fmt.Println(slice)
}
