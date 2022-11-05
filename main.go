package main

import (
	"github.com/jerrychan807/go-indicator/lib"
	"github.com/jerrychan807/okex"
)

func main() {
	println("hello")

	client := lib.GetOkxClient()
	// 获取K线
	instId := "MASK-USDT-SWAP"
	res := lib.GetInstKLineInfo(client, instId, okex.Bar1H, 72)
	// k线数据是倒序的，第一个数据是最新的
	candlesLength := len(res.Candles)
	var max_list []float64
	for index, _ := range res.Candles[5 : candlesLength-5] {
		// 采用 9 根蜡烛的窗口,取最大值
		//var nKlineMaxList []float64
		var current_max float64
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

		if len(max_list)==5 &&  {
			
		}

	}
}
