package go_lifx

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

const AUTH = "https://api.lifx.com/v1/lights/all"
const DELTA = "https://api.lifx.com/v1/lights/all/state"

var (
	authString string
	client     = &http.Client{}
)

type Entry struct {
	Name  string
	Color string
	Time  time.Time
}

func Init(token string) (status string) {

	authString = fmt.Sprintf("Bearer %s", token)
	req, err := http.NewRequest("GET", AUTH, nil)
	if err != nil {
		log.Fatalln("Failed to create request", err)
	}
	req.Header.Set("Authorization", authString)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln("Failed to authenticate", err)
	}
	if resp.Status != "200 OK" {
		log.Fatalln("Failed to authenticate", resp.Status)
	}
	log.Printf("Successfully authenticated")

	status = resp.Status
	return
}

func LocalChange(color string, name string) (entry Entry) {
	body := strings.NewReader(fmt.Sprintf("color=%s", color))
	req, err := http.NewRequest("PUT", DELTA, body)
	if err != nil {
		log.Fatalln("Failed to create request", err)
	}
	req.Header.Set("Authorization", authString)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln("Failed to change color")
	}
	log.Println(resp.Body.Close())
	entry = Entry{name, color, time.Now()}
	return
}
