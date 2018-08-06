# Currencies Exchange API with GoLang and Postgres

## Description

This service is to manage a daily basis display foreign exchange rate for currencies

## Installation

if you have docker:
  - run docker-compose up
  - try to access `http//localhost:8081/ping`

if you dont have docker:
  - install postgres
  - install golang
  - clone this project (put it on your GOPATH)
  - install go dep
  - run dep ensure
  - create your db
  - modify .env file
  - run go build && ./currencies-exchange
  - try to access `http//localhost:{port}/ping`

## Usage

to use this API please take look on this use cases
  - User wants to input daily exchange rate data
    
    try access `/addRate` method `POST` with this request parameter structure
    ```
    {
      "rate": 5,
      "date": "2018-02-17",
      "from_to_id": 3
    }```

  - User has a list of exchange rates to be tracked
    
    try access `/rates/{date}` method `GET`

  - User wants to see the exchange rate trend from the most recent 7 data points
    
    try access `/recentRates/{from}/{to}` method `GET`

  - User wants to add an exchange rate to the list
    
    try access `/addCurrency` method `POST` with this request parameter structure
    ```
    {
      "from": "EUR",
      "to": "USD"
    }```

  - User wants to remove an exchange rate from the list
    
    try access `/deleteCurrency/{id}` method `DELETE`

  - User wants to see all all currencies data
    
    try access `/allCurrencies` method `GET`

  - User wants to see all rates data
    
    try access `/rates` method `GET`

