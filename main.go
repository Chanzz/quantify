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
	low, high, err := logics.GetHistoryLowAndHigh(originalData)
	if err != nil {
		fmt.Printf("get history low and high err: %s", err.Error())
	}
	fmt.Printf("low:%g", low)
	fmt.Printf("high:%g", high)
}
