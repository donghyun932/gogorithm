package main

import "fmt"

func CombSort(slice *[]int) {
	gap := len(*slice)
	shrink := 1.5
	sorted := false

	for !sorted {
		gap = int( float64(gap) / shrink)
		if gap > 1{
			sorted = false
		} else {
			gap = 1
			sorted = true
		}

		for i:=0; i + gap < len(*slice); i++ {
			if (*slice)[i] > (*slice)[i + gap] {
				(*slice)[i], (*slice)[i + gap] = (*slice)[i+gap], (*slice)[i]
				sorted = false
			}
		}
	}
}

func main(){
	slice := []int{4,2,3,5,6,1,9,8,7}
	CombSort(&slice)
	fmt.Println(slice)
}