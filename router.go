package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func router() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/ping", pingController).Methods("GET")
	myRouter.HandleFunc("/allCurrencies", allCurrencies).Methods("GET")
	myRouter.HandleFunc("/addCurrency", addCurrency).Methods("POST")
	myRouter.HandleFunc("/deleteCurrency/{id}", deleteCurrency).Methods("DELETE")
	myRouter.HandleFunc("/addRate", addRate).Methods("POST")
	myRouter.HandleFunc("/rates", showRates).Methods("GET")
	myRouter.HandleFunc("/rates/{date}", showRatesByDate).Methods("GET")
	myRouter.HandleFunc("/recentRates/{from}/{to}", showRecentRates).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+os.Getenv("APP_PORT"), myRouter))
}
