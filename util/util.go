package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/shopspring/decimal"
	"log"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"
)

func BigFloat4Decimal(s1 string) string {
	f1 := new(big.Float)
	f1.SetString(s1)
	sprintf := fmt.Sprintf("%.4f", f1)
	return sprintf
}

// @title float64转换为百分比%,保留两位小数
func Float64ToPercentage(value float64) string {
	return strconv.FormatFloat(value*100, 'f', 2, 64) + "%"
}

// @title float64转换字符串,保留x位小数
// @param value float64 "浮点数"
// @param prec int "保留小数位"
func Float64ToStr(value float64, prec int) string {
	return strconv.FormatFloat(value, 'f', prec, 64)
}

// @title 字符串转float64类型
func StrToFloat64(floatStr string) float64 {
	s, _ := strconv.ParseFloat(floatStr, 64)
	return s
}

// @title 字符串生成md5字符串
func Md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// @title 检查文件地址下文件是否存在的函数
func FileExists(fileAddr string) bool {
	if _, err := os.Stat(fileAddr); os.IsNotExist(err) {
		return false
	}
	return true
}

// @title 错误处理函数
func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// @title 错误处理函数
func GetNowTimeStr() string {
	nowStr := time.Now().Format("2006-01-02 15:04:05") //获取当前时间
	return nowStr
}

// @title float64取整运算,四舍五入
func GetFloat64Round(value float64) float64 {
	return math.Ceil(value - 0.5)
}

// @title 查看接口返回状态,code为0则返回true
func CheckReqStatus(code string) bool {
	if code == "0" {
		return true
	} else {
		return false
	}
}

// @title 根据精度字符串处理浮点数的小数位
// @param price float64 "浮点数"
// @param tickSz float64 "精度" 示例数据"0.000001"
func HandleFloat64ByPrecision(price float64, tickSz float64) string {
	tickSzStr := decimal.NewFromFloat(tickSz).String()
	decimals := len(strings.Split(fmt.Sprintf("%v", tickSzStr), ".")[1]) // 小数位数
	formatStr := "%." + strconv.Itoa(decimals) + "f"
	profitPrice := fmt.Sprintf(formatStr, price) // 保留小数位数
	//value, _ := strconv.ParseFloat(profitPrice, 64)
	return profitPrice
}

// @title 根据精度处理浮点数的小数位
// @param price float64 "浮点数"
// @param pricePrecision int "价格小数点位数" 示例数据"5"
func HandleFloat64ByPrecisionNum(price float64, pricePrecision int) string {
	formatStr := "%." + strconv.Itoa(pricePrecision) + "f"
	profitPrice := fmt.Sprintf(formatStr, price) // 保留小数位数
	//value, _ := strconv.ParseFloat(profitPrice, 64)
	return profitPrice
}

func StringToFloat64(value string) float64 {
	v, _ := strconv.ParseFloat(value, 64)
	return v
}

// @title 转成正的浮点数
func StringToPositiveFloat64(value string) float64 {
	positiveString := strings.Replace(value, "-", "", -1)
	v, _ := strconv.ParseFloat(positiveString, 64)
	return v
}

// @title 转成正的字符串
func StringToPositiveString(value string) string {
	positiveString := strings.Replace(value, "-", "", -1)
	return positiveString
}

func reverseArray(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}


