package controllers

import (
	"log"

	"github.com/gotrading/app/models"
	"github.com/gotrading/bitflyer"
	"github.com/gotrading/config"
)

func StreamIngestionData() {
	// bitflyer.goの構造体Tickerをchannelにして
	// 変数tickerChannlに格納
	var tickerChannl = make(chan bitflyer.Ticker)

	// 自分のAPIキーでNewして変数apiClientに格納
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)

	// goroutineを回してリアルタイムでAPIから情報を取得する
	go apiClient.GetRealTimeTicker(config.Config.ProductCode, tickerChannl)

	// goroutineにすることで他の関数をブロックすることなく、
	// 並列でデータを取りに行く
	go func() {
		for ticker := range tickerChannl {
			// ログにリアルタイムで取得しているデータを出力する
			log.Printf("action=StreamIngestionData, %v", ticker)

			// 下記のfor分はログ出力とは関係ない、キャンドルを作るためのもの？
			for _, duration := range config.Config.Durations {
				isCreated := models.CreateCandleWithDuration(ticker, ticker.ProductCode, duration)

				// 下記のようにログを出すとfor文で３回出力される
				// log.Println("isCreated is working")

				if isCreated == true && duration == config.Config.TradeDuration {
					// TODO
				}
			}
		}
	}()
}
