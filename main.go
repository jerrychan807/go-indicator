package main

import (
	"fmt"
	"github.com/jerrychan807/go-indicator/lib"
)

func main() {
	fmt.Println("Starting")

	jsonRes := lib.GetTradingViewResponse("crypto", "BINANCE", "BTCUSDT", "1h")
	analysisData := lib.ParseJsonData(jsonRes)
	fmt.Println("analysisData: ", analysisData.BBupper)
}
