package bstring

// 判断切片中某个元素是否存在
func IsExistString (i string, s []string) (res bool){
	for _, v := range s {
		if v == i {
			return true
		}
	}
	return false
}
