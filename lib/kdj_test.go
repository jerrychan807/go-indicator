package lib

import (
	"fmt"
	"github.com/jerrychan807/go-indicator/models"
	"github.com/jerrychan807/okex"
	"testing"
	"time"
)

func TestKDJ(t *testing.T) {
	t.Parallel()
	client := GetOkxClient()
	instId := "ETH-USDT-SWAP"
	bar := "1D"
	kLineInfo := GetInstKLineInfo(client, instId, okex.BarSize(bar), 30)
	fmt.Println(kLineInfo)
	var kLineList []*models.Kline
	for index, _ := range kLineInfo.Candles {
		var newKline models.Kline
		reverseIndex := len(kLineInfo.Candles) - 1 - index // 索引倒序
		newKline.Low = kLineInfo.Candles[reverseIndex].L   //
		newKline.High = kLineInfo.Candles[reverseIndex].H
		newKline.Open = kLineInfo.Candles[reverseIndex].O
		newKline.Close = kLineInfo.Candles[reverseIndex].C
		newKline.Vol = kLineInfo.Candles[reverseIndex].VolCcyQuote
		newKline.KlineTime = time.Time(kLineInfo.Candles[reverseIndex].TS)
		kLineList = append(kLineList, &newKline)
		//if index == 8 {
		//	break
		//}
	}
	fmt.Println(len(kLineList))
	fmt.Println(kLineList[0].Close)

	stockList := NewKDJ(kLineList, 9).Calculation().GetPoints()
	for _, v := range stockList {
		fmt.Printf("Time:%s\t Middle:%.5f Up:%.5f Low:%.5f\n", v.Time.Format("2006-01-02 15:04:05"), v.K, v.D, v.J)
	}

	//Time:2023-08-01 00:00:00         Middle:100.00000 Up:100.00000 Low:100.00000
	//Time:2023-08-02 00:00:00         Middle:83.08994 Up:94.36331 Low:60.54320
	//Time:2023-08-03 00:00:00         Middle:71.00432 Up:86.57698 Low:39.85900

}
