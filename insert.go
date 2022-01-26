package main

import "fmt"

func InsertData() {
	pres := db.Query(`INSERT INTO royaltypool.raihan.client (KEY, VALUE) VALUES('airline_4485',
	{ "callsign": "MY-AIR",
	  "country": "United States",
	  "iata": "Z1",
	  "icao": "AQZ",
	  "name": "80-My Air",
	  "id": "4444",
	   "type": "airline"} )`)
	fmt.Println(pres.Result, pres.Errors)
}
