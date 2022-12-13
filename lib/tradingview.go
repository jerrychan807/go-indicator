package lib

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/jerrychan807/go-indicator/common"
)

const (
	FULL_API_URL        = "https://scanner.tradingview.com/crypto/scan"
	INTERVAL_1_MINUTE   = "1m"
	INTERVAL_5_MINUTES  = "5m"
	INTERVAL_15_MINUTES = "15m"
	INTERVAL_1_HOUR     = "1h"
	INTERVAL_4_HOURS    = "4h"
	INTERVAL_1_DAY      = "1d"
	INTERVAL_1_WEEK     = "1W"
	INTERVAL_1_MONTH    = "1M"
)

var (
	recommendsList  = []string{"Recommend.Other", "Recommend.All", "Recommend.MA"}
	oscillatorsList = []string{"RSI", "RSI[1]", "Stoch.K", "Stoch.D", "Stoch.K[1]", "Stoch.D[1]", "CCI20", "CCI20[1]", "ADX", "ADX+DI", "ADX-DI", "ADX+DI[1]", "ADX-DI[1]", "AO", "AO[1]", "Mom", "Mom[1]", "MACD.macd", "MACD.signal", "Rec.Stoch.RSI", "Stoch.RSI.K", "Rec.WR", "W.R", "Rec.BBPower", "BBPower", "Rec.UO", "UO", "close"}
	maList          = []string{"EMA5", "SMA5", "EMA10", "SMA10", "EMA20", "SMA20", "EMA30", "SMA30", "EMA50", "SMA50", "EMA100", "SMA100", "EMA200", "SMA200"}
	maSimpleList    = []string{"Rec.Ichimoku", "Ichimoku.BLine", "Rec.VWMA", "VWMA", "Rec.HullMA9", "HullMA9", "open", "P.SAR", "BB.lower", "BB.upper", "AO[2]", "volume", "change", "low", "high"}
)

func concatAppend(slices [][]string) []string {
	var tmp []string
	for _, s := range slices {
		tmp = append(tmp, s...)
	}
	return tmp
}

type Data struct {
	Symbols struct {
		Tickers []string `json:"tickers"`
		Query   struct {
			Types []string `json:"types"`
		} `json:"query"`
	} `json:"symbols"`
	Columns []string `json:"columns"`
}

// PrepareData prepare payload for request
func PrepareData(symbol, interval string, indicators []string) ([]byte, error) {
	// Default, 1 Day
	dataInterval := ""

	if interval == INTERVAL_1_MINUTE {
		// 1 Minute
		dataInterval = "|1"
	} else if interval == INTERVAL_5_MINUTES {
		// 5 Minutes
		dataInterval = "|5"
	} else if interval == INTERVAL_15_MINUTES {
		// 15 Minutes
		dataInterval = "|15"
	} else if interval == INTERVAL_1_HOUR {
		// 1 Hour
		dataInterval = "|60"
	} else if interval == INTERVAL_4_HOURS {
		// 4 Hour
		dataInterval = "|240"
	} else if interval == INTERVAL_1_WEEK {
		// 1 Week
		dataInterval = "|1W"
	} else if interval == INTERVAL_1_MONTH {
		// 1 Month
		dataInterval = "|1M"
	} else {
		if interval != INTERVAL_1_DAY {
			fmt.Println("Interval is empty or not valid, defaulting to 1 day.")
			// Default, 1 Day
			dataInterval = ""
		}
	}

	//indicators := concatAppend([][]string{recommendsList, oscillatorsList, maList, maSimpleList})

	data := Data{}
	data.Symbols.Tickers = []string{symbol}
	for _, ind := range indicators {
		data.Columns = append(data.Columns, fmt.Sprintf("%s%s", ind, dataInterval))
	}
	return json.Marshal(data)
}

func GetTradingViewResponse(screener, exchange, symbol, interval string) string {
	indicators := concatAppend([][]string{recommendsList, oscillatorsList, maList, maSimpleList})
	payload, _ := PrepareData(fmt.Sprintf("%s:%s", exchange, symbol), interval, indicators)
	//fmt.Println(string(payload))

	c := colly.NewCollector()
	var jsonData string

	//c.OnRequest(func(r *colly.Request) {
	//	r.Headers.Set("Content-Type", "application/json;charset=UTF-8")
	//})

	// extract status code
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("[*] response received", r.StatusCode)
		//fmt.Println(r.Ctx.Get("url"))
		// 打印body html源码
		//fmt.Println("[*] Source Code", string(r.Body))
		jsonData = string(r.Body)
		//fmt.Println(jsonData)
	})

	c.PostRaw(FULL_API_URL, payload)
	return jsonData
}

type ApiResJson struct {
	TotalCount int `json:"totalCount"`
	Data       []struct {
		S string    `json:"s"`
		D []float64 `json:"d"`
	} `json:"data"`
}

