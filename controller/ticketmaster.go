package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/paul-tong/buzzer_beater/model"
)

var defaultEventCount = "3"
var ticketmasterAPI string = "I5H7fVWSs6yLf8RfSfImVVbbzYQL61J7"

// get events from ticketmaster api
// serachType: "keyword", "eventId"
// TO_DO: may need to handle exceptions like when event locks some information
func searchEventFromTicketmaster(searchType string, searchWords string, count string) ([]model.EventViewModle, error) {
	var url string

	if searchType == "keyword" {
		keyWords := strings.Join(strings.Fields(searchWords), "+") // split keywords by space(may have multiple spaces), then join with +
		url = "https://app.ticketmaster.com/discovery/v2/events.json?&keyword=" + keyWords + "&size=" + count + "&apikey=" + ticketmasterAPI
	} else if searchType == "eventId" {
		eventId := searchWords
		url = "https://app.ticketmaster.com/discovery/v2/events.json?&id=" + eventId + "&size=1&apikey=" + ticketmasterAPI
	} else {
		log.Println("Undifined event serach type!")
	}

	// get response from url and
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Request ticketmaster api error: " + err.Error())
	}
	defer resp.Body.Close()

	// read content of the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Read ticketmaster response error: " + err.Error())
	}

	// convert json to struct, this struct contains cover like "Embeded"
	eventWithCover, err := UnmarshalTicketmasterEvent(body)
	if err != nil {
		log.Println("Unmarshal ticketmaster event error: " + err.Error())
	}

	// get the list of events
	events := eventWithCover.Embedded.Events
	if len(events) == 0 {
		log.Println("Event count returned from ticketmaster is zero")
	}

	// build a list of eventViewModel
	var eventsModel []model.EventViewModle
	for _, event := range events {
		var eventModel = model.EventViewModle{ID: event.ID, Name: event.Name, Date: event.Dates.Start.LocalDate, URL: event.URL, SeatMap: event.Seatmap.StaticURL}
		eventsModel = append(eventsModel, eventModel)
		//fmt.Println(event.Name)
	}

	// convert json into map
	/*var data map[string]interface{} // may, key is string, value is interface(can be any data type)
	json.Unmarshal(body, &data)

	// get useful infromation from map
	events := (data["_embedded"].(map[string]interface{})["events"]).([]interface{}) // convert data["embedded"] to map
	for _, event := range events {
		eventMap := event.(map[string]interface{}) // convert to map
		name := eventMap["name"]
		id := eventMap["id"]
		url := eventMap["url"]
		date := eventMap["dates"].(map[string]interface{})["start"].(map[string]interface{})["localDate"]
		seatMap := eventMap["seatmap"].(map[string]interface{})["staticUrl"]
		fmt.Println("name: " + name.(string) + " id: " + id.(string) + " url: " + url.(string) + " seatMap: " + seatMap.(string) + " date: " + date.(string))
	}*/

	return eventsModel, err
}

// Structs used to convert event json gotten from ticketmaster api
// Generated from https://app.quicktype.io/
// To parse and unparse this JSON data, add this code to your project and do:
//
//    ticketmasterEvent, err := UnmarshalTicketmasterEvent(bytes)
//    bytes, err = ticketmasterEvent.Marshal()
func UnmarshalTicketmasterEvent(data []byte) (TicketmasterEvent, error) {
	var r TicketmasterEvent
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TicketmasterEvent) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type TicketmasterEvent struct {
	Embedded TicketmasterEventEmbedded `json:"_embedded"`
	Links    TicketmasterEventLinks    `json:"_links"`
	Page     Page                      `json:"page"`
}

type TicketmasterEventEmbedded struct {
	Events []Event `json:"events"`
}

type Event struct {
	Name            string           `json:"name"`
	Type            string           `json:"type"`
	ID              string           `json:"id"`
	Test            bool             `json:"test"`
	URL             string           `json:"url"`
	Locale          string           `json:"locale"`
	Images          []Image          `json:"images"`
	Sales           Sales            `json:"sales"`
	Dates           Dates            `json:"dates"`
	Classifications []Classification `json:"classifications"`
	Promoter        Promoter         `json:"promoter"`
	Promoters       []Promoter       `json:"promoters"`
	PleaseNote      string           `json:"pleaseNote"`
	PriceRanges     []PriceRange     `json:"priceRanges"`
	Products        []Product        `json:"products"`
	Seatmap         Seatmap          `json:"seatmap"`
	TicketLimit     TicketLimit      `json:"ticketLimit"`
	Links           EventLinks       `json:"_links"`
	Embedded        EventEmbedded    `json:"_embedded"`
}

type Classification struct {
	Primary  bool  `json:"primary"`
	Segment  Genre `json:"segment"`
	Genre    Genre `json:"genre"`
	SubGenre Genre `json:"subGenre"`
	Type     Genre `json:"type"`
	SubType  Genre `json:"subType"`
	Family   bool  `json:"family"`
}

type Genre struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Dates struct {
	Start            Start  `json:"start"`
	Timezone         string `json:"timezone"`
	Status           Status `json:"status"`
	SpanMultipleDays bool   `json:"spanMultipleDays"`
}

