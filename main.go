package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type Booking struct {
	Details string `json:"details,notes"`
	Start   string `json:"start"`
	End     string `json:"end"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Name    string `json:"name"`
	Guests  int    `json:"guests"`
}

type Slot struct {
	Start string
	End   string
	Type  string
}

var bookingList []Booking
var booking Booking
var slotList []Slot

func slotHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		slotJSON, err := json.Marshal(slotList)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(slotJSON)
		return
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func bookingHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		bookingJSON, err := json.Marshal(bookingList)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(bookingJSON)
		return

	case http.MethodPost:
		reqString, err := ioutil.ReadAll(r.Body)
		if err == nil {
			err := json.Unmarshal(reqString, &booking)
			if err != nil {
				fmt.Println(err)
				io.WriteString(w, "hello")
				return
			} else {
				w.WriteHeader(http.StatusOK)
				return
			}
		}
		return
	}
}

func main() {
	booking := &Booking{
		Details: "Testing",
		Start:   "2020-09-02T18:00:00",
		End:     "2020-09-02T20:00:00",
		Email:   "test@domain.com",
		Phone:   "07111111111",
		Guests:  2,
		Name:    "Test Person",
	}
	bookingList = append(bookingList, *booking)
	bookingList = append(bookingList, *booking)
	bookingList = append(bookingList, *booking)

	slot := &Slot{
		Start: "2020-09-13 12:00:00",
		End:   "2020-09-13 13:00:00",
		Type:  "all",
	}
	slotList = append(slotList, *slot)
	slotList = append(slotList, *slot)
	slotList = append(slotList, *slot)
	slotList = append(slotList, *slot)

	http.HandleFunc("/booking", bookingHandler)
	http.HandleFunc("/slot", slotHandler)
	http.ListenAndServe(":5000", nil)
}
