func maximumSum(arr [] int32) int64 {
	min := arr[0]
	for_, val := range arr {
		if val < min {
			min = val
		}
	}
	total := int64(0)
	temp := int64(min)
	for i := 0; i < len(arr); i++ {
		total = total + int64 (arr[i])
		if total > temp {
			temp = total
		}
		if total < 0 {
			total = 0
		}
	}
	return temp
}