type TradingViewAnalysis struct {
	RecommendOther float64 // Recommend.Other
	RecommendAll   float64 // Recommend.All
	RecommendMA    float64 //  Recommend.MA
	RSI            float64 //  RSI
	RSI1           float64 // RSI[1]
	StochK         float64 // Stoch.K
	StochD         float64 // Stoch.D
	StochK1        float64 // Stoch.K[1]
	StochD1        float64 // Stoch.D[1]
	CCI20          float64 // CCI20
	CCI201         float64 // CCI20[1]
	ADX            float64 // ADX
	ADXPlusDI      float64 // ADX+DI
	ADXMinusDI     float64 // ADX-DI
	ADXPlusDI1     float64 // ADX+DI[1]
	ADXMinusDI1    float64 // ADX-DI[1]
	AO             float64 // AO
	AO1            float64 // AO[1]
	Mom            float64 // Mom
	Mom1           float64 // Mom[1]
	MACDMacd       float64 // MACD.macd
	MACDSignal     float64 // MACD.signal
	RecStochRSI    float64 // Rec.Stoch.RSI
	StochRSIK      float64 // Stoch.RSI.K
	RecWR          float64 // Rec.WR
	WR             float64 // W.R
	RecBBPower     float64 // Rec.BBPower
	BBPower        float64 // BBPower
	RecUO          float64 // Rec.UO
	UO             float64 // UO
	Close          float64
	EMA5           float64
	SMA5           float64
	EMA10          float64
	SMA10          float64
	EMA20          float64
	SMA20          float64
	EMA30          float64
	SMA30          float64
	EMA50          float64
	SMA50          float64
	EMA100         float64
	SMA100         float64
	EMA200         float64
	SMA200         float64
	RecIchimoku    float64 // Rec.Ichimoku
	IchimokuBLine  float64 // Ichimoku.BLine
	RecVWMA        float64
	VWMA           float64
	RecHullMA9     float64
	HullMA9        float64
	Open           float64
	PSAR           float64 // P.SAR
	BBlower        float64 // BB.lower
	BBupper        float64 // BB.upper
	AO2            float64 // AO[2]
	Volume         float64
	Change         float64
	Low            float64
	High           float64
}

func ParseJsonData(jsonData string) TradingViewAnalysis {
	var p ApiResJson
	// 解析json数据
	err := json.Unmarshal([]byte(jsonData), &p)
	if err != nil {
		common.Logger.Errorf("Parse FutureIncreaseRank Api JsonData Fail: %s", err)
	}

	fmt.Printf("[*] Api Json Data After handled: %+v \n", p)
	//return p
	var analysisData TradingViewAnalysis
	analysisData.RecommendOther = p.Data[0].D[0]
	analysisData.RecommendAll = p.Data[0].D[1]
	analysisData.RecommendMA = p.Data[0].D[2]
	analysisData.RSI = p.Data[0].D[3]
	analysisData.RSI1 = p.Data[0].D[4]
	analysisData.StochK = p.Data[0].D[5]
	analysisData.StochD = p.Data[0].D[6]
	analysisData.StochK1 = p.Data[0].D[7]
	analysisData.StochD1 = p.Data[0].D[8]
	analysisData.CCI20 = p.Data[0].D[9]
	analysisData.CCI201 = p.Data[0].D[10]
	analysisData.ADX = p.Data[0].D[11]
	analysisData.ADXPlusDI = p.Data[0].D[12]
	analysisData.ADXMinusDI = p.Data[0].D[13]
	analysisData.ADXPlusDI1 = p.Data[0].D[14]
	analysisData.ADXMinusDI1 = p.Data[0].D[15]
	analysisData.AO = p.Data[0].D[16]
	analysisData.AO1 = p.Data[0].D[17]
	analysisData.Mom = p.Data[0].D[18]
	analysisData.Mom1 = p.Data[0].D[19]
	analysisData.MACDMacd = p.Data[0].D[20]
	analysisData.MACDSignal = p.Data[0].D[21]
	analysisData.RecStochRSI = p.Data[0].D[22]
	analysisData.StochRSIK = p.Data[0].D[23]
	analysisData.RecWR = p.Data[0].D[24]
	analysisData.WR = p.Data[0].D[25]
	analysisData.RecBBPower = p.Data[0].D[26]
	analysisData.BBPower = p.Data[0].D[27]
	analysisData.RecUO = p.Data[0].D[28]
	analysisData.UO = p.Data[0].D[29]
	analysisData.Close = p.Data[0].D[30]
	analysisData.EMA5 = p.Data[0].D[31]
	analysisData.SMA5 = p.Data[0].D[32]
	analysisData.EMA10 = p.Data[0].D[33]
	analysisData.SMA10 = p.Data[0].D[34]
	analysisData.EMA20 = p.Data[0].D[35]
	analysisData.SMA20 = p.Data[0].D[36]
	analysisData.EMA30 = p.Data[0].D[37]
	analysisData.SMA30 = p.Data[0].D[38]
	analysisData.EMA50 = p.Data[0].D[39]
	analysisData.SMA50 = p.Data[0].D[40]
	analysisData.EMA100 = p.Data[0].D[41]
	analysisData.SMA100 = p.Data[0].D[42]
	analysisData.EMA200 = p.Data[0].D[43]
	analysisData.SMA200 = p.Data[0].D[44]
	analysisData.RecIchimoku = p.Data[0].D[45]
	analysisData.IchimokuBLine = p.Data[0].D[46]
	analysisData.RecVWMA = p.Data[0].D[47]
	analysisData.VWMA = p.Data[0].D[48]
	analysisData.RecHullMA9 = p.Data[0].D[49]
	analysisData.HullMA9 = p.Data[0].D[50]
	analysisData.Open = p.Data[0].D[51]
	analysisData.PSAR = p.Data[0].D[52]
	analysisData.BBlower = p.Data[0].D[53]
	analysisData.BBupper = p.Data[0].D[54]
	analysisData.AO2 = p.Data[0].D[55]
	analysisData.Volume = p.Data[0].D[56]
	analysisData.Change = p.Data[0].D[57]
	analysisData.Low = p.Data[0].D[58]
	analysisData.High = p.Data[0].D[59]

	//fmt.Println("analysisData: ", analysisData.BBlower)
	//fmt.Println("analysisData: ", analysisData.BBupper)
	return analysisData
}
