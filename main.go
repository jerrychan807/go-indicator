package main

import (
	"fmt"
	"github.com/jerrychan807/go-indicator/lib"
	"github.com/jerrychan807/okex"
	"github.com/jerrychan807/okex/responses/market"
	"math"
)

/*
首先，我们选取一个包含 9 根蜡烛的窗口（如黄色框所示），并找到窗口内蜡烛的最大值。
然后，将窗口向上移动一根蜡烛（绿色框）并找到窗口内蜡烛的最大值。
做同样的事情，直到你移动了 5 次。
如果具有最高值的蜡烛在整个步骤中保持不变，我们就有了一个枢轴点。
*/

func GetKlinesAveragePrice(res market.Candle) float64 {
	var highestPrice float64 // k线最高价
	var lowestPrice float64  // k线最低价
	for _, candle := range res.Candles {
		if candle.H > highestPrice {
			highestPrice = candle.H
		}
		// 取所有K线中的最低价
		if lowestPrice == float64(0) {
			lowestPrice = candle.L
		}
		if candle.L < lowestPrice {
			lowestPrice = candle.L
		}
	}
	averagePrice := (highestPrice + lowestPrice) / 2
	return averagePrice
}

func IsFarFromLevel(value float64, levels []float64, res market.Candle) bool {
	averagePrice := GetKlinesAveragePrice(res)
	for _, level := range levels {
		absoluteValue := math.Abs(value - level)
		if absoluteValue > averagePrice {
			return false
		}
	}
	return true
}

func main() {
	println("hello")

	client := lib.GetOkxClient()
	// 获取K线
	instId := "MASK-USDT-SWAP"
	res := lib.GetInstKLineInfo(client, instId, okex.Bar1H, 168)
	// k线数据是倒序的，第一个数据是最新的
	candlesLength := len(res.Candles)
	var max_list []float64
	var pivots []float64
	for index, _ := range res.Candles[0 : candlesLength-5] {
		if index < 5 {
			continue
		}
		// 采用 9 根蜡烛的窗口,取最大值
		//var nKlineMaxList []float64
		var current_max float64
		// 在9根蜡烛中取最大值
		for _, candle := range res.Candles[index-5 : index+4] {
			if candle.H > current_max {
				current_max = candle.H
			}
		}
		// 如果我们找到一个新的最大值，则清空 max_list
		for _, element := range max_list {
			if current_max == element {
				max_list = []float64{}
			}
		}

		max_list = append(max_list, current_max)
		// 如果移动 5 次后最大值保持不变
		if len(max_list) == 5 && IsFarFromLevel(current_max, pivots, res) {
			pivots = append(pivots, current_max)
		}
	}

	fmt.Println("pivots:", pivots)
}
