package logics

import (
	"math"
	"quantify/messages"
	"strconv"
)

func GetHistoryLowAndHigh(originalData *messages.OriginalData) (float64, float64, error) {
	var historyLow float64
	var historyHigh float64
	historyLow = math.MaxFloat64
	historyHigh = 0
	for _, i := range originalData.TimeSeries {
		low, err := strconv.ParseFloat(i.Low, 64)
		if err != nil {
			return 0, 0, err
		}
		high, err := strconv.ParseFloat(i.High, 64)
		if err != nil {
			return 0, 0, err
		}
		if low < historyLow {
			historyLow = low
		}
		if high > historyHigh {
			historyHigh = high
		}
	}
	return historyLow, historyHigh, nil
}

func Grid(ratio int, originalData *messages.OriginalData) {

}
