package controllers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gotrading/app/models"
	"github.com/gotrading/config"
)

// htmlファイルをキャッシュする？
var templates = template.Must(template.ParseFiles("app/views/chart.html"))

func viewChartHandler(w http.ResponseWriter, r *http.Request) {
	// ExecuteTemplateの第３引数は渡したいデータ
	err := templates.ExecuteTemplate(w, "chart.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// JSONのエラーを返すstruct, あとでmarshalして返す
type JSONError struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

//　エラーをJSONで返すための関数？
func APIError(w http.ResponseWriter, errMessage string, code int) {
	// httpヘッダの情報を追加
	w.Header().Set("Content-Type", "application/json")

	// 第３引数のcodeをヘッダに追加
	w.WriteHeader(code)

	// JSONErrorのstructに第二引数のerrMessageと第３引数のcodeを入れて
	// 変数jsonErrorに格納
	jsonError, err := json.Marshal(JSONError{Error: errMessage, Code: code})
	if err != nil {
		log.Fatal(err)
	}
	// ResponseWriterに変数jsonErrorを追加
	w.Write(jsonError)
}

var apiValidPath = regexp.MustCompile("^/api/candle/$")

func apiMakeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// r.URLにapiValidPathとマッチする文字列があるか見て変数mに格納
		m := apiValidPath.FindStringSubmatch(r.URL.Path)
		// apiValidPathとマッチするものがない場合はAPIErrorを返す
		if len(m) == 0 {
			APIError(w, "Not found", http.StatusNotFound)
		}
		// apiValidPathとマッチするものがあれば、
		// fnにresponseWriterとhttp.Requestを入れて返す
		fn(w, r)
	}
}

func apiCandleHandler(w http.ResponseWriter, r *http.Request) {
	// URLでkeyがproduct_codeのqueryを変数productCodeに格納
	productCode := r.URL.Query().Get("product_code")
	if productCode == "" {
		// responseWriterとエラーメッセージとステータスを返す
		APIError(w, "No product_code param", http.StatusBadRequest)
		return
	}
	strLimit := r.URL.Query().Get("limit")
	limit, err := strconv.Atoi(strLimit)
	if strLimit == "" || err != nil || limit < 0 || limit > 1000 {
		limit = 1000
	}

	duration := r.URL.Query().Get("duration")
	if duration == "" {
		duration = "1m"
	}
	durationTime := config.Config.Durations[duration]

	df, _ := models.GetAllCandle(productCode, durationTime, limit)

	js, err := json.Marshal(df)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func StartWebServer() error {
	// "/api/candle/"のURLはapiMakeHandlerの処理を行う
	http.HandleFunc("/api/candle/", apiMakeHandler(apiCandleHandler))
	http.HandleFunc("/chart/", viewChartHandler)
	return http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), nil)
}
