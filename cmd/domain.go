package main

import (
	"time"
)

//Sector is xxx
type Sector struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

//Countries is xxx
type Countries struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

//Paper is xxx
type Paper struct {
	ID                int64   `json:"id"`
	FIGI              string  `json:"FIGI"`
	ISIN              string  `json:"ISIN"`
	Ticker            string  `json:"ticker"`
	Type              int64   `json:"type"`
	Name              string  `json:"name"`
	SectorID          int64   `json:"sector_id"`
	CountryID         int64   `json:"country_id"`
	ExchangeID        int64   `json:"exchange_id"`
	DTStart           string  `json:"dt_start"`
	HistoryStartDate  string  `json:"history_start_date"`
	DtLastparse       string  `json:"dt_lastparse"`
	DtLastparseMin    string  `json:"dt_lastparse_min"`
	TradeStart        string  `json:"trade_start"`
	TradeStop         string  `json:"trade_stop"`
	Lots              int64   `json:"lots"`
	MinPriceIncrement float64 `json:"min_price_increment"`
	TcsParams         string  `json:"tcs_params"`
}

// PaperCandle is shit
type PaperCandle struct {
	ID      int64     `json:"id"`
	PaperID int64     `json:"paper_id"`
	O       float64   `json:"o"`
	C       float64   `json:"c"`
	H       float64   `json:"h"`
	L       float64   `json:"l"`
	V       int64     `json:"v"`
	D       time.Time `json:"d"`
}

// Avg is return avg
func (c PaperCandle) Avg() float64 {
	return (c.C + c.O + c.H + c.L) / 4
}

//Weekend is shit
type Weekend struct {
	ID            int64     `json:"id"`
	D             time.Time `json:"d"`
	IsWeekend     string    `json:"is_weekend"`
	BeforeWeekend string    `json:"before_weekend"`
	AfterWeekend  string    `json:"after_weekend"`
}
