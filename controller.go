package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/currencies-exchange/models"
	"github.com/gorilla/mux"
)

type currencyExchange struct {
	From string `json:"from,omitempty"`
	To   string `json:"to,omitempty"`
}

func pingController(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusOK)
	response.Write([]byte("{\"response\": \"pong\"}"))
}

func allCurrencies(response http.ResponseWriter, request *http.Request) {
	currencies, err := models.GetAllFromTo()
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(response, currencies)
}

func addCurrency(response http.ResponseWriter, request *http.Request) {
	var params currencyExchange
	err := json.NewDecoder(request.Body).Decode(&params)

	if err != nil {
		panic(err)
	}

	msg := models.CreateFromTo(params.From, params.To)
	fmt.Fprintf(response, msg)
}

func deleteCurrency(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	msg := models.DeleteFromTo(id)
	fmt.Fprintf(response, msg)
}

func addRate(response http.ResponseWriter, request *http.Request) {
	var rate models.Rate

	err := json.NewDecoder(request.Body).Decode(&rate)

	if err != nil {
		panic(err)
	}

	msg := models.CreateRate(rate)
	fmt.Fprintf(response, msg)
}

func showRates(response http.ResponseWriter, request *http.Request) {
	rates, err := models.GetAllRates()

	if err != nil {
		panic(err)
	}

	fmt.Fprintf(response, rates)
}

func showRatesByDate(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	date := vars["date"]

	rates, err := models.GetRatesByDate(date)

	if err != nil {
		panic(err)
	}

	fmt.Fprintf(response, rates)
}

func showRecentRates(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	from := vars["from"]
	to := vars["to"]

	rates, err := models.GetRecentRates(from, to)

	if err != nil {
		panic(err)
	}

	fmt.Fprintf(response, rates)
}
