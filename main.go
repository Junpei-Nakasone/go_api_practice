package main

import (
	"fmt"

	"github.com/gotrading/app/controllers"
	"github.com/gotrading/app/models"
	"github.com/gotrading/config"
	"github.com/gotrading/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	fmt.Println(models.DbConnection)
	controllers.StreamIngestionData()
	controllers.StartWebServer()
}
