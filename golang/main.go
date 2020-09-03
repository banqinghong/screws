package main

import (
    "fmt"
    "github.com/banqinghong/screws/golang/number"
    )

func main(){
    fmt.Println("main starting")
    s := []float64{1.1, 2.2, 5, 6}
    maxNum, err := number.GetMaxNumOfFloat64(s)
    if err != nil {
        fmt.Println("err:", err)
    }else {
        fmt.Println("maxNum is ", maxNum)
    }
}
