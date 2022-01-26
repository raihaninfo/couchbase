package main

import "fmt"

func Read() {
	pres := db.Query("SELECT * FROM royaltypool.raihan.client WHERE id='4444'")
	fmt.Println(pres.Result)
}
