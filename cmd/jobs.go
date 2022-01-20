package main

import (
	"encoding/json"
	"fmt"
	"log"

	ia "github.com/oridisboss/investAPI"
)

func LoadPapers(Type string) {
	if Type == "shares" {
		LoadShare(Type)
	}
	if Type == "bonds" {
		LoadBonds(Type)
	}
	if Type == "etfs" {
		LoadEtfs(Type)
	}
	if Type == "futures" {
		LoadFutures(Type)
	}
	if Type == "currencies" {
		LoadCurrencies(Type)
	}
}

func LoadFutures(Type string) {
	var ir ia.InstrumentsRequest
	ir.InstrumentStatus = ia.InstrumentStatus_INSTRUMENT_STATUS_ALL
	var l LogItem
	l.StartLog("Instruments.Futures " + ir.String())

	r, err := Instruments.Futures(ctx, &ir)
	if err != nil {
		st := conn.GetState()
		text := fmt.Sprintf("Error loading %s  %v %v \n", Type, err, st)
		l.StopLog(text)

		return
	}
	l.StopLog("OK")

	shares := r.GetInstruments()

	for i := range shares {
		p := GetPaperByFIGI(shares[i].Figi)

		p.FIGI = shares[i].Figi
		p.ISIN = "-"
		p.Ticker = shares[i].Ticker
		p.Lot = int64(shares[i].Lot)
		p.Name = shares[i].Name
		p.Country = shares[i].CountryOfRisk
		p.Exchange = shares[i].Exchange
		p.MinPriceIncrement = float64(shares[i].MinPriceIncrement)
		p.Type = Type
		p.TradingStatus = int64(shares[i].TradingStatus)
		p.ApiTradeAvailableFlag = shares[i].ApiTradeAvailableFlag

		j, err := json.Marshal(shares[i])
		if err == nil {
			p.Extra = string(j)
		}

		dbg.Save(&p)
	}
	ClearPapers()
}

func LoadEtfs(Type string) {
	var ir ia.InstrumentsRequest
	ir.InstrumentStatus = ia.InstrumentStatus_INSTRUMENT_STATUS_ALL

	var l LogItem
	l.StartLog("Instruments.Etfs " + ir.String())

	r, err := Instruments.Etfs(ctx, &ir)
	if err != nil {
		st := conn.GetState()
		text := fmt.Sprintf("Error loading %s  %v %v \n", Type, err, st)
		l.StopLog(text)
		return
	}
	shares := r.GetInstruments()
	for i := range shares {
		p := GetPaperByFIGI(shares[i].Figi)

		p.FIGI = shares[i].Figi
		p.ISIN = shares[i].Isin
		p.Ticker = shares[i].Ticker
		p.Lot = int64(shares[i].Lot)
		p.Name = shares[i].Name
		p.Country = shares[i].CountryOfRisk
		p.Exchange = shares[i].Exchange
		p.MinPriceIncrement = float64(shares[i].MinPriceIncrement)
		p.Type = Type
		p.TradingStatus = int64(shares[i].TradingStatus)
		p.ApiTradeAvailableFlag = shares[i].ApiTradeAvailableFlag

		j, err := json.Marshal(shares[i])
		if err == nil {
			p.Extra = string(j)
		}

		dbg.Save(&p)
	}
	ClearPapers()
}

func LoadCurrencies(Type string) {
	var ir ia.InstrumentsRequest
	ir.InstrumentStatus = ia.InstrumentStatus_INSTRUMENT_STATUS_ALL

	var l LogItem
	l.StartLog("Instruments.Currencies " + ir.String())

	r, err := Instruments.Currencies(ctx, &ir)
	if err != nil {
		st := conn.GetState()
		text := fmt.Sprintf("Error loading %s  %v %v \n", Type, err, st)
		l.StopLog(text)
		return
	}
	l.StopLog("OK")

	shares := r.GetInstruments()
	for i := range shares {
		p := GetPaperByFIGI(shares[i].Figi)

		p.FIGI = shares[i].Figi
		p.ISIN = shares[i].Isin
		p.Ticker = shares[i].Ticker
		p.Lot = int64(shares[i].Lot)
		p.Name = shares[i].Name
		p.Country = shares[i].CountryOfRisk
		p.Exchange = shares[i].Exchange
		p.MinPriceIncrement = float64(shares[i].MinPriceIncrement)
		p.Type = Type
		p.TradingStatus = int64(shares[i].TradingStatus)
		p.ApiTradeAvailableFlag = shares[i].ApiTradeAvailableFlag

		j, err := json.Marshal(shares[i])
		if err == nil {
			p.Extra = string(j)
		}

		dbg.Save(&p)
	}
	ClearPapers()
}

func LoadShare(Type string) {
	var ir ia.InstrumentsRequest
	ir.InstrumentStatus = ia.InstrumentStatus_INSTRUMENT_STATUS_ALL

	var l LogItem
	l.StartLog("Instruments.Shares " + ir.String())

	r, err := Instruments.Shares(ctx, &ir)
	if err != nil {
		st := conn.GetState()
		text := fmt.Sprintf("Error loading %s  %v %v \n", Type, err, st)
		l.StopLog(text)
		return
	}
	l.StopLog("OK")

	shares := r.GetInstruments()
	for i := range shares {
		p := GetPaperByFIGI(shares[i].Figi)

		p.FIGI = shares[i].Figi
		p.ISIN = shares[i].Isin
		p.Ticker = shares[i].Ticker
		p.Lot = int64(shares[i].Lot)
		p.Name = shares[i].Name
		p.Country = shares[i].CountryOfRisk
		p.Exchange = shares[i].Exchange
		p.MinPriceIncrement = float64(shares[i].MinPriceIncrement)
		p.Type = Type
		p.TradingStatus = int64(shares[i].TradingStatus)
		p.ApiTradeAvailableFlag = shares[i].ApiTradeAvailableFlag

		j, err := json.Marshal(shares[i])
		if err == nil {
			p.Extra = string(j)
		}

		dbg.Save(&p)
	}
	ClearPapers()
}

func LoadBonds(Type string) {
	var ir ia.InstrumentsRequest
	ir.InstrumentStatus = ia.InstrumentStatus_INSTRUMENT_STATUS_ALL
	var l LogItem
	l.StartLog("Instruments.Bonds " + ir.String())

	r, err := Instruments.Bonds(ctx, &ir)
	if err != nil {
		st := conn.GetState()
		text := fmt.Sprintf("Error loading %s  %v %v \n", Type, err, st)
		l.StopLog(text)
		return
	}
	l.StopLog("OK")

	log.Println("loading... \n", err)
	shares := r.GetInstruments()
	for i := range shares {
		p := GetPaperByFIGI(shares[i].Figi)

		p.FIGI = shares[i].Figi
		p.ISIN = shares[i].Isin
		p.Ticker = shares[i].Ticker
		p.Lot = int64(shares[i].Lot)
		p.Name = shares[i].Name
		p.Country = shares[i].CountryOfRisk
		p.Exchange = shares[i].Exchange
		p.MinPriceIncrement = float64(shares[i].MinPriceIncrement)
		p.Type = Type
		p.TradingStatus = int64(shares[i].TradingStatus)
		p.ApiTradeAvailableFlag = shares[i].ApiTradeAvailableFlag

		j, err := json.Marshal(shares[i])
		if err == nil {
			p.Extra = string(j)
		}

		dbg.Save(&p)
	}
	ClearPapers()
}

func AddJob(jobType string, jobParam string) {

	if jobType == "reload_papers" {
		println("reload " + jobParam)
		LoadPapers(jobParam)
	}
}
