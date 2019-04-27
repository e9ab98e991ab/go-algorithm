package bubblesort

import (
	"github.com/go-algorithm/utils"
)

func Sort(list []int, left, right int)  {

	if right == 0 {
		return
	}
	for index,num := range list {
		if index < right && num > list[index + 1] {
			utils.SwapGo(list, index, index + 1)
		}
	}
	Sort(list, left, right - 1)
}
