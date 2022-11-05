package config

import (
	"github.com/jerrychan807/go-indicator/models"
	"github.com/spf13/viper"
	"log"
	"runtime"
)

func InitConfig() {
	//path, err := os.Getwd()
	//if err != nil {
	//	//	panic(err)
	//	//}
	//viper.AddConfigPath("/jcoin/cake-syrup-pools-lover/config")
	sysType := runtime.GOOS
	if sysType == "linux" { // LINUX系统
		viper.AddConfigPath("/jcoin/go-indicator/config")
	}
	if sysType == "windows" { // windows系统
		viper.AddConfigPath("D:\\code\\github\\go-indicator\\config")
	}
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
}

// @title 读取配置文件中的配置
// @description 读取配置,返回配置结构体
func GetConfig() models.Config {
	var AllConfig models.Config
	InitConfig()
	// 交易所配置
	AllConfig.RestUrl = viper.GetString("okx.restUrl")
	AllConfig.ApiKey = viper.GetString("dexsecret.apiKey")
	AllConfig.SecretKey = viper.GetString("dexsecret.secretKey")
	AllConfig.Passphrase = viper.GetString("dexsecret.passphrase")
	AllConfig.Env = viper.GetString("dexsecret.env")
	AllConfig.DexName = viper.GetString("dexsecret.dexName")

	// TgBot配置
	AllConfig.TgToken = viper.GetString("tg.token")
	AllConfig.TgUserIdList = viper.GetIntSlice("alert.tgUserIdList")

	return AllConfig
}
