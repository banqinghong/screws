package main

import (
	"fmt"
	"github.com/banqinghong/screws/golang/btime"
)

func main() {
	//fmt.Println("main starting")
	//s := []float64{1, 3, 3}
	//fmt.Println("sum: ", number.GetAvgOfFloatList(s))
	//bfile.ReadFileLine("/tmp/test.txt")
	timeStr := "2021-03-04 02:10:00.000"
	time, _ := btime.String2Time(timeStr)
	lastMonth := time.AddDate(0, -1, 0).Format("2006-01")
	fmt.Println(lastMonth)
}
