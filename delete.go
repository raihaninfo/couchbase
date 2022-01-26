package main

import "fmt"

func Delete() {
	pres := db.Query("DELETE FROM royaltypool.raihan.client WHERE id='4444'")
	fmt.Println(pres.Result, pres.Errors)
}
