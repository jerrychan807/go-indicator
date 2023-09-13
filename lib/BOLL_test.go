package lib

import (
	"fmt"
	"github.com/jerrychan807/go-indicator/models"
	"github.com/jerrychan807/okex"
	"testing"
	"time"
)

func TestBOLL(t *testing.T) {
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
		newKline.Close = kLineInfo.Candles[reverseIndex].C
		newKline.KlineTime = time.Time(kLineInfo.Candles[reverseIndex].TS)
		kLineList = append(kLineList, &newKline)
		//if index == 8 {
		//	break
		//}
	}
	//fmt.Println(len(kLineList))
	//fmt.Println(kLineList[0].Close)

	//计算新的BOLL
	stockList := NewBOLL(kLineList).Calculation().GetPoints()
	for _, v := range stockList {
		fmt.Printf("Time:%s\t Middle:%.5f Up:%.5f Low:%.5f\n", v.Time.Format("2006-01-02 15:04:05"), v.MID, v.UP, v.Low)
	}

}
