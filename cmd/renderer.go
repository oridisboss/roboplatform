package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

func logHandler(w http.ResponseWriter, r *http.Request) {
	var tpl = template.Must(template.ParseFiles(GetTemplatesFiles("logs.html")...))
	var data LogData
	data.Main = GetTPLData(r)
	dbg.Find(&data.Logs)
	tpl.Execute(w, data)
}

func papersHandler(w http.ResponseWriter, r *http.Request) {
	var tpl = template.Must(template.ParseFiles(GetTemplatesFiles("papers.html")...))
	url := strings.ReplaceAll(r.RequestURI, "/", "")

	var data PaperData
	data.Main = GetTPLData(r)
	data.PaperTypeName = url

	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			log.Println(w, "ParseForm() err: %v", err)
			return
		}
		if _, ok := r.Form["type"]; ok {
			AddJob("reload_papers", url)
		}
	}

	data.PaperTypes = append(data.PaperTypes, PaperType{Name: "Base", Papers: *GetPapers(url, "api_trade_available_flag is true")})
	data.PaperTypes = append(data.PaperTypes, PaperType{Name: "All", Papers: *GetPapers(url, "")})

	tpl.Execute(w, data)
}

func profileHandler(w http.ResponseWriter, r *http.Request) {

	var tpl = template.Must(template.ParseFiles(GetTemplatesFiles("profile.html")...))

	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			log.Println(w, "ParseForm() err: %v", err)
			return
		}
		if val, ok := r.Form["token"]; ok {
			MyUser.SetToken(val[0])
		}
	}
	tpl.Execute(w, GetMainTPLData(r))
}
