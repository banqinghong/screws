package number

// 获取列表最大值
func GetMaxIntNumOfList (s []int) (maxNum int){
	if len(s) == 0 {
		return 0
	}
	for _, v := range s {
		if v > maxNum {
			maxNum = v
		}
	}
	return maxNum
}

// 求和
func GetSumOfIntList(s []int) (sum int)  {
	if len(s) == 0 {
		return 0
	}
	sum = 0
	for _, v := range s {
		sum += v
	}
	return sum
}

// 求平均值
func GetAvgOfIntList(s []int) (avg float64)  {
	if len(s) == 0 {
		return 0
	}
	sum := 0
	for _, v := range s {
		sum += v
	}
	avg = float64(sum) / float64(len(s))
	return Float2float(avg, 4)
}
