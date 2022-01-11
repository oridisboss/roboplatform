package main

import "net/http"

type menuItem struct {
	Name    string
	Url     string
	Handler func(w http.ResponseWriter, r *http.Request)
}

type menuTree struct {
	Name  string
	Items []menuItem
}

var Menues = []menuTree{}

type TPLData struct {
	Menues     []menuTree
	RequestURI string
	User       User
}

type MainTPLData struct {
	Main TPLData
}

func GetTPLData(r *http.Request) TPLData {
	var data TPLData
	data.RequestURI = r.RequestURI
	data.Menues = Menues
	data.User = MyUser
	return data
}

func GetMainTPLData(r *http.Request) MainTPLData {
	var data MainTPLData
	data.Main = GetTPLData(r)
	return data
}

func GetTemplatesFiles(extraFile string) []string {
	return []string{
		"templates/" + extraFile,
		"templates/body.html",
		"templates/footer.html",
		"templates/head.html",
	}
}

var MainMenu = []menuTree{
	{"Data", []menuItem{
		{"Robots", "/robots", profileHandler},
		{"Shares", "/papers", papersHandler},
	},
	},
	{"System", []menuItem{
		{"Home", "/", indexHandler},
		{"Profile", "/profile", profileHandler},
	},
	},
}
