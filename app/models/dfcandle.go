package models

import "time"

// UIにcandleのデータを渡す

// DataFrameに対して特定の情報だけ取れるようにする
// Candleは同じmodelsパッケージのcandle.goにある構造体Candle
// ProductCodeとDurationはDBテーブルを指定するために必要な情報で、
// 構造体Candleからは必要な情報だけ取れるようにCandlesに格納している？
type DataFrameCandle struct {
	ProductCode string        `json:"product_code"`
	Duration    time.Duration `json:"duration"`
	Candles     []Candle      `json:"candles"`
}

// DataFrameCandleのtimeだけが入っているスライスを返すことができる関数
func (df *DataFrameCandle) Times() []time.Time {
	// 変数名sのスライスを作る,
	// キャパシティがdfcandleのlenでtimeが入るスライス
	s := make([]time.Time, len(df.Candles))

	// dfCandleの長さでfor文を回す
	for i, candle := range df.Candles {
		// dfCandleのtimeだけが入っているスライスを返す
		s[i] = candle.Time
	}
	return s
}

func (df *DataFrameCandle) Opens() []float64 {
	s := make([]float64, len(df.Candles))
	for i, candle := range df.Candles {
		s[i] = candle.Open
	}
	return s
}

func (df *DataFrameCandle) Closes() []float64 {
	s := make([]float64, len(df.Candles))
	for i, candle := range df.Candles {
		s[i] = candle.Close
	}
	return s
}

func (df *DataFrameCandle) Highs() []float64 {
	s := make([]float64, len(df.Candles))
	for i, candle := range df.Candles {
		s[i] = candle.High
	}
	return s
}

func (df *DataFrameCandle) Low() []float64 {
	s := make([]float64, len(df.Candles))
	for i, candle := range df.Candles {
		s[i] = candle.Low
	}
	return s
}

func (df *DataFrameCandle) Volume() []float64 {
	s := make([]float64, len(df.Candles))
	for i, candle := range df.Candles {
		s[i] = candle.Volume
	}
	return s
}
