package main

import "sync"

//GetPaper is xxx
func GetPaper(paperID int64) Paper {
	var p Paper
	dbg.First(&p, paperID)
	return p
}

//GetPaper is xxx
func GetPaperByFIGI(figi string) Paper {
	var p Paper
	dbg.Where("FIGI=?", figi).First(&p)
	return p
}

var muPapers sync.Mutex
var gPapers map[int]*[]Paper

//GetSectors is xxx
func GetSectors() []Sector {
	var sectors []Sector
	dbg.Find(&sectors)
	return sectors
}

//GetCountries is xxx
func GetCountries() []Countries {
	var countries []Countries
	dbg.Find(&countries)
	return countries
}

//GetPapers is xxx
func GetPapers(exchangeID int) *[]Paper {
	muPapers.Lock()
	if len(gPapers) == 0 {
		gPapers = make(map[int]*[]Paper)
	}
	if gPapers[exchangeID] == nil {
		gPapers[exchangeID] = new([]Paper)
	}
	if len(*gPapers[exchangeID]) == 0 {
		if exchangeID == 0 {
			dbg.Find(gPapers[exchangeID])
		} else {
			dbg.Where("exchange_id=?", exchangeID).Find(gPapers[exchangeID])
		}
	}
	muPapers.Unlock()
	return gPapers[exchangeID]
}

//GetPapersBySector is xxx
func GetPapersBySector(sectorID int64) []Paper {
	var papers []Paper
	dbg.Where("sector_id=?", sectorID).Find(&papers)
	return papers
}

//GetPapersByCountry is xxx
func GetPapersByCountry(sectorID int64) []Paper {
	var papers []Paper
	dbg.Where("sector_id=?", sectorID).Find(&papers)
	return papers
}
