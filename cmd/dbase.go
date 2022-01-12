package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbg *gorm.DB

func DatabaseInit() {
	var err error
	dbg, err = gorm.Open(sqlite.Open("papers.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	dbg.AutoMigrate(&Paper{})
	dbg.AutoMigrate(&Setting{})
}
