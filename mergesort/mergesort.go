package mergesort

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
