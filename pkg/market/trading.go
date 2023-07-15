package market

import (
	"backend-api-go/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type TradingFundd struct {
	Stat   string     `json:"stat"`
	Date   string     `json:"date"`
	Title  string     `json:"title"`
	Fields []string   `json:"fields"`
	Data   [][]string `json:"data"`
	Params Params
	Notes  []string `json:"notes"`
	Hints  string   `json:"hints"`
}

func TradingFund() {
	// 建立 HTTP 客戶端
	client := &http.Client{}

	// 建立 GET 請求
	req, err := http.NewRequest("GET", "https://www.twse.com.tw/rwd/zh/fund/BFI82U?type=day&dayDate=20230630&weekDate=20230626&monthDate=20230630&response=json&_=1688112429317", nil)
	if err != nil {
		fmt.Println("建立請求時發生錯誤：", err)
		return
	}

	// 發送請求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("發送請求時發生錯誤：", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	var result models.TradingFund
	// 讀取回應內容
	json.Unmarshal(body, &result)
	//json.NewDecoder(resp.Body)
	if err != nil {
		fmt.Println("讀取回應時發生錯誤：", err)
		return
	}

	// 輸出回應內容
	fmt.Println(json.Marshal(result.Params))
}
