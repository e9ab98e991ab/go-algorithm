package radixsort

import (
	"strconv"
)

func Sort(list []int)  {
	baseList := make([][]int, 10)
	maxDigist := maxDigist(list)
	for i:=0;i<maxDigist;i++ {
		for _,value := range list {
			baseList[getDigist(value, i)] = append(baseList[getDigist(value, i)], value)
		}

		j := 0
		for index,value :=range baseList {
			if len(value) > 0 {
				for _,v := range value {
					list[j] = v
					j++
				}
			}
			baseList[index] = nil
		}
	}



}

func maxDigist(list []int) int {
	maxDigist := 1
	for _,value := range list {
		if len(strconv.Itoa(value)) > maxDigist {
			maxDigist = len(strconv.Itoa(value))
		}
	}
	return maxDigist
}

func getDigist(number int, index int) int  {
	strNum := strconv.Itoa(number)
	if index > len(strNum) - 1 {
		return 0
	}
	index = len(strNum) - 1 - index
	//fmt.Println("index = ", index)
	result,error := strconv.Atoi(string(strNum[index]))
	if error != nil {

		return -1
	} else {
		return result
	}
}
