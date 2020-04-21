package main

import (
	"github.com/gorilla/mux"
	"TutorialGoRestApi-Firebird/apis"
	"log"
	"net/http"
)

func main(){

	router := mux.NewRouter()

	router.HandleFunc("/restapi/sleeph5",apis.GetAllSleeph5).Methods("GET")

	err := http.ListenAndServe(":1234",router)

	if err != nil {
		log.Println(err)
	}

}