package shellsort

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
