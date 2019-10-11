package main

import (
	"fmt"

	"github.com/gotrading/config"
	"github.com/gotrading/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	fmt.Println(config.Config.ApiKey)
}
