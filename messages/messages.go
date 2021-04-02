package messages

type OriginalData struct {
	MetaData   MetaData              `json:"Meta Data"`
	TimeSeries map[string]TimeSeries `json:"Time Series (Daily)"`
}

type MetaData struct {
	Information string `json:"1. Information"`
	Symbol      string `json:"2. Symbol"`
	LastRefresh string `json:"3. Last Refreshed"`
	OutputSize  string `json:"4. Output Size"`
	TimeZone    string `json:"5. Time Zone"`
}

type TimeSeries struct {
	Open             float64 `json:"1. open,string"`
	High             float64 `json:"2. high,string"`
	Low              float64 `json:"3. low,string"`
	Close            float64 `json:"4. close,string"`
	AdjustedClose    float64 `json:"5. adjusted close,string"`
	Volume           string  `json:"6. volume"`
	DividendAmount   string  `json:"7. dividend amount"`
	SplitCoefficient string  `json:"8. split coefficient"`
}

type MacdOriginalData struct {
	MetaData          MacdMetaData                 `json:"Meta Data"`
	TechnicalAnalysis map[string]TechnicalAnalysis `json:"Technical Analysis: MACD"`
}

type MacdMetaData struct {
	Symbol        string `json:"1: Symbol"`
	Indicator     string `json:"2: Indicator"`
	LastRefreshed string `json:"3: Last Refreshed"`
	Interval      string `json:"4: Interval"`
	FastPeriod    int    `json:"5.1: Fast Period"`
	SlowPeriod    int    `json:"5.2: Slow Period"`
	SignalPeriod  int    `json:"5.3: Signal Period"`
	SeriesType    string `json:"6: Series Type"`
	TimeZone      string `json:"7: Time Zone"`
}

type TechnicalAnalysis struct {
	MACD       string `json:"MACD"`
	MACDHist   string `json:"macd_hist"`
	MACDSignal string `json:"macd_signal"`
}
