package main

//Paper is xxx
type Paper struct {
	ID                    int64   `json:"Id" gorm:"primaryKey"`
	FIGI                  string  `json:"FIGI"`
	ISIN                  string  `json:"ISIN"`
	Ticker                string  `json:"Ticker"`
	Type                  string  `json:"Type"`
	Name                  string  `json:"Name"`
	Country               string  `json:"Country"`
	Exchange              string  `json:"Exchange"`
	TradingStatus         int64   `json:"TradingStatus"`
	Lot                   int64   `json:"Lot"`
	MinPriceIncrement     float64 `json:"MinPriceIncrement"`
	Extra                 string  `json:"Extra"`
	ApiTradeAvailableFlag bool    `json:"apiTradeAvailableFlag"`
}
