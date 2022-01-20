package main

import (
	"time"
)

type LogItem struct {
	Id       int64 `gorm:"primaryKey"`
	Dt       time.Time
	Duration float64
	Url      string
	Responce string
}

func (l *LogItem) StartLog(url string) {
	l.Dt = time.Now()
	l.Url = url
}

func (l *LogItem) StopLog(resp string) {
	l.Duration = float64(float64(time.Now().UnixNano()-l.Dt.UnixNano()) / float64(time.Second))
	l.Responce = resp

	l.Save()
}

func (l *LogItem) Save() {
	dbg.Save(l)
}
