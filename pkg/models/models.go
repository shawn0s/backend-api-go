package models

import (
	"gorm.io/gorm"
)

type Params struct {
	Type       string `json:"type"`
	DayDate    string `json:"dayDate"`
	WeekDate   string `json:"weekDate"`
	MonthDate  string `json:"monthDate"`
	Response   string `json:"response"`
	Controller string `json:"controller"`
	Action     string `json:"action"`
	Lang       string `json:"lang"`
	Timestamp  string `json:"_"`
}

type Stock struct {
	gorm.Model
	StockNo            string  `json:"stockNo"`
	StockName          string  `json:"stockName"`
	StockType          string  `json:"stockType"`
	StockDate          string  `json:"stockDate"`
	Date               string  `json:"date"`
	Amount             int     `json:"amount"`
	Avg                float32 `json:"avg"`
	Change             float32 `json:"change"`
	Close              float32 `json:"close"`
	Diff               float32 `json:"diff"`
	High               float32 `json:"high"`
	Low                float32 `json:"low"`
	Open               float32 `json:"open"`
	PeRatio            float32 `json:"peRatio"`
	Volume             int     `json:"volume"`
	DealerBuy          int     `json:"dealerBuy"`
	DealerSell         int     `json:"dealerSell"`
	DealerTotal        int     `json:"dealerTotal"`
	DealerTotalAmount  int     `json:"dealerTotalAmount"`
	ForeignBuy         int     `json:"foreignBuy"`
	ForeignSell        int     `json:"foreignSell"`
	ForeignTotal       int     `json:"foreignTotal"`
	ForeignTotalAmount int     `json:"foreignTotalAmount"`
	FundBuy            int     `json:"fundBuy"`
	FundSell           int     `json:"fundSell"`
	FundTotal          int     `json:"fundTotal"`
	FundTotalAmount    int     `json:"fundTotalAmount"`
}
