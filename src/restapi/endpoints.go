package restapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// CreateTicketOptions rest api to handle the post requests for creation tickets
func CreateTicketOptions(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var ticketOptions TicketOptions

	err := decoder.Decode(&ticketOptions)
	if err != nil {
		log.Println(err.Error())
	}

	getTikcetOptions, err := CreateTicketOption(&ticketOptions)
	if err != nil {
		log.Println(err.Error())
	}

	json.NewEncoder(w).Encode(getTikcetOptions)
}

func RenderTicketOptionsByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	getTikcetOptionByID, err := GetTicketOptionByID(id)
	if err != nil {
		log.Println(err.Error())
	}

	json.NewEncoder(w).Encode(getTikcetOptionByID)
}

func PurchaseTicketOptionsByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	decoder := json.NewDecoder(r.Body)
	var purchaseTickets PurchaseTickets

	err := decoder.Decode(&purchaseTickets)
	if err != nil {
		log.Println(err.Error())
	}
	getTikcetOptionByID, err := GetTicketOptionByID(id)
	if err != nil {
		log.Println(err.Error())
	}

	w.Header().Set("Content-Type", "appliction/json")
	if purchaseTickets.Quantity > getTikcetOptionByID.Allocation {

		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte(fmt.Sprintf("%v - Something bad happened!", http.StatusExpectationFailed)))
		return
	}

	w.WriteHeader(http.StatusOK)

}
