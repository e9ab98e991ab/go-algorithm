package countsort

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