type Start struct {
	LocalDate      string `json:"localDate"`
	LocalTime      string `json:"localTime"`
	DateTime       string `json:"dateTime"`
	DateTBD        bool   `json:"dateTBD"`
	DateTBA        bool   `json:"dateTBA"`
	TimeTBA        bool   `json:"timeTBA"`
	NoSpecificTime bool   `json:"noSpecificTime"`
}

type Status struct {
	Code string `json:"code"`
}

type EventEmbedded struct {
	Venues      []Venue      `json:"venues"`
	Attractions []Attraction `json:"attractions"`
}

type Attraction struct {
	Name            string           `json:"name"`
	Type            string           `json:"type"`
	ID              string           `json:"id"`
	Test            bool             `json:"test"`
	URL             string           `json:"url"`
	Locale          string           `json:"locale"`
	Images          []Image          `json:"images"`
	Classifications []Classification `json:"classifications"`
	UpcomingEvents  UpcomingEvents   `json:"upcomingEvents"`
	Links           AttractionLinks  `json:"_links"`
}

type Image struct {
	Ratio       *Ratio  `json:"ratio,omitempty"`
	URL         string  `json:"url"`
	Width       int64   `json:"width"`
	Height      int64   `json:"height"`
	Fallback    bool    `json:"fallback"`
	Attribution *string `json:"attribution,omitempty"`
}

type AttractionLinks struct {
	Self First `json:"self"`
}

type First struct {
	Href string `json:"href"`
}

type UpcomingEvents struct {
	Total        int64 `json:"_total"`
	Ticketmaster int64 `json:"ticketmaster"`
}

type Venue struct {
	Name                    string          `json:"name"`
	Type                    string          `json:"type"`
	ID                      string          `json:"id"`
	Test                    bool            `json:"test"`
	URL                     string          `json:"url"`
	Locale                  string          `json:"locale"`
	Images                  []Image         `json:"images"`
	PostalCode              string          `json:"postalCode"`
	Timezone                string          `json:"timezone"`
	City                    City            `json:"city"`
	State                   State           `json:"state"`
	Country                 Country         `json:"country"`
	Address                 Address         `json:"address"`
	Location                Location        `json:"location"`
	Markets                 []Genre         `json:"markets"`
	Dmas                    []DMA           `json:"dmas"`
	BoxOfficeInfo           BoxOfficeInfo   `json:"boxOfficeInfo"`
	ParkingDetail           string          `json:"parkingDetail"`
	AccessibleSeatingDetail string          `json:"accessibleSeatingDetail"`
	GeneralInfo             GeneralInfo     `json:"generalInfo"`
	UpcomingEvents          UpcomingEvents  `json:"upcomingEvents"`
	Links                   AttractionLinks `json:"_links"`
}

type Address struct {
	Line1 string `json:"line1"`
}

type BoxOfficeInfo struct {
	PhoneNumberDetail     string `json:"phoneNumberDetail"`
	OpenHoursDetail       string `json:"openHoursDetail"`
	AcceptedPaymentDetail string `json:"acceptedPaymentDetail"`
	WillCallDetail        string `json:"willCallDetail"`
}

type City struct {
	Name string `json:"name"`
}

type Country struct {
	Name        string `json:"name"`
	CountryCode string `json:"countryCode"`
}

type DMA struct {
	ID int64 `json:"id"`
}

type GeneralInfo struct {
	GeneralRule string `json:"generalRule"`
	ChildRule   string `json:"childRule"`
}

type Location struct {
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
}

type State struct {
	Name      string `json:"name"`
	StateCode string `json:"stateCode"`
}

type EventLinks struct {
	Self        First   `json:"self"`
	Attractions []First `json:"attractions"`
	Venues      []First `json:"venues"`
}

type PriceRange struct {
	Type     string  `json:"type"`
	Currency string  `json:"currency"`
	Min      float64 `json:"min"`
	Max      float64 `json:"max"`
}

type Product struct {
	ID   string `json:"id"`
	URL  string `json:"url"`
	Type string `json:"type"`
	Name string `json:"name"`
}

type Promoter struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Sales struct {
	Public   Public    `json:"public"`
	Presales []Presale `json:"presales"`
}

type Presale struct {
	StartDateTime string `json:"startDateTime"`
	EndDateTime   string `json:"endDateTime"`
	Name          string `json:"name"`
}

type Public struct {
	StartDateTime string `json:"startDateTime"`
	StartTBD      bool   `json:"startTBD"`
	EndDateTime   string `json:"endDateTime"`
}

type Seatmap struct {
	StaticURL string `json:"staticUrl"`
}

type TicketLimit struct {
	Info string `json:"info"`
}

type TicketmasterEventLinks struct {
	First First `json:"first"`
	Self  First `json:"self"`
	Next  First `json:"next"`
	Last  First `json:"last"`
}

type Page struct {
	Size          int64 `json:"size"`
	TotalElements int64 `json:"totalElements"`
	TotalPages    int64 `json:"totalPages"`
	Number        int64 `json:"number"`
}

type Ratio string

const (
	The16_9 Ratio = "16_9"
	The3_2  Ratio = "3_2"
	The4_3  Ratio = "4_3"
)
