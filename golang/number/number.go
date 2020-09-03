package number

import "errors"

func GetMaxNumOfInt (s []int) (maxNum int, err error){
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

func GetMaxNumOfFloat64 (s []float64) (maxNum float64, err error){
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



