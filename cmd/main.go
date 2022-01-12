package main

import (
	"flag"
	"math/rand"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	flag.Parse()

	DatabaseInit()
	LoadSettings()
	GRPCConnect()

	Menues = MainMenu

	mux := http.NewServeMux()
	mux.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("static/css"))))
	mux.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("static/fonts"))))
	mux.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("static/js"))))
	mux.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("static/img"))))

	for i := range Menues {
		for j := range Menues[i].Items {
			mux.HandleFunc(Menues[i].Items[j].Url, Menues[i].Items[j].Handler)
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	http.ListenAndServe(":"+port, mux)
}
