package main

import (
	"fmt"
	"quantify/logics"
)

func main() {
	originalData, err := logics.GetOriginalDataByHttp("512980.SHH")
	if err != nil {
		fmt.Printf("get original data by http err: %s", err.Error())
		return
	}
	//low, high := logics.GetHistoryLowAndHigh(originalData)
	//fmt.Printf("low:%g", low)
	//fmt.Printf("high:%g", high)
	logics.Grid(7, 0.860, originalData)
}
