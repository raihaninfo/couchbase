package main

import "fmt"

func UpdateData() {
	pres := db.Query(`UPDATE royaltypool.raihan.client SET country = "Bangladesh" WHERE id = "4444"`)
	fmt.Println(pres.Result, pres.Errors)
}
