package main

import "fmt"

func SelectionSort(slice *[]int){
	for i := range *slice {
		minIndex := i
		for j := i + 1; j < len(*slice); j++ {
			if (*slice)[j] < (*slice)[minIndex] {
				minIndex = j
			}
		}
		(*slice)[i], (*slice)[minIndex] = (*slice)[minIndex], (*slice)[i]
	}
}

func main(){
	slice := []int{4,2,3,5,6,1,9,8,7}
	SelectionSort(&slice)
	fmt.Println(slice)
}