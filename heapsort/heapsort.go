package heapsort

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
