package bucketsort

import (
	"github.com/go-algorithm/quicksort"
)

func Sort(list []int)  []int{
	max := max(list)
	min := min(list)
	base := 0
	if min < 0 {
		base = -min
	} else {
		base = min
	}
	max = (max + base)/10
	min = (min + base)/10
	bucket := make([][]int, max - min + 1)
	var result []int
	for _,value := range list {
		i := (int)((value+base)/10)
		bucket[i] = append(bucket[i], value)
	}

	for _,value := range bucket {
		if len(value) > 0 {
			quicksort.Sort(value, 0, len(value)-1)
		}
	}

	for _,value := range bucket {
		if len(value) > 0 {
			for _,v := range value {
				result = append(result,v)
			}
		}
	}
	return result
}

func max(list []int) int  {
	max := list[0]
	for _,value := range list {
		if value > max {
			max = value
		}
	}
	return max
}

func min(list []int) int  {
	min := list[0]
	for _,value := range list {
		if value < min {
			min = value
		}
	}
	return min
}
