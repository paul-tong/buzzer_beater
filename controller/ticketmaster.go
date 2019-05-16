package controller

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var defaultEventCount = "1"
var ticketmasterAPI string = "I5H7fVWSs6yLf8RfSfImVVbbzYQL61J7"

func searchEventByKeyWords(keyWords string, count string) string {
	keyWords = strings.Join(strings.Fields(keyWords), "+") // split keywords by space(may have multiple spaces), then join with +
	url := "https://app.ticketmaster.com/discovery/v2/events.json?&keyword=" + keyWords + "&size=" + count + "&apikey=" + ticketmasterAPI
	//url := "https://app.ticketmaster.com/discovery/v2/events.json?&keyword=bostonBruins&size=2&apikey=" + ticketmasterAPI

	resp, err := http.Get(url)
	if err != nil {
		log.Println("Request ticketmaster api error: " + err.Error())
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Read ticketmaster response data error: " + err.Error())
	}

	events := string(body)
	log.Println("ticketmaster data: " + events)

	return events
}
