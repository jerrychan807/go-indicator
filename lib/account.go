package lib

import (
	"context"
	"github.com/jerrychan807/go-indicator/common"
	"github.com/jerrychan807/go-indicator/config"
	"github.com/jerrychan807/okex"
	"github.com/jerrychan807/okex/api"
	"github.com/sirupsen/logrus"
	"log"
)

// @title 获取okx客户端
func GetOkxClient() *api.Client {
	allConfig := config.GetConfig()

	dest := okex.DemoServer
	if allConfig.Env == "prod" {
		dest = okex.NormalServer // The main API server
	} else {
		dest = okex.DemoServer
	}
	common.Logger.WithFields(logrus.Fields{"env": allConfig.Env}).Info("GetOkxClient")
	ctx := context.Background()
	client, err := api.NewClient(ctx, allConfig.ApiKey, allConfig.SecretKey, allConfig.Passphrase, dest)
	if err != nil {
		log.Fatalln(err)
	}
	return client
}
