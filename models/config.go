package models

// Config 配置
type Config struct {
	RestUrl      string // RestUrl okx域名
	TgToken      string // TgBotToken
	TgUserIdList []int
	ApiKey       string //
	SecretKey    string //
	Passphrase   string //
	Env          string // 生产or测试环境
	DexName      string // 交易所名称
}

type ApiConfig struct {
}
