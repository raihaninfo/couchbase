package main

import "fmt"

func InsertData() {
	pres := db.Query(`INSERT INTO royaltypool.raihan.login (KEY, VALUE) VALUES('555',
	{ "name": "Raihan",
	  "country": "bangladesh",
	  "id": "4444",
	   "type": "admin"} )`)
	fmt.Println(pres.Result, pres.Errors)
}
