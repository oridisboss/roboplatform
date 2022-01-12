package main

import "sync"

var muPapers sync.Mutex
var gPapers map[string]*[]Paper

//GetPaper is xxx
func GetPaper(paperID int64) Paper {
	var p Paper
	muPapers.Lock()
	dbg.First(&p, paperID)
	muPapers.Unlock()
	return p
}

//GetPaper is xxx
func GetPaperByFIGI(figi string) Paper {
	var p Paper
	muPapers.Lock()
	dbg.Where("FIGI=?", figi).First(&p)
	muPapers.Unlock()
	return p
}

func ClearPapers() {
	muPapers.Lock()
	gPapers = nil
	muPapers.Unlock()
}

//GetPapers return list of all papers
func GetPapers(PaperType string, param string) *[]Paper {
	index := PaperType + param

	muPapers.Lock()
	if len(gPapers) == 0 {
		gPapers = make(map[string]*[]Paper)
	}
	if gPapers[index] == nil {
		gPapers[index] = new([]Paper)
	}

	if len(*gPapers[index]) == 0 {
		if param == "" {
			dbg.Where("Type='" + PaperType + "'").Find(gPapers[index])
		} else {
			println("Type='" + PaperType + "' AND " + param)
			dbg.Where("Type='" + PaperType + "' AND " + param).Find(gPapers[index])
		}
	}
	muPapers.Unlock()

	return gPapers[index]
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
