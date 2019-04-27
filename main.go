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

var  data = []int{8, 3, 6, 9, 11, 2, 7, 23, 65, 13, 9}

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
	/***
	 * 重置链表中数据
	 */
	data = []int{8, 3, 6, 9, 11, 2, 7, 23, 65, 13, 9}
	fmt.Println("原始数据:", data, "\n")
}
