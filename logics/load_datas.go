package logics

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"quantify/messages"
)

const dataUrl = "https://www.alphavantage.co/query?function=TIME_SERIES_DAILY_ADJUSTED&symbol=%s&outputsize=full&apikey=xxxxx"
const MacdDataUrl = "https://www.alphavantage.co/query?function=MACD&symbol=%s&interval=daily&series_type=close&apikey=xxxxxx"

func GetOriginalDataByHttp(symbol string) (*messages.OriginalData, error) {
	client := &http.Client{}
	url := fmt.Sprintf(dataUrl, symbol)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("new request err: %s", err.Error())
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("do request err: %s", err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("io readall err: %s", err.Error())
		return nil, err
	}

	var originData messages.OriginalData
	err = json.Unmarshal(body, &originData)
	if err != nil {
		fmt.Printf("json unmarshal err: %s", err.Error())
		return nil, err
	}
	return &originData, nil
}

func formatOriginData(data *messages.OriginalData) *messages.OriginalData {
	return data
}
