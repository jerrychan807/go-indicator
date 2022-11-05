package lib

import (
	"github.com/jerrychan807/go-indicator/common"
	"github.com/jerrychan807/go-indicator/util"
	"github.com/jerrychan807/okex"
	"github.com/jerrychan807/okex/api"
	"github.com/jerrychan807/okex/requests/rest/market"
	market_responses "github.com/jerrychan807/okex/responses/market"
	"github.com/sirupsen/logrus"
)

// @title 获取产品的当前价格
func GetInstCurrentPrice(response market_responses.Ticker) float64 {
	currentPrice := response.Tickers[0].Last
	common.Logger.WithFields(logrus.Fields{"currentPrice": currentPrice}).Info("GetInstCurrentPrice")
	return float64(currentPrice)
}

// @title 获取单个产品的K线数据
// @param client *api.Client "api客户端"
// @param instId string "产品ID"
// @param bar string "K线周期,时间粒度，默认值1m如 [1m/3m/5m/15m/30m/1H/2H/4H/6H/12H/1D/1W/1M/3M/6M/1Y]"
// @param kLineNum string "返回k线条数"
func GetInstKLineInfo(client *api.Client, instId string, bar okex.BarSize, kLineNum int) (response market_responses.Candle) {
	var req = market.GetCandlesticks{InstID: instId, Bar: bar, Limit: int64(kLineNum)}
	response, err := client.Rest.Market.GetCandlesticks(req)
	if err != nil {
		common.Logger.Panicf("GetInstKLineInfo Fail: %s", err)
	}
	common.Logger.WithFields(logrus.Fields{"instId": instId}).Info("GetInstKLineInfo")
	return
}

// @title 获取单个产品的合约面值
func CalculateKlineAveragePrice(response market_responses.Candle, tickSz float64) string {
	var highestPrice float64 // k线最高价
	var lowestPrice float64  // k线最低价
	for _, candle := range response.Candles {
		// 取所有K线中的最高价
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
	value := util.HandleFloat64ByPrecision(averagePrice, tickSz) // 根据精度处理价格的小数点
	common.Logger.WithFields(logrus.Fields{"highestPrice": highestPrice, "lowestPrice": lowestPrice, "averagePrice": value}).Info("CalculateKlineAveragePrice")
	return value
}
