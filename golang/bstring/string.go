package bstring

// 判断切片中某个元素是否存在
func IsExistString(i string, s []string) bool {
	if len(s) == 0 {
		return false
	}
	if i == "" {
		return true
	}
	for _, v := range s {
		if v == i {
			return true
		}
	}
	return false
}
