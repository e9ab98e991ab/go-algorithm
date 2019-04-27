## go常见的十大算法
1. BubbleSort(冒泡排序)
```go
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

```
2. BucketSort(桶排序)
```go

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
```

3. CountSort (计数排序)
```go

func Sort(list []int) []int{
	max := max(list)
	min := min(list)
	base := -min
	max = max - base
	min = min - base
	numbers := make([]int, max - min + 1)
	for _,value := range list{
		numbers[value + base] = numbers[value + base] + 1
	}
	var result []int
	for i,value := range numbers {
		for j:=value;j>0 && value > 0;j-- {
			result = append(result, i - base)
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

``` 
4. HeapSort(堆排序)
```go

func Sort(list []int)  {
	length := len(list)
	for {
		if length < 1 {
			break
		}
		index := length/2 -1
		for ;index>=0;index-- {
			swap(list, index, length-1)
		}
		tmp := list[0]
		list[0] = list[length - 1]
		list[length - 1] = tmp
		length--
	}
}

func swap(list []int, index int, length int)  {

	left := 2*index + 1
	right := 2*index + 2
	if left <= length && list[left] > list[index] {
		tmp := list[index]
		list[index] = list[left]
		list[left] = tmp
	}
	if right <= length && list[right] > list[index] {
		tmp := list[index]
		list[index] = list[right]
		list[right] = tmp
	}
}

```
5. InsertSort(插入排序)
```go

func Sort(list []int, left, right int)  {
	for index := left;index <= right;index++ {
		if index > 0 {
			for i:=index;i>0;i-- {
				current := list[i]
				pre := list[i-1]
				if current <= pre {
					utils.SwapGo(list, i, i-1)
				} else {
					break
				}
			}
		}
	}
}

```
6. MergeSort(合并排序)
```go

func Sort(list []int, left, right int) []int{
	return mergeSort(list[left:right-left + 1])
}

func mergeSort(list []int) []int {
	if len(list) < 2 {
		return list
	} else {
		return merge(mergeSort(list[:len(list)/2]), mergeSort(list[len(list)/2:]))

	}
}

func merge(list0, list1 []int) []int{
	var result []int
	index0 := 0
	index1 := 0
	for {
		if index0 < len(list0) && index1 < len(list1) {
			if list0[index0] < list1[index1] {
				result = append(result, list0[index0])
				index0 = index0 + 1
			} else {
				result = append(result, list1[index1])
				index1 = index1 + 1
			}
		} else {
			break
		}
	}
	if index0 < len(list0) {
		result = append(result, list0[index0:]...)
	}
	if index1 < len(list1) {
		result = append(result, list1[index1:]...)
	}
	return result
}

```
7. QuickSort(快速排序)
```go

import "github.com/go-algorithm/utils"

func Sort(list []int, left, right int)  {
	if right < left {
		return
	}
	flag := list[left]
	start := left
	end := right
	for {
		if start == end {
			break
		}
		for list[end] >= flag && end > start {
			end--
		}
		for list[start] <= flag && end > start {
			start++
		}
		if end > start {
			utils.SwapGo(list, start, end)
		}
	}
	utils.SwapGo(list, left, start)
	Sort(list, left, start - 1)
	Sort(list, start + 1, right)
}

```
8. RadixSort(基数排序)
```go

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

```
9. SelectSort(选择排序)
```go

import "github.com/go-algorithm/utils"

func Sort(list []int, left, right int)  {
	if left == right {
		return
	}
	minIndex := left
	for i:=left;i<=right;i++ {
		if list[i] <= list[minIndex] {
			minIndex = i
		}
	}
	utils.SwapGo(list, left, minIndex)
	Sort(list, left + 1, right)
}

```
10. shellsort(希尔排序)
```go
func Sort(list []int, left, right int)  {
	increment := len(list)/3 + 1
	for {
		if increment < 1 {
			break
		}
		for i:=left;i<increment;i++ {
			for j:=i+increment;j<=right;j++ {
				if list[j] < list[j-increment] {
					tmp := list[j]
					list[j] = list[j-increment]
					list[j-increment] = tmp
				}
			}
		}
		increment--
	}
}
```
main方法
```go
package main

import (
	"fmt"
	"github.com/go-algorithm/bubblesort"
	"github.com/go-algorithm/quicksort"
	"github.com/go-algorithm/selectsort"
	"github.com/go-algorithm/insertsort"
	"github.com/go-algorithm/shellsort"
	"github.com/go-algorithm/radixsort"
	"github.com/go-algorithm/heapsort"
	"github.com/go-algorithm/bucketsort"
	"github.com/go-algorithm/countsort"
	"github.com/go-algorithm/mergesort"
)

var data = []int{8, 3, 6, 9, 11, 2, 7, 23, 65, 13, 9}

func main() {
	datePrintln("桶排序")
	datePrintln("计数排序")
	datePrintln("冒泡排序")
	datePrintln("快速排序")
	datePrintln("选择排序")
	datePrintln("插入排序")
	datePrintln("希尔排序")
	datePrintln("合并排序")
	datePrintln("基数排序")
	datePrintln("堆排序")

}

func datePrintln(name string) {
	var result []int
	fmt.Println("初始数据:", data)
	switch name {
	case "桶排序":
		result = bucketsort.Sort(data)
		break
	case "计数排序":
		result = countsort.Sort(data)
		break
	case "合并排序":
		result = mergesort.Sort(data, 0, len(data)-1)
		break
	case "冒泡排序":
		bubblesort.Sort(data, 0, len(data)-1)
		result = data
		break
	case "快速排序":
		quicksort.Sort(data, 0, len(data)-1)
		result = data
		break
	case "选择排序":
		selectsort.Sort(data, 0, len(data)-1)
		result = data
		break
	case "插入排序":
		insertsort.Sort(data, 0, len(data)-1)
		result = data
		break
	case "希尔排序":
		shellsort.Sort(data, 0, len(data)-1)
		result = data
		break
	case "基数排序":
		radixsort.Sort(data)
		result = data
		break
	case "堆排序":
		heapsort.Sort(data)
		result = data
		break
	}
	fmt.Println(name+":", result)
	data = []int{8, 3, 6, 9, 11, 2, 7, 23, 65, 13, 9}
	fmt.Println("原始数据:", data, "\n")
}

```