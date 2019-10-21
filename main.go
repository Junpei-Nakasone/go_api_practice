package main

import (
	"github.com/gotrading/app/controllers"
	"github.com/gotrading/config"
	"github.com/gotrading/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	controllers.StreamIngestionData()
	controllers.StartWebServer()
}
