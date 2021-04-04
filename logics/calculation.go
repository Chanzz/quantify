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

// 无限资金模式(梭哈)
func ShowHandGrid(ratio int, buyIn float64, originalData *messages.OriginalData) {
	var dateList []string
	for date := range originalData.TimeSeries {
		dateList = append(dateList, date)
	}
	// 排序
	sort.SliceStable(dateList, func(i, j int) bool {
		return dateList[i] < dateList[j]
	})

	// 总份额
	var stock float64
	// 总成本
	var cost float64
	// 总收入
	var income float64
	// 买入点
	buyInPoint := buyIn
	// 卖出点
	var sellOutPoint float64
	// 买入份额
	var baseStock float64
	baseStock = 10000
	// 买入次数
	var buyTimes int
	// 上次买入成本
	var lastCost float64
	// 上次买入份额
	var lastStock float64
	for _, date := range dateList {
		if s, ok := originalData.TimeSeries[date]; ok {
			// 达到买入点
			if s.Low < buyInPoint {
				// 计算总成本
				lastCost = buyInPoint * baseStock
				cost += lastCost
				// 计算总份额
				lastStock = baseStock
				stock += baseStock
				// 计算下次买入份额
				baseStock = getSellOut(baseStock, ratio)
				// 计算下次卖出点
				sellOutPoint = getSellOut(buyInPoint, ratio)
				// 计算下次买入点
				buyInPoint = getBuyIn(buyIn, ratio)
				buyTimes++
				// t+1卖出
				if buyTimes == 1 {
					continue
				}
			}
			// 达到卖出点
			if s.High > sellOutPoint {
				// 曾经买入
				if buyTimes > 0 {
					buyTimes--
					// 留利润
					stock -= getSellOutStock(lastStock, lastCost, sellOutPoint)
					income += sellOutPoint * getSellOutStock(lastStock, lastCost, sellOutPoint)
				}
			}
			fmt.Printf("date:%s data:%+v \n", date, s)
			fmt.Printf("stock:%g \n", stock)
			fmt.Printf("cost:%g \n", cost)
			fmt.Printf("income:%g \n", income)
		}

	}

}

func getBuyIn(buyIn float64, ratio int) float64 {
	nextBuyIn, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", buyIn*(1-float64(ratio)/100)), 64)
	return nextBuyIn
}

func getSellOut(buyIn float64, ratio int) float64 {
	nextSellOut, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", buyIn*(1+float64(ratio)/100)), 64)
	return nextSellOut
}

// 计算留下份额
func getSellOutStock(lastStock, lastCost, sellOutPoint float64) float64 {
	income := lastStock * sellOutPoint
	return lastStock - ((income - lastCost) / sellOutPoint)
}
