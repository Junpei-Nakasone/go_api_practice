package models

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"github.com/gotrading/config"
)

const (
	tableNameSignalEvents = "signal_events"
)

// 変数DbConnectionに*sql.DBを格納することで、DBの操作ができる(?)
var DbConnection *sql.DB

func GetCandleTableName(productCode string, duration time.Duration) string {
	return fmt.Sprintf("%s_%s", productCode, duration)
}

func init() {
	var err error

	// *sql.DBが入ってる変数にsql.Openして、第一引数をドライバー、第二引数を
	// DB名にして実行し、DBを作成できる
	DbConnection, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}

	// 作成したDBで実行したいSQL文を変数cmdに格納
	cmd := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			time DATETIME PRIMARY KEY TO NOT NULL,
			product_code STRING,
			side STRING,
			price FLOAT,
			size FLOAT)`, tableNameSignalEvents)
	DbConnection.Exec(cmd)

	for _, duration := range config.Config.Durations {
		tableName := GetCandleTableName(config.Config.ProductCode, duration)
		c := fmt.Sprintf(`
			CREATE TABLE IF NOT EXISTS %s (
				time DATETIME PRIMARY KEY NOT NULL,
				open FLOAT,
				close FLOAT,
				high FLOAT,
				low open FLOAT,
				volume FLOAT)`, tableName)
		DbConnection.Exec(c)
	}
}
