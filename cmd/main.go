package main

import (
	"context"
	"database/sql"
	"flag"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	pb "investAPI"

	"google.golang.org/grpc"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MyUser User
var dbg *gorm.DB

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var tpl = template.Must(template.ParseFiles(GetTemplatesFiles("screener.html")...))
	tpl.Execute(w, GetMainTPLData(r))
}

type ProfileData struct {
	Main   TPLData
	Papers []Paper
}

func papersHandler(w http.ResponseWriter, r *http.Request) {
	var tpl = template.Must(template.ParseFiles(GetTemplatesFiles("papers.html")...))
	var data ProfileData
	data.Main = GetTPLData(r)
	data.Papers = *GetPapers(0)

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

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	flag.Parse()

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	instruments := pb.NewInstrumentsServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var ir pb.InstrumentsRequest
	ir.InstrumentStatus = pb.InstrumentStatus_INSTRUMENT_STATUS_ALL
	r, err := instruments.Shares(ctx, &ir)
	log.Println(r)

	db, err := sql.Open("mysql", "root:va123321@/go?charset=utf8")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	dsn := "root:va123321@tcp(127.0.0.1:3306)/go?charset=utf8mb4&parseTime=True&loc=Local"
	dbg, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	Menues = MainMenu

	mux := http.NewServeMux()

	mux.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	mux.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("fonts"))))
	mux.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	mux.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))

	for i := range Menues {
		for j := range Menues[i].Items {
			mux.HandleFunc(Menues[i].Items[j].Url, Menues[i].Items[j].Handler)
		}
	}

	http.ListenAndServe(":"+port, mux)
}

func NewInstrumentsServiceClient(conn *grpc.ClientConn) {
	panic("unimplemented")
}
