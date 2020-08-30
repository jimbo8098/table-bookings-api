package main

import (
	"encoding/json"
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

var bookingList []Booking
var booking Booking

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
		err := json.Unmarshal(r.Body, &booking)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
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

	/*
		fmt.Println(string(bookingJSON))

		newbooking := Booking{}
		err = json.Unmarshal(bookingJSON, &newbooking)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Booking name %s", newbooking.Name)
	*/

	http.HandleFunc("/booking", bookingHandler)
	http.ListenAndServe(":5000", nil)
}
