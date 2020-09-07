package main

import (
    "fmt"
    "github.com/banqinghong/screws/golang/number"
)

func main(){
    fmt.Println("main starting")
    s := []float64{1, 3, 3}
    fmt.Println("sum: ", number.GetAvgOfFloatList(s))
}
