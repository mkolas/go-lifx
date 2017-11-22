package main

import (
	"github.com/mkolas/go-lifx"

	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	time2 "time"
)

type Configuration struct {
	Token string
	Port  string
}

var (
	db *sql.DB
)

func main() {
	confFile, err := os.Open("config/token.json")
	if err != nil {
		log.Fatalln(err)
	}
	decoder := json.NewDecoder(confFile)
	configuration := Configuration{}
	err = decoder.Decode(&configuration)
	if err != nil {
		log.Fatalln(err)
	}

	go_lifx.Init(configuration.Token)

	// init a database

	db, err = sql.Open("sqlite3", "config/colors.db")
	if err != nil {
		log.Fatal("Unable to create database")
	}
	createStmt := `
	create table if not exists colors(name text, color text, time datetime);
	`

	defer db.Close()

	_, err = db.Exec(createStmt)
	if err != nil {
		log.Printf("%q, %s\n", err, createStmt)
	}

	http.Handle("/lifx/", http.StripPrefix("/lifx/", http.FileServer(http.Dir("web"))))
	http.Handle("/change", http.HandlerFunc(handleChange))
	http.Handle("/recent", http.HandlerFunc(getRecent))
	http.ListenAndServe(configuration.Port, nil)
	fmt.Println("Serving light changer")
	<-make(chan struct{})
	return
}

func getLastEntries() (lastEntries []*go_lifx.Entry) {
	rows, err := db.Query("select * from colors order by time desc limit 10")
	if err != nil {
		log.Fatalln("Failed to retrieve rows.", err)
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		var color string
		var time time2.Time
		err = rows.Scan(&name, &color, &time)
		if err != nil {
			log.Fatalln("Failed to read row.", err)
		}
		entry := &go_lifx.Entry{
			Name:  name,
			Color: color,
			Time:  time,
		}
		lastEntries = append(lastEntries, entry)
	}
	return
}

func getRecent(w http.ResponseWriter, r *http.Request) {
	lastEntries := getLastEntries()
	t, _ := template.ParseFiles("web/templates/entry.html.tmpl")
	t.Execute(w, lastEntries)
}

func handleChange(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request...")
	color := r.FormValue("color")
	name := r.FormValue("name")
	fmt.Println("Color is... ", color)
	entry := go_lifx.LocalChange(color, name)

	insertStmt, err := db.Prepare("insert into colors(name, color, time) values(?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	insertStmt.Exec(entry.Name, entry.Color, entry.Time)

	lastEntries := getLastEntries()

	t, _ := template.ParseFiles("web/templates/entry.html.tmpl")
	t.Execute(w, lastEntries)
}
