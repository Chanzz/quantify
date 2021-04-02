package logics

import (
	"fmt"
	"math"
	"quantify/messages"
	"sort"
	"strconv"
)

func GetHistoryLowAndHigh(originalData *messages.OriginalData) (float64, float64) {
	var historyLow float64
	var historyHigh float64
	historyLow = math.MaxFloat64
	historyHigh = 0
	for _, i := range originalData.TimeSeries {
		if i.Low < historyLow {
			historyLow = i.Low
		}
		if i.High > historyHigh {
			historyHigh = i.High
		}
	}
	return historyLow, historyHigh
}

func Grid(ratio int, buyIn1 float64, originalData *messages.OriginalData) {
	var dateList []string
	for date := range originalData.TimeSeries {
		dateList = append(dateList, date)
	}
	sort.SliceStable(dateList, func(i, j int) bool {
		return dateList[i] < dateList[j]
	})

	// 总份额
	var stock int
	// 总成本
	var cost float64
	// 保留两位小数
	// 计算所有买入卖出点
	sellOut1, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", buyIn1*(1+float64(ratio)/100)), 64)
	buyIn2, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", buyIn1*(1-float64(ratio)/100)), 64)
	sellOut2, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", buyIn2*(1+float64(ratio)/100)), 64)
	buyIn3, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", buyIn2*(1-float64(ratio)/100)), 64)
	sellOut3, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", buyIn3*(1+float64(ratio)/100)), 64)
	//fmt.Printf("buyIn1:%g \n",buyIn1)
	//fmt.Printf("sellOut1:%g \n",sellOut1)
	//fmt.Printf("buyIn2:%g \n",buyIn2)
	//fmt.Printf("sellOut2:%g \n",sellOut2)
	//fmt.Printf("buyIn3:%g \n",buyIn3)
	//fmt.Printf("sellOut3:%g \n",sellOut3)

	for _, date := range dateList {
		if s, ok := originalData.TimeSeries[date]; ok {
			// buy
			if s.Low < buyIn1 {
				//cost=
			}
		}
	}
}
