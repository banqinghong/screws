package number

import (
	"errors"
	"fmt"
	"strconv"
)

// 小数精度转换
func Float2float (num float64, precision int) float64 {
	formatString := fmt.Sprintf("%%.%vf", precision)
	floatNum, _ := strconv.ParseFloat(fmt.Sprintf(formatString, num), 64)
	return floatNum
}

// 获取列表最大值
func GetMaxFloat64NumOfList (s []float64) (maxNum float64, err error){
	if len(s) == 0 {
		return 0, errors.New("no number found")
	}
	for _, v := range s {
		if v > maxNum {
			maxNum = v
		}
	}
	return maxNum, nil
}

// 求和
func GetSumOfFloatList(s []float64) (sum float64)  {
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
func GetAvgOfFloatList(s []float64) (avg float64)  {
	if len(s) == 0 {
		return 0
	}
	sum := 0.0
	for _, v := range s {
		sum += v
	}
	avg = sum / float64(len(s))
	return Float2float(avg, 4)
}

// 判断切片中某个元素是否存在
func IsExistFloat (i float64, s []float64) (res bool){
	for _, v := range s {
		if v == i {
			return true
		}
	}
	return false
}
