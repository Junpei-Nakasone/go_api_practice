package main

import (
	"fmt"

	"github.com/gotrading/bitflyer"
	"github.com/gotrading/config"
	"github.com/gotrading/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	fmt.Println(apiClient.GetBalance())
}
