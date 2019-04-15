package main

import (
	"log"
	"net/http"
	"restapi"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/ticket_options", restapi.CreateTicketOptions).Methods("POST")
	router.HandleFunc("/ticket_options/{id}", restapi.RenderTicketOptionsByID).Methods("GET")
	router.HandleFunc("/ticket_options/{id}/purchases", restapi.PurchaseTicketOptionsByID).Methods("POST")

	http.Handle("/", router)

	log.Printf("Listening on port 3000\n")

	log.Println(http.ListenAndServe(":3000", nil))

}